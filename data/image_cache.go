package data

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/conf"
	"gitlab.ctyuncdn.cn/ias/ias-core/mqtt"
	"net/http"
	neturl "net/url"
	"strconv"
	"time"
)

func NewImageCacheRepo(cfg *conf.Bootstrap, broker biz.MqttRepo, logger log.Logger) (biz.ImageCacheRepo, error) {
	switch cfg.Scene {
	case "jme":
		return NewImageCacheByApi(logger)
	case "jmv":
		return NewImageCacheByMqtt(broker, logger)
	}
	return nil, fmt.Errorf("newImageCacheRepo error, invalid scene: %s", cfg.Scene)
}

type imageCacheByApi struct {
	client    *transhttp.Client
	agentAddr *neturl.URL

	log *log.Helper
}

func NewImageCacheByApi(logger log.Logger) (biz.ImageCacheRepo, error) {
	conn, err := transhttp.NewClient(
		context.Background(),
		transhttp.WithTimeout(time.Minute),
		transhttp.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
		),
	)
	if err != nil {
		return nil, err
	}
	agentAddr, err := neturl.Parse("http://localhost:8082") // lite与agent在同一台宿主机上
	if err != nil {
		return nil, err
	}
	return &imageCacheByApi{
		client:    conn,
		agentAddr: agentAddr,
		log:       log.NewHelper(logger),
	}, nil
}

func (i *imageCacheByApi) generalRequest(ctx context.Context, url string, data []byte) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = i.client.Do(req)
	return err
}

func (i *imageCacheByApi) CacheImage(ctx context.Context, data *biz.ImageCache) error {
	url := i.agentAddr.JoinPath("/api/v1/image/cache")
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return i.generalRequest(ctx, url.String(), buf)
}

func (i *imageCacheByApi) UploadImage(ctx context.Context, data *biz.ImageUpload) error {
	url := i.agentAddr.JoinPath("/api/v1/image/upload")
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return i.generalRequest(ctx, url.String(), buf)
}

func (i *imageCacheByApi) ClearImage(ctx context.Context, data *biz.ImageClear) error {
	url := i.agentAddr.JoinPath("/api/v1/image/clear")
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return i.generalRequest(ctx, url.String(), buf)
}

type imageCacheByMqtt struct {
	mqtt biz.MqttRepo

	log *log.Helper
}

func NewImageCacheByMqtt(mqtt biz.MqttRepo, logger log.Logger) (biz.ImageCacheRepo, error) {
	return &imageCacheByMqtt{
		mqtt: mqtt,
		log:  log.NewHelper(logger),
	}, nil
}

func (i *imageCacheByMqtt) CacheImage(ctx context.Context, cache *biz.ImageCache) error {
	b, err := json.Marshal(cache)
	if err != nil {
		return err
	}
	deviceID, err := strconv.ParseUint(cache.DeviceID, 10, 64)
	if err != nil {
		return err
	}
	i.mqtt.Publish(mqtt.GetTopicImageCache(deviceID), 1, b)
	return nil
}

func (i *imageCacheByMqtt) UploadImage(ctx context.Context, cache *biz.ImageUpload) error {
	b, err := json.Marshal(cache)
	if err != nil {
		return err
	}
	deviceID, err := strconv.ParseUint(cache.DeviceID, 10, 64)
	if err != nil {
		return err
	}
	token := i.mqtt.Publish(mqtt.GetTopicImageUpload(deviceID), 1, b)
	token.Wait()
	return token.Error()
}

func (i *imageCacheByMqtt) ClearImage(ctx context.Context, cache *biz.ImageClear) error {
	b, err := json.Marshal(cache)
	if err != nil {
		return err
	}
	deviceID, err := strconv.ParseUint(cache.DeviceID, 10, 64)
	if err != nil {
		return err
	}
	token := i.mqtt.Publish(mqtt.GetTopicImageClear(deviceID), 2, b)
	token.Wait()
	return token.Error()
}
