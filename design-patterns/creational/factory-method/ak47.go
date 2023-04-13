package factorymethod

type Ak47 struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{
		Gun: Gun{
			Name:        "Ak47 machine gun",
			Power:       4,
			Description: "Shooting with power 100",
		},
	}
}
