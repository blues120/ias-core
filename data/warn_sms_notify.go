package data

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redsync/redsync/v4"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/valyala/fasthttp"

	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/conf"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
)

const (
	redisKeySendMsg    = "vics:sms"  // 控制短信发送频率的redis key
	redsyncMutexForSms = "mutex-sms" // 用于发送短信的分布式锁标识
)

type smsNotifyRepo struct {
	data      *Data
	ytxClient *ytxClient

	log *log.Helper
}

func NewSmsNotifyRepo(data *Data, ytxClient *ytxClient, logger log.Logger) biz.SmsNotifyRepo {
	return &smsNotifyRepo{
		data:      data,
		ytxClient: ytxClient,
		log:       log.NewHelper(log.With(logger, "clazz", "smsNotifyRepo")),
	}

}

// 短信错误信息biz转ent类型
func SmsErrBizToEnt(data *biz.SmsErr) *ent.WarnSmsErr {
	return &ent.WarnSmsErr{
		AppName:  data.AppName,
		RecordID: data.RecordID,
		ErrorMsg: data.ErrorMsg,
	}
}

// RecordSmsErr 保存短信发送失败记录
func (r *smsNotifyRepo) RecordSmsErr(ctx context.Context, data *biz.SmsErr) error {
	_, err := r.data.db.WarnSmsErr(ctx).Create().
		SetWarnSmsErr(SmsErrBizToEnt(data)).
		Save(ctx)

	return err
}

// FilterUnwarnedPhoneNumberList 过滤手机号：1.每个手机号一天最多发送10条短信;2.每个任务ID+手机号的组合，10分钟最多发一次短信;
func (r *smsNotifyRepo) FilterUnwarnedPhoneNumberList(ctx context.Context, taskID uint64, phoneNumberList []string) (unwarnedPhoneNumberList []string) {
	unwarnedPhoneNumberList = make([]string, 0, len(phoneNumberList))
	for _, phoneNumber := range phoneNumberList {
		// 判断手机号每日短信发送限制
		sentTimes, err := r.data.rdb.Get(ctx, getDailyWarnedKey(phoneNumber)).Int()
		if err != nil {
			if err != redis.Nil {
				r.log.Errorf("FilterUnwarnedPhoneNumberList err: %+v", err)
			}
		} else {
			maxTimes := 10 // 每日最大发送短信次数
			if sentTimes >= maxTimes {
				r.log.Debugf("FilterUnwarnedPhoneNumberList: phoneNumber: %s. The maximum number(%d) of SMS messages sent on the day is reached", phoneNumber, maxTimes)
				continue
			}
		}
		// 判断task ID+手机号的短时发送次数限制
		_, err = r.data.rdb.Get(ctx, getWarnedKey(taskID, phoneNumber)).Result()
		if err != nil {
			if err == redis.Nil {
				unwarnedPhoneNumberList = append(unwarnedPhoneNumberList, phoneNumber)
			} else {
				r.log.Errorf("FilterUnwarnedPhoneNumberList err: %+v", err)
			}
		} else {
			r.log.Debugf("FilterUnwarnedPhoneNumberList: filtered taskId %d and phoneNumber: %s", taskID, phoneNumber)
		}
	}
	return
}

// MarkPhoneAsWarned 1.更新当日的短信发送次数; 2.记录已发送报警的标识：任务ID+手机号的组合，过期时间为10分钟
func (r *smsNotifyRepo) MarkPhoneAsWarned(ctx context.Context, taskID uint64, phoneNumber string) {
	// 每日RedisKey设置
	sentTimes, err := r.data.rdb.Get(ctx, getDailyWarnedKey(phoneNumber)).Int()
	if err != nil && err != redis.Nil {
		r.log.Errorf("MarkPhoneAsWarned err: %+v", err)
		return
	} else {
		now := time.Now()
		todayEnd := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
		expireTime := time.Until(todayEnd)
		err = r.data.rdb.Set(ctx, getDailyWarnedKey(phoneNumber), sentTimes+1, expireTime).Err() // 设置过期时间到当日23:59:59
		if err != nil {
			r.log.Errorf("MarkPhoneAsWarned err: %+v", err)
			return
		}
	}
	// 短时redisKey设置
	err = r.data.rdb.Set(ctx, getWarnedKey(taskID, phoneNumber), nil, 10*time.Minute).Err()
	if err != nil {
		r.log.Errorf("MarkPhoneAsWarned err: %+v", err)
		return
	}
}

