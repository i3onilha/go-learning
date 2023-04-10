package factorymethod

type Musket struct {
	Gun
}

func NewMusket() IGun {
	return &Musket{
		Gun{
			Name:        "Musket pistol",
			Power:       1,
			Description: "Shooting with power 1",
		},
	}
}
