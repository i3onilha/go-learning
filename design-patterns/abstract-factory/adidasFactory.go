package abstractfactory

type adidasFactory struct{}

func (a *adidasFactory) makeShoe(logo string, size int) IShoe {
	return &adidasShoe{
		Shoe: Shoe{
			logo: logo,
			size: size,
		},
	}
}

func (a *adidasFactory) makeShirt(logo, size string) IShirt {
	return &adidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: size,
		},
	}
}
