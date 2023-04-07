package abstractfactory

type ISportsFactory interface {
	makeShoe(logo string, size int) IShoe
	makeShirt(logo, size string) IShirt
}

func GetSportsFactory(brand string) ISportsFactory {
	if brand == "adidas" {
		return &adidasFactory{}
	} else if brand == "nike" {
		return &nikeFactory{}
	}
	return nil
}
