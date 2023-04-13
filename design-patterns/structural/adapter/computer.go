package adapter

type Computer struct{}

func (m *Computer) PluginUSBPort(dev USBDevice) bool {
	return dev.InsertIntoUSBPort() == USB
}
