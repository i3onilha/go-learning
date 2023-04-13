package builder

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

var builderTypes = map[string]IBuilder{
	"normal": newNormalBulder(),
	"igloo":  newIglooBulder(),
}

func GetBuilder(builderType string) IBuilder {
	if builder, ok := builderTypes[builderType]; ok {
		return builder
	}
	return nil
}
