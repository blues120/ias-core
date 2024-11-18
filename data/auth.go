package data

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/conf"
	"github.com/blues120/ias-core/data/ent"
	"github.com/blues120/ias-core/data/ent/user"
	"github.com/blues120/ias-core/errors"
	"github.com/blues120/ias-core/pkg/codec"
)

type userRepo struct {
	data *Data

	log *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{data: data, log: log.NewHelper(logger)}
}

func (r *userRepo) Create(ctx context.Context, user *biz.User) error {
	return r.data.db.User(ctx).Create().SetName(user.Name).SetPassword(user.Password).Exec(ctx)
}

func (r *userRepo) FindByName(ctx context.Context, name string) (*biz.User, error) {
	u, err := r.data.db.User(ctx).Query().Where(user.Name(name)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrorUserNotFound("")
		}
		return nil, err
	}
	return &biz.User{
		Id:       u.ID,
		Name:     u.Name,
		Password: u.Password,
	}, nil
}

type captchaRepo struct {
	captcha *base64Captcha.Captcha

	log *log.Helper
}

func NewCaptchaRepo(conf *conf.Auth, data *Data, logger log.Logger) biz.CaptchaRepo {
	store := &redisStore{conf: conf, data: data}
	captcha := base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, store)
	return &captchaRepo{captcha: captcha, log: log.NewHelper(logger)}
}

func (r *captchaRepo) Generate(ctx context.Context) (*biz.Captcha, error) {
	id, b64, err := r.captcha.Generate()
	if err != nil {
		return nil, err
	}
	return &biz.Captcha{
		Id:        id,
		Base64Img: b64,
	}, nil
}

func (r *captchaRepo) Verify(ctx context.Context, id string, answer string) (bool, error) {
	realAnswer := r.captcha.Store.Get(id, false)
	if realAnswer == "" {
		return false, errors.ErrorUserCaptchaExpired("验证码已过期")
	}
	return answer == realAnswer, nil
}

// redisStore implement base64Captcha.Store
type redisStore struct {
	conf *conf.Auth
	data *Data
}

// Verify captcha answer directly
func (r *redisStore) Verify(id, answer string, clear bool) bool {
	return r.Get(id, clear) == answer
}

// Set sets the digits for the captcha id.
func (r *redisStore) Set(id string, value string) error {
	return r.data.rdb.SetNX(context.Background(), r.getRedisCaptcha(id), value, r.conf.CaptchaExpire.AsDuration()).Err()
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (r *redisStore) Get(id string, clear bool) string {
	ctx := context.Background()

	result, err := r.data.rdb.Get(ctx, r.getRedisCaptcha(id)).Result()
	if clear && err == nil {
		r.data.rdb.Del(ctx, r.getRedisCaptcha(id))
	}

	return result
}

func (r *redisStore) getRedisCaptcha(key string) string {
	return "ias:captcha:" + key
}

type jwtRepo struct {
	conf *conf.Auth
	data *Data

	log *log.Helper
}

func NewJwtRepo(conf *conf.Auth, data *Data, logger log.Logger) biz.JwtRepo {
	return &jwtRepo{conf: conf, data: data, log: log.NewHelper(logger)}
}

func (r *jwtRepo) NewToken(ctx context.Context, claims biz.CustomClaims, accessExpiration, refreshExpiration time.Duration) (*biz.Token, error) {
	// accessToken
	accessExpiresAt := time.Now().Add(accessExpiration)
	accessToken, err := r.newAccessToken(ctx, claims, accessExpiresAt)
	if err != nil {
		return nil, err
	}

	// refreshToken
	refreshExpiresAt := time.Now().Add(refreshExpiration)
	refreshToken, err := r.newRefreshToken(ctx, claims.UserId, refreshExpiration)
	if err != nil {
		return nil, err
	}

	return &biz.Token{
		AccessToken:            accessToken,
		AccessTokenExpireTime:  accessExpiresAt.UnixMilli(),
		RefreshToken:           refreshToken,
		RefreshTokenExpireTime: refreshExpiresAt.UnixMilli(),
	}, nil

}

func (r *jwtRepo) GetUserIdFromRefreshToken(ctx context.Context, refreshToken string) (uint64, error) {
	ret, err := r.data.rdb.Get(ctx, r.getRedisRefreshToken(refreshToken)).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, errors.ErrorUserRefreshTokenNotFound("登录状态已过期, 请重新登录")
		}
		return 0, err
	}
	userIdStr, oldVersion := r.parseRedisRefreshTokenValue(ret)
	currentVersion, err := r.data.rdb.Get(ctx, r.getRedisRefreshTokenVersion(userIdStr)).Int64()
	if err != nil && err != redis.Nil {
		return 0, err
	}
	if currentVersion != oldVersion {
		return 0, errors.ErrorUserRefreshTokenNotFound("登录状态已过期, 请重新登录")
	}
	return strconv.ParseUint(userIdStr, 10, 64)
}

func (r *jwtRepo) DeleteRefreshToken(ctx context.Context, refreshToken string) error {
	return r.data.rdb.Del(ctx, r.getRedisRefreshToken(refreshToken)).Err()
}

// newAccessToken 新建 accessToken
func (r *jwtRepo) newAccessToken(ctx context.Context, claims biz.CustomClaims, expireAt time.Time) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expireAt),
	}

	token := jwt.NewWithClaims(getSignMethod(r.conf.SignMethod), claims)
	tokenStr, err := token.SignedString([]byte(r.conf.JwtKey))
	if err != nil {
		return "", err
	}
	// add token prefix
	return "Bearer " + tokenStr, nil
}

func getSignMethod(method string) jwt.SigningMethod {
	switch method {
	case "hs512":
		return jwt.SigningMethodHS512
	case "hs256":
		return jwt.SigningMethodHS256
	case "hs384":
		return jwt.SigningMethodHS384
	default:
		return jwt.SigningMethodHS256
	}
}

// newRefreshToken 新建 refreshToken
func (r *jwtRepo) newRefreshToken(ctx context.Context, userId uint64, expiration time.Duration) (string, error) {
	token := codec.UUIDHex()
	userIdStr := strconv.FormatUint(userId, 10)

	version, err := r.data.rdb.Get(ctx, r.getRedisRefreshTokenVersion(userIdStr)).Int64()
	if err != nil && err != redis.Nil {
		return "", err
	}
	err = r.data.rdb.Set(ctx, r.getRedisRefreshToken(token), r.getRedisRefreshTokenValue(userIdStr, version), expiration).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *jwtRepo) getRedisRefreshToken(token string) string {
	return "ias:refresh:" + token
}

func (r *jwtRepo) getRedisRefreshTokenVersion(userIdStr string) string {
	return "ias:refresh:version:" + userIdStr
}

func (r *jwtRepo) getRedisRefreshTokenValue(userIdStr string, version int64) string {
	return userIdStr + ":" + strconv.FormatInt(version, 10)
}

func (r *jwtRepo) parseRedisRefreshTokenValue(val string) (userIdStr string, version int64) {
	arr := strings.Split(val, ":")
	if len(arr) != 2 {
		return "", 0
	}
	version, _ = strconv.ParseInt(arr[1], 10, 64)
	return arr[0], version
}
