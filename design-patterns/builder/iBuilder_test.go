package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHousesBuilder(t *testing.T) {
	normalBuilder := GetBuilder("normal")
	iglooBuilder := GetBuilder("igloo")

	director := NewDirector(normalBuilder)

	normalHouse := director.BuildHouse()
	assert.Equal(t, "Wooden window", normalHouse.windowType)
	assert.Equal(t, "Wooden door", normalHouse.doorType)
	assert.Equal(t, 2, normalHouse.floor)

	director.SetBuilder(iglooBuilder)

	iglooHouse := director.BuildHouse()
	assert.Equal(t, "Snow window", iglooHouse.windowType)
	assert.Equal(t, "Snow door", iglooHouse.doorType)
	assert.Equal(t, 1, iglooHouse.floor)
}
