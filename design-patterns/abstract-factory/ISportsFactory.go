package abstractfactory

type ISportsFactory interface {
	makeShoe(logo string, size int) IShoe
	makeShirt(logo, size string) IShirt
}

var brands = map[string]ISportsFactory{
	"adidas": &adidasFactory{},
	"nike":   &nikeFactory{},
}

func GetSportsFactory(brand string) ISportsFactory {
	if brands[brand] == nil {
		return nil
	}
	return brands[brand]
}
