package streaming

type onvif struct {
	*rtsp
}

func NewOnvif(addr string) (Protocol, error) {
	pt, err := NewRtsp(addr)
	if err != nil {
		return nil, err
	}
	return &onvif{pt.(*rtsp)}, nil
}

func (r *onvif) Type() ProtocolType {
	return ProtocolTypeOnvif
}
