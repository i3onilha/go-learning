package adapter

const (
	USB = "USB Driver Plugged In"
	P2M = "P2 Male Driver Plugged In"
	P2F = "P2 Female Driver Plugged In"

	P2MErr = "No P2M Device Plugged In"
)

type USBDevice interface {
	InsertIntoUSBPort() string
}

type P2MDevice interface {
	InsertIntoP2FPort() string
}
