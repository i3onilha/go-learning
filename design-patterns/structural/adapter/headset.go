package adapter

type Headset struct{}

func (h *Headset) InsertIntoP2FPort() string {
	return P2M
}
