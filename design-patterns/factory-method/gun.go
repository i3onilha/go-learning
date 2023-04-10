package factorymethod

type Gun struct {
	Name        string
	Power       int
	Description string
}

func (g *Gun) SetName(name string) {
	g.Name = name
}

func (g *Gun) GetName() string {
	return g.Name
}

func (g *Gun) SetPower(power int) {
	g.Power = power
}

func (g *Gun) GetPower() int {
	return g.Power
}

func (g *Gun) GetGunDescription() string {
	return g.Description
}
