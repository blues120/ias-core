package streaming

type gb28181 struct {
	*rtsp
}

func NewGb28181(addr string) (Protocol, error) {
	pt, err := NewRtsp(addr)
	if err != nil {
		return nil, err
	}
	return &gb28181{pt.(*rtsp)}, nil
}

func (r *gb28181) Type() ProtocolType {
	return ProtocolTypeGB28181
}
