package builder

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBulder() *IglooBuilder {
	return &IglooBuilder{}
}

func (i *IglooBuilder) setWindowType() {
	i.windowType = "Snow window"
}

func (i *IglooBuilder) setDoorType() {
	i.doorType = "Snow door"
}

func (i *IglooBuilder) setNumFloor() {
	i.floor = 1
}

func (i *IglooBuilder) getHouse() House {
	return House{
		windowType: i.windowType,
		doorType:   i.doorType,
		floor:      i.floor,
	}
}
