package builder

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBulder()
	}
	if builderType == "igloo" {
		return newIglooBulder()
	}
	return nil
}
