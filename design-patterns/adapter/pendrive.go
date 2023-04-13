package adapter

type PenDrive struct{}

func (p *PenDrive) InsertIntoUSBPort() string {
	return USB
}
