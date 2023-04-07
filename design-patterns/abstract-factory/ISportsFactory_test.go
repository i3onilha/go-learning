package abstractfactory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestISportsFactory(t *testing.T) {
	adidasFactory, nikeFactory := GetSportsFactory("adidas"), GetSportsFactory("nike")
	adidasShoe := adidasFactory.makeShoe("adidas", 41)
	adidasShirt := adidasFactory.makeShirt("adidas", "L")
	nikeShoe := nikeFactory.makeShoe("nike", 42)
	nikeShirt := nikeFactory.makeShirt("nike", "M")
	assert.Equal(t, "adidas", adidasShoe.getLogo())
	assert.Equal(t, "adidas", adidasShirt.getLogo())
	assert.Equal(t, "nike", nikeShoe.getLogo())
	assert.Equal(t, "nike", nikeShirt.getLogo())
	assert.Equal(t, 41, adidasShoe.getSize())
	assert.Equal(t, 42, nikeShoe.getSize())
	assert.Equal(t, "L", adidasShirt.getSize())
	assert.Equal(t, "M", nikeShirt.getSize())
}
