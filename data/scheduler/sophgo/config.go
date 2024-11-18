package sophgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/conf"
	coreMqtt "github.com/blues120/ias-core/mqtt"
)

func TaskToAlgoConfig(bc *conf.Bootstrap, ta *biz.Task) coreMqtt.TaskStart {
	algoConfigJson := ta.AlgoConfig
	algoConfigMap := make(map[string]interface{})
	json.Unmarshal([]byte(algoConfigJson), &algoConfigMap)
	confidence := 0.25
	if targetDetect, ok := algoConfigMap["targetDetect"]; ok {
		if targetDetect != nil {
			confidence = 0.01 * targetDetect.(float64)
		}
	}

	// telestream api算法特有配置
	customConfig := make(map[string]interface{})
	if val, ok := algoConfigMap["customConfig"]; ok {
		if val2, ok2 := val.(map[string]interface{}); ok2 {
			customConfig = val2
		}
	}

	var period uint32
	period, err := biz.ParsePeriod(ta.AlgoConfig)
	if err != nil {
		period = 0
	}

	if period > 0 {
		customConfig["warn_time_interval"] = period
	}

	controlTime, err := biz.ParseControlTime(ta.AlgoConfig)
	if err != nil {
		controlTime = ""
	}

	config := coreMqtt.TaskStart{
		Agent: coreMqtt.Agent{
			Server: coreMqtt.Server{
				File: coreMqtt.File{
					DataOutputDir: "/data/data",
					MetaOutputDir: "/data/meta",
				},
			},
		},
		Algorithm: coreMqtt.Algorithm{
			Confidence:   confidence,
			Config:       algoConfigMap,
			ID:           ta.Algo.Algorithm.AlgoNameEn,
			Mode:         "realtime",
			Period:       int64(period),
			Trigger:      1,
			Version:      ta.Algo.Algorithm.Version,
			Prefix:       ta.Algo.Algorithm.Prefix,
			CustomConfig: customConfig,
		},
		CallbackURL: fmt.Sprintf("http://%s", bc.Callback.Addr),
		Task: coreMqtt.Task{
			TaskID:        fmt.Sprintf("%d", ta.Id),
			WorkspaceID:   "",
			Period:        strconv.FormatUint(uint64(period), 10),
			ControlPeriod: controlTime,
			CallbackURL:   fmt.Sprintf("http://%s", bc.Callback.Addr),
			AlgoConfig:    ta.AlgoConfig,
			Type:          ta.Type,
		},
		Version: "v1.1",
	}

	if ta.Algo.Algorithm.AlgoID != "" {
		config.Algorithm.ModelNames = []string{ta.Algo.Algorithm.AlgoID}
	}

	cameraMultiImgBox := make(map[uint64]string, 0)
	for _, ca := range ta.Cameras {
		if ca.MultiImgBox != "" && ca.MultiImgBox != "[]" { // 多区域场景
			cameraMultiImgBox[ca.Camera.Id] = ca.MultiImgBox
		}
	}

	sources := make([]coreMqtt.Source, 0, len(ta.Cameras))
	for _, tc := range ta.Cameras {
		camera := tc.Camera
		codec := camera.StreamingInfo.CodecName
		width := int64(camera.StreamingInfo.Width)
		height := int64(camera.StreamingInfo.Height)
		item := coreMqtt.Source{
			CameraID:          fmt.Sprintf("%d", camera.Id),
			Codec:             &codec,
			DropFrameInterval: int64(ta.Algo.Interval),
			Source:            camera.StreamingProtocol.Source(),
			Width:             &width,
			Height:            &height,
			Events:            make([]coreMqtt.Event, 0),
			Areas:             make([]coreMqtt.Area, 0),
			Lines:             make([]coreMqtt.Line, 0),
		}
		if strings.HasPrefix(item.Source, "rtsp") {
			item.SourceType = "RTSP"
		} else {
			item.SourceType = "VIDEO"
		}
		if multiImgBox, has := cameraMultiImgBox[camera.Id]; has {
			if multiImgBox != "" {
				item.Events = ParseEvents(ta.Algo.Algorithm.Provider, multiImgBox, uint(*item.Width), uint(*item.Height))
				item.Areas, item.Lines = ParseAreaLines(multiImgBox)
			}
		}
		sources = append(sources, item)
	}
	config.Sources = sources
	sourceMarshal, _ := json.Marshal(sources)
	config.Task.Sources = string(sourceMarshal)
	return config
}

