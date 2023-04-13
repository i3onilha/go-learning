package adapter

import "testing"

func TestAdapter(t *testing.T) {
	computer := &Computer{}
	scenarios := []struct {
		computer *Computer
		device   USBDevice
		expected bool
	}{
		{
			computer: computer,
			device:   &PenDrive{},
			expected: true,
		},
		{
			computer: computer,
			device: &P2M2USBAdapter{
				p2m: &Headset{},
			},
			expected: true,
		},
		{
			computer: computer,
			device:   &P2M2USBAdapter{},
			expected: false,
		},
	}
	for _, s := range scenarios {
		if s.computer.PluginUSBPort(s.device) != s.expected {
			t.Errorf("Expected %v, got %v", s.expected, s.computer.PluginUSBPort(s.device))
		}
	}
}
