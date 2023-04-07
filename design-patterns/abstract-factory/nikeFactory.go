package abstractfactory

type nikeFactory struct{}

func (n *nikeFactory) makeShoe(logo string, size int) IShoe {
	return &nikeShoe{
		Shoe: Shoe{
			logo: logo,
			size: size,
		},
	}
}

func (n *nikeFactory) makeShirt(logo, size string) IShirt {
	return &nikeShirt{
		Shirt: Shirt{
			logo: logo,
			size: size,
		},
	}
}