func ParseEvents(provider string, eventStr string, width, height uint) []coreMqtt.Event {
	ret := make([]coreMqtt.Event, 0)

	if eventStr == "" {
		return ret
	}
	// 解析原始输入数据结构
	var input map[string][][]struct {
		X float64
		Y float64
	}
	if err := jsoniter.Unmarshal([]byte(eventStr), &input); err != nil {
		return ret
	}

	detectDirections := make([]int64, 0)
	targetLine := make([]int64, 0)
	targetDirection := make([]int64, 0)

	for k, v := range input {
		if len(v) == 0 {
			continue
		}
		if k == "targetArea" || k == "excludeArea" {
			var detectAreas []int64
			detectAreasShape := []int64{0}
			detectAreasShapeItemNum := 0
			for _, item := range v {
				for _, c := range item {
					detectAreas = append(detectAreas, int64(c.X*float64(width)))
					detectAreas = append(detectAreas, int64(c.Y*float64(height)))
				}
				detectAreasShapeItemNum++
				detectAreasShape = append(detectAreasShape, int64(len(item)))
			}
			detectAreasShape[0] = int64(detectAreasShapeItemNum)
			event := coreMqtt.Event{
				DetectAreas:      detectAreas,
				DetectAreasShape: detectAreasShape,
			}
			if k == "targetArea" {
				name := "in-filter"
				event.Name = &name
			} else {
				name := "out-filter"
				event.Name = &name
			}
			ret = append(ret, event)
		} else if k == "targetLine" {
			for _, item := range v {
				for _, c := range item {
					targetLine = append(targetLine, int64(c.X*float64(width)))
					targetLine = append(targetLine, int64(c.Y*float64(height)))
				}
			}
		} else if k == "targetDirection" {
			for _, item := range v {
				for _, c := range item {
					targetDirection = append(targetDirection, int64(c.X*float64(width)))
					targetDirection = append(targetDirection, int64(c.Y*float64(height)))
				}
			}
		}
	}

	// 说明是画线和方向的人流量检测算法
	if len(targetLine) > 0 {
		switch provider {
		case "sophgo_city", "sophgo_park", "ctyun_telestream":
			detectDirectionsCount := int64(len(targetDirection) / 4)
			detectLinesCount := int64(len(targetLine) / 4)
			event := coreMqtt.Event{
				DetectDirections:      targetDirection,
				DetectDirectionsCount: &detectDirectionsCount,
				DetectLines:           targetLine,
				DetectLinesCount:      &detectLinesCount,
			}
			ret = append(ret, event)
		case "ctyun_ias":
			detectDirections = append(detectDirections, targetDirection...)
			detectDirections = append(detectDirections, targetLine...)
			event := coreMqtt.Event{
				DetectDirections: detectDirections,
			}
			ret = append(ret, event)
		}
	}
	return ret
}

// ParseAreaLines telestream 画区域画线规范
func ParseAreaLines(eventStr string) ([]coreMqtt.Area, []coreMqtt.Line) {
	areas := make([]coreMqtt.Area, 0)
	lines := make([]coreMqtt.Line, 0)

	if eventStr == "" {
		return areas, lines
	}
	// 解析原始输入数据结构
	var input map[string][][]struct {
		X float64
		Y float64
	}
	if err := jsoniter.Unmarshal([]byte(eventStr), &input); err != nil {
		return areas, lines
	}
	// {"targetArea":[[{"x":0.39,"y":0.21},{"x":0.37,"y":0.95},{"x":0.85,"y":0.91},{"x":0.83,"y":0.21}]],"excludeArea":[]}
	// 二维数组转化为一维数组
	if v, ok := input["targetArea"]; ok && len(v) > 0 {
		for index, item := range v {
			area := coreMqtt.Area{
				ID:    strconv.Itoa(index),
				Coord: make([]float64, 0),
			}
			for _, c := range item {
				area.Coord = append(area.Coord, c.X, c.Y)
			}
			areas = append(areas, area)
		}
	}
	// {"targetLine":[[{"x":0.07,"y":0.61},{"x":0.42,"y":0.58}]],"targetDirection":[[{"x":0.21,"y":0.82},{"x":0.24,"y":0.41}]]}
	// 二维数组转化为一维数组
	if v, ok := input["targetLine"]; ok && len(v) > 0 {
		for index, item := range v {
			line := coreMqtt.Line{
				ID:        strconv.Itoa(index),
				Coord:     make([]float64, 0),
				Direction: make([]float64, 0),
			}
			for _, c := range item {
				line.Coord = append(line.Coord, c.X, c.Y)
			}
			lines = append(lines, line)
		}
	}
	if v, ok := input["targetDirection"]; ok && len(v) > 0 && len(v) <= len(lines) {
		for index, item := range v {
			for _, c := range item {
				lines[index].Direction = append(lines[index].Direction, c.X, c.Y)
			}
		}
	}

	return areas, lines
}

// PostRequest 构造http POST请求
func PostRequest(url string, bodyData []byte) (*Response, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, fmt.Errorf("HTTP POST ERROR: %v", err)
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("RESPONSE READ ERROR: %v", err)
	}
	var resp2 *Response
	err = jsoniter.Unmarshal(respData, &resp2)
	if err != nil {
		return nil, fmt.Errorf("RESPONSE UNMARSHAL ERROR: %v", err)
	}

	return resp2, nil
}

type Response struct {
	Code int                    `json:"code"` // 接口调用的状态码，0 表示成功, 非 0 表示失败
	Data map[string]interface{} `json:"data"` // 接口调用返回的具体数据
	Msg  string                 `json:"msg"`  // 接口调用结果的文本描述
}
