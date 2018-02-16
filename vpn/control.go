package vpn

type Control struct {
}

type ByteCountFunc func(user string, downKiBs, upKiBs int)

func NewControl(conf *Config, byteCount ByteCountFunc) (*Control, error) {
	return nil, nil
}

func (c *Control) Close() {

}

func (c *Control) Kill(user string) {

}
