package factorymethod

import "testing"

func TestIGun(t *testing.T) {
	scenario := []struct {
		name   string
		power  int
		result string
	}{
		{"ak47", 100, "Shooting with power 100"},
		{"musket", 1, "Shooting with power 1"},
	}
	for _, s := range scenario {
		gun, err := GetGun(s.name)
		if err != nil {
			t.Errorf("Expected nil, got %s", err)
			return
		}
		if gun.GetGunDescription() != s.result {
			t.Errorf("Expected %s, got %s", s.result, gun.GetGunDescription())
		}
	}
}
