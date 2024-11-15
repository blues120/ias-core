package streaming

type offlineVideo struct {
	*rtsp
}

func NewOfflineVideo(addr string) (Protocol, error) {
	pt, err := NewRtsp(addr)
	if err != nil {
		return nil, err
	}
	return &offlineVideo{pt.(*rtsp)}, nil
}

func (r *offlineVideo) Type() ProtocolType {
	return ProtocolTypeOfflineVideo
}