// EopPost 发送告警短信
func (r *smsNotifyRepo) EopPost(body []byte) ([]byte, error) {
	return r.ytxClient.EopPost(body)
}

// AddLock 申请加锁
func (r *smsNotifyRepo) AddLock(taskId uint64) (*redsync.Mutex, error) {
	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.
	mutex := r.data.rs.NewMutex(fmt.Sprintf("%s:%v", redsyncMutexForSms, taskId))
	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	err := mutex.Lock()
	return mutex, err
}

// ReleaseLock 释放锁
func (r *smsNotifyRepo) ReleaseLock(m *redsync.Mutex) error {
	// Release the lock so other processes or threads can obtain a lock.
	ok, err := m.Unlock()
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("unlock failed")
	}
	return nil
}

// ytxClient 短信客户端
type ytxClient struct {
	url, ak, sk string

	log *log.Helper
}

func NewYtxClient(conf *conf.YtxClient, logger log.Logger) *ytxClient {
	// 创建云通信结构体
	return &ytxClient{
		url: conf.Url,
		ak:  conf.Ak,
		sk:  conf.Sk,
		log: log.NewHelper(log.With(logger, "clazz", "ytxClient")),
	}
}

// EopPost 报警短信HTTP请求构造并发送
func (c *ytxClient) EopPost(body []byte) ([]byte, error) {
	httpclient := fasthttp.Client{}
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	// SETUP1:获取AccessKey和SecurityKey // 直接从ytxClient读取

	// SETUP2:构造时间戳
	eopDate := time.Now().Format("20060102T150405Z")

	// SETUP3:构造请求流水号
	id := uuid.New().String()

	// SETUP4:构造待签名字符串
	req.SetRequestURI(c.url)
	args := req.URI().QueryArgs()
	args.Sort(bytes.Compare)
	headerStr := "ctyun-eop-request-id:" + id + "\n" + "eop-date:" + eopDate + "\n"
	queryStr := args.String()
	calculateContentHash := getSha256(body)
	signStr := headerStr + "\n" + queryStr + "\n" + calculateContentHash

	// SETUP5:构造签名
	kTime := hmacSha256(eopDate, c.sk)
	kAk := hmacSha256(c.ak, string(kTime))
	kData := hmacSha256(eopDate[:8], string(kAk))
	signatureDate := hmacSha256(signStr, string(kData))
	signature := base64.StdEncoding.EncodeToString(signatureDate)

	// SETUP6:构造请求头
	req.Header.Add("Eop-date", eopDate)
	req.Header.Add("ctyun-eop-request-id", id)
	signatureHeader := c.ak + " Headers=ctyun-eop-request-id;eop-date Signature=" + signature
	req.Header.Add("Eop-Authorization", signatureHeader)
	c.log.Debugf("sms request: %+v", req)
	req.SetBody(body)

	if err := httpclient.Do(req, resp); err != nil {
		return nil, err
	}

	b := resp.Body()

	return b, nil
}

// 构造消息签名使用
func hmacSha256(data, secret string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))

	return h.Sum(nil)
}

// 构造消息签名使用
func getSha256(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	b := hash.Sum(nil)
	return hex.EncodeToString(b)
}

func getWarnedKey(taskID uint64, phoneNumber string) string {
	return fmt.Sprintf("%s:%d:%s", redisKeySendMsg, taskID, phoneNumber)
}

func getDailyWarnedKey(phoneNumber string) string {
	return fmt.Sprintf("%s:%s", redisKeySendMsg, phoneNumber)
}
