package streaming

type Info struct {
	CodecName string `json:"codec_name"` // 编码名称
	Width     int32  `json:"width"`      // 宽
	Height    int32  `json:"height"`     // 高
}
