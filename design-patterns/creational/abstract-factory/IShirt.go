package abstractfactory

type IShirt interface {
	setLogo(logo string)
	setSize(size string)
	getLogo() string
	getSize() string
}

type Shirt struct {
	logo string
	size string
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) setSize(size string) {
	s.size = size
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) getSize() string {
	return s.size
}
