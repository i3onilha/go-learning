package factorymethod

import "fmt"

var gunTypes = map[string]IGun{
	"ak47":   NewAk47(),
	"musket": NewMusket(),
}

func GetGun(gunType string) (IGun, error) {
	if gun, ok := gunTypes[gunType]; ok {
		return gun, nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
