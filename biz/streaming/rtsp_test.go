//go:build skip
// +build skip

package streaming

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// 执行单元测试前，请确保这个流是在线的
const onlineRtspUrl = "rtsp://127.0.0.1:8554/qqq"

func TestNewRtsp(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{name: "string", args: args{addr: "S/XGT}Lb<x%;2{RU"}, wantErr: ErrInvalidRtspAddr},
		{name: "rtmp", args: args{addr: "rtmp://127.0.0.1:8554/03261"}, wantErr: ErrInvalidRtspAddr},
		{name: "http", args: args{addr: "http://127.0.0.1:8554/03261"}, wantErr: ErrInvalidRtspAddr},
		{name: "rtsp", args: args{addr: "rtsp://127.0.0.1:8554/03261"}, wantErr: nil},
		{name: "rtsp with auth", args: args{addr: "rtsp://user:password@44.178.93.38/media/video1"}, wantErr: nil},
		{name: "online rtsp", args: args{addr: onlineRtspUrl}, wantErr: nil},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := NewRtsp(tt.args.addr)
			require.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRtspIsOnline(t *testing.T) {
	type args struct {
		ctx     context.Context
		addr    string
		timeout time.Duration
	}
	tests := []struct {
		name       string
		args       args
		wantNoErr  bool
		wantOnline bool
	}{
		{
			name:       "offline(non-existent)",
			args:       args{ctx: context.Background(), addr: "rtsp://127.0.0.1:8554/non-existent", timeout: time.Second * 10},
			wantNoErr:  false,
			wantOnline: false,
		},
		{
			name:       "online",
			args:       args{ctx: context.Background(), addr: onlineRtspUrl, timeout: time.Second * 10},
			wantNoErr:  true,
			wantOnline: true,
		},
		{
			name:       "network unreachable",
			args:       args{ctx: context.Background(), addr: "rtsp://111.11.111.111:8554/non-existent", timeout: time.Second * 10},
			wantNoErr:  false,
			wantOnline: false,
		},
		{
			name:       "with authentication",
			args:       args{ctx: context.Background(), addr: "rtsp://user:password@192.168.1.62:554", timeout: time.Second * 3},
			wantNoErr:  true,
			wantOnline: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			rtsp, err := NewRtsp(tt.args.addr)
			require.NoError(t, err)
			got, err := rtsp.IsOnline(tt.args.ctx, tt.args.timeout)
			require.Equal(t, tt.wantNoErr, err == nil)
			require.Equal(t, tt.wantOnline, got)
		})
	}
}

func TestRtspGetSnapshot(t *testing.T) {
	type args struct {
		ctx  context.Context
		addr string
	}
	tests := []struct {
		name      string
		args      args
		wantNoErr bool
		wantEmpty bool
	}{
		{name: "offline(non-existent) rtsp", args: args{ctx: context.Background(), addr: "rtsp://127.0.0.1:8554/non-existent"}, wantNoErr: false, wantEmpty: true},
		{name: "online rtsp", args: args{ctx: context.Background(), addr: onlineRtspUrl}, wantNoErr: true, wantEmpty: false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			rtsp, err := NewRtsp(tt.args.addr)
			require.NoError(t, err)
			got, err := rtsp.GetSnapshot(tt.args.ctx)
			require.Equal(t, tt.wantNoErr, err == nil)
			require.Equal(t, tt.wantEmpty, len(got) == 0)
		})
	}
}

func TestExtractInfo(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected *Info
	}{
		{"error1", []byte(`{
			interleave TCP based rtp used: 0, 1
			BMvidDecCreateW5 board id 0 coreid 0
			libbmvideo.so addr : /system/lib/libbmvideo.so, name_len: 12
			vpu firmware addr: /system/lib/vpu_firmware/chagall_dec.bin
			VERSION=0, REVISION=213135
				"programs": [

				],
				"streams": [
					{

						"width": 1440,
						"height": 1080
					}
				]
			[ERROR] Failed to DEC_PIC_HDR(ERROR REASON: 00005000) error code is 0x1
			InstIdx 0: BMVidDecSeqInitW5 failed Error code is 0x1
		}`), nil},
		{"correctextra", []byte(`{
interleave TCP based rtp used: 0, 1
BMvidDecCreateW5 board id 0 coreid 0
libbmvideo.so addr : /system/lib/libbmvideo.so, name_len: 12
vpu firmware addr: /system/lib/vpu_firmware/chagall_dec.bin
VERSION=0, REVISION=213135
    "streams": [
        {
            "index": 0,
            "codec_name": "h264",
            "codec_long_name": "H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10",
            "profile": "High",
            "codec_type": "video",
            "codec_time_base": "1/24",
            "codec_tag_string": "[0][0][0][0]",
            "codec_tag": "0x0000",
            "width": 1920,
            "height": 1080,
            "coded_width": 1920,
            "coded_height": 1088,
            "has_b_frames": 2,
            "sample_aspect_ratio": "1:1",
            "display_aspect_ratio": "16:9",
            "pix_fmt": "yuv420p",
            "level": 40,
            "chroma_location": "left",
            "field_order": "progressive",
            "refs": 1,
            "r_frame_rate": "24/1",
            "avg_frame_rate": "24/1",
            "time_base": "1/90000",
            "start_pts": 15000,
            "start_time": "0.166667",
            "bits_per_raw_sample": "8",
            "disposition": {
                "default": 0,
                "dub": 0,
                "original": 0,
                "comment": 0,
                "lyrics": 0,
                "karaoke": 0,
                "forced": 0,
                "hearing_impaired": 0,
                "visual_impaired": 0,
                "clean_effects": 0,
                "attached_pic": 0,
                "timed_thumbnails": 0
            }
        }
[ERROR] Failed to DEC_PIC_HDR(ERROR REASON: 00005000) error code is 0x1
InstIdx 0: BMVidDecSeqInitW5 failed Error code is 0x1
    ]
}`), &Info{CodecName: "h264", Width: int32(1920), Height: int32(1080)}},

		//{"EmptyInput", []byte(""), &Info{}}, // empty input should return an empty Info struct with no error
		//{"NullInput", nil, &Info{}},         // null input should return an empty Info struct with no error (Go does not allow passing a nil slice to a function that expects a byte slice)
	}
	for _, testCase := range testCases {
		result, _ := extractInfo(testCase.input)
		fmt.Println(result)
		fmt.Println(testCase.expected)

		if result != testCase.expected && (result == nil || (&Info{}) == testCase.expected) {
			t.Errorf("%s: expected %v, got %v", testCase.name, testCase.expected, result)
			continue
		}

	}

}
