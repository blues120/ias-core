package sophgo

import (
	"context"
	"encoding/json"
	"fmt"
	coreMqtt "github.com/blues120/ias-core/mqtt"
	neturl "net/url"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/conf"
)

const (
	UriTaskStart = "/api/v1/task/start"
	UriTaskStop  = "/api/v1/task/stop"
)

type scheduler struct {
	bc      *conf.Bootstrap
	baseUrl *neturl.URL // 算能算法服务地址

	log *log.Helper
}

func NewScheduler(bc *conf.Bootstrap, logger log.Logger) biz.SchedulerRepo {
	if bc == nil || bc.Sophgo == nil || logger == nil {
		return nil
	}

	// 解析算能算法服务地址
	url, err := neturl.Parse(bc.Sophgo.Addr)
	if err != nil {
		return nil
	}

	return &scheduler{
		bc:      bc,
		baseUrl: url,
		log:     log.NewHelper(logger),
	}
}

// Close 关闭
func (s *scheduler) Close() error {
	return nil
}

// Start 启动算能算法
func (s *scheduler) Start(ctx context.Context, ta *biz.Task) error {
	url := s.baseUrl.JoinPath(UriTaskStart).String()

	// 生成算能算法配置
	algoConfig := TaskToAlgoConfig(s.bc, ta)
	data, _ := json.Marshal(&algoConfig)
	s.log.Debugf("【算能任务启动】算法配置: %s", string(data))

	resp, err := PostRequest(url, data)
	if err != nil {
		s.log.Errorf("【算能任务启动】POST请求失败: %v", err)
		return err
	} else if resp.Code != 0 {
		err = fmt.Errorf("【算能任务启动】算能端报错 %v", resp.Msg)
		s.log.Error(err)
		return err
	}

	s.log.Infof("【算能任务启动】成功")
	return nil
}

// Stop 停止算能算法
func (s *scheduler) Stop(ctx context.Context, ta *biz.Task) error {
	url := s.baseUrl.JoinPath(UriTaskStop).String()

	stopParams := coreMqtt.TaskStop{
		TaskID: fmt.Sprintf("%d", ta.Id),
		Algorithm: coreMqtt.Algorithm{
			Prefix: ta.Algo.Algorithm.Prefix,
		},
	}
	data, _ := json.Marshal(&stopParams)

	resp, err := PostRequest(url, data)
	if err != nil {
		s.log.Errorf("【算能任务停止】POST请求失败: %v", err)
		return err
	} else if resp.Code != 0 {
		err = fmt.Errorf("【算能任务停止】算能端报错 %v", resp.Msg)
		return err
	}

	return nil
}

func (s *scheduler) GetLog(ctx context.Context, ta *biz.Task, conn *websocket.Conn) error {
	// TODO implement
	panic("implement me")
}

func (s *scheduler) GetStatuses(ctx context.Context, taskIdList []uint) ([]biz.TaskIdStatus, error) {
	// TODO implement
	panic("implement me")
}
