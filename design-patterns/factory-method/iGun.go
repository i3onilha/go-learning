package factorymethod

type IGun interface {
	SetName(name string)
	GetName() string
	SetPower(power int)
	GetPower() int
	GetGunDescription() string
}
