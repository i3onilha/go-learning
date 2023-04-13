package adapter

type P2M2USBAdapter struct {
	p2m P2MDevice
}

func NewP2M2USBAdapter(dev P2MDevice) *P2M2USBAdapter {
	if dev.InsertIntoP2FPort() == P2M {
		return &P2M2USBAdapter{
			p2m: dev,
		}
	}
	return nil
}

func (p *P2M2USBAdapter) InsertIntoUSBPort() string {
	if p.p2m == nil {
		return P2MErr
	}
	return USB
}
