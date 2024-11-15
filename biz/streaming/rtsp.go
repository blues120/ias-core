package streaming

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
	customerr "gitlab.ctyuncdn.cn/ias/ias-core/errors"
)

var (
	ErrInvalidRtspAddr = customerr.ErrorInvalidRtspAddr("rtsp 地址不合法")
)

type rtsp struct {
	source *url.URL
}

func NewRtsp(addr string) (Protocol, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, ErrInvalidRtspAddr
	}
	if u.Scheme != string(ProtocolTypeRtsp) {
		return nil, ErrInvalidRtspAddr
	}
	return &rtsp{source: u}, nil
}

func (r *rtsp) Source() string {
	spSource, _ := url.QueryUnescape(r.source.String()) // 转成原始字符串
	return spSource
}

func (r *rtsp) Type() ProtocolType {
	return ProtocolTypeRtsp
}

func (r *rtsp) IsOnline(ctx context.Context, timeout time.Duration) (bool, error) {
	cmd := exec.Command("timeout",
		fmt.Sprintf("%vs", timeout.Seconds()),
		"ffprobe",
		"-rtsp_transport",
		"tcp",
		r.Source(),
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, fmt.Errorf(string(output))
	}
	return true, nil
}

func (r *rtsp) GetSnapshot(ctx context.Context) ([]byte, error) {
	cmd := exec.Command("ffmpeg",
		"-y",
		"-rtsp_transport",
		"tcp",
		"-i",
		r.Source(),
		"-vframes",
		"1",
		"-f",
		"image2",
		"pipe:",
	)

	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (r *rtsp) GetInfo(ctx context.Context) (*Info, error) {
	cmd := exec.Command("ffprobe",
		"-v",
		"error",
		"-show_streams",
		"-select_streams",
		"v:0",
		"-print_format",
		"json",
		"-rtsp_transport",
		"tcp",
		"-i",
		r.Source(),
	)

	buf, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var info struct {
		Streams []*Info `json:"streams"`
	}

	if err = jsoniter.Unmarshal(buf, &info); err != nil {

		//不正确时buf返回以下数据：
		//buf = []byte(`{
		//	interleave TCP based rtp used: 0, 1
		//	BMvidDecCreateW5 board id 0 coreid 0
		//	libbmvideo.so addr : /system/lib/libbmvideo.so, name_len: 12
		//	vpu firmware addr: /system/lib/vpu_firmware/chagall_dec.bin
		//	VERSION=0, REVISION=213135
		//		"programs": [
		//
		//		],
		//		"streams": [
		//			{
		//				"codec_name": "h264",
		//				"width": 1440,
		//				"height": 1080
		//			}
		//	[ERROR] Failed to DEC_PIC_HDR(ERROR REASON: 00005000) error code is 0x1
		//	InstIdx 0: BMVidDecSeqInitW5 failed Error code is 0x1
		//		]
		//}`)

		//将buf中的codec_name width height正则化为Info结构体的格式并返回
		return extractInfo(buf)
	}

	if len(info.Streams) > 0 {
		return info.Streams[0], nil
	}

	return nil, customerr.ErrorStreamInfoNotfound("未找到流信息")
}

func (r *rtsp) GetInfoWithTimeout(ctx context.Context, timeout time.Duration) (*Info, error) {
	// 创建一个带有超时的context
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx,
		"ffprobe",
		"-v",
		"error",
		"-show_streams",
		"-select_streams",
		"v:0",
		"-print_format",
		"json",
		"-rtsp_transport",
		"tcp",
		"-i",
		r.Source(),
	)

	buf, err := cmd.Output()
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return nil, fmt.Errorf("获取视频流信息超时: %w", ctx.Err())
	}
	if err != nil {
		return nil, err
	}

	var info struct {
		Streams []*Info `json:"streams"`
	}

	if err = jsoniter.Unmarshal(buf, &info); err != nil {
		//将buf中的codec_name width height正则化为Info结构体的格式并返回
		return extractInfo(buf)
	}

	if len(info.Streams) > 0 {
		return info.Streams[0], nil
	}

	return nil, customerr.ErrorStreamInfoNotfound("未找到流信息")
}

func extractInfo(buf []byte) (*Info, error) {
	// 1.使用正则表达式匹配并提取 codec_name、width 和 height
	codecNamePattern := `"codec_name"\s*:\s*"([^"]+)"`
	widthPattern := `"width"\s*:\s*(\d+)`
	heightPattern := `"height"\s*:\s*(\d+)`
	reCodecName := regexp.MustCompile(codecNamePattern)
	reWidth := regexp.MustCompile(widthPattern)
	reHeight := regexp.MustCompile(heightPattern)
	// 使用正则表达式来提取匹配的值
	data := string(buf)
	codecNameMatches := reCodecName.FindStringSubmatch(data)
	widthMatches := reWidth.FindStringSubmatch(data)
	heightMatches := reHeight.FindStringSubmatch(data)

	// 2.提取匹配到的信息
	if len(codecNameMatches) != 2 || len(widthMatches) != 2 || len(heightMatches) != 2 {
		return nil, customerr.ErrorStreamInfoNotfound("未找到流信息")
	}
	codecName := codecNameMatches[1]
	widthStr := widthMatches[1]
	heightStr := heightMatches[1]

	// 3.将提取的字符串转换为整数
	width, err := strconv.Atoi(widthStr)
	if err != nil {
		return nil, fmt.Errorf("宽度(width)解析失败： %w", err)
	}

	height, err := strconv.Atoi(heightStr)
	if err != nil {
		return nil, fmt.Errorf("高度(height)解析失败： %w", err)
	}
	// 4. 将提取的信息填充到 Info 结构体中
	info := Info{
		CodecName: codecName,
		Width:     int32(width),
		Height:    int32(height),
	}
	//5.返回结构体
	return &info, nil
}
