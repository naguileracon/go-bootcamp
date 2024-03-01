package products

type SmallProduct struct {
	productPrice float64
}

func (sp SmallProduct) Price() float64 {
	return sp.productPrice
}

type MediumProduct struct {
	productPrice float64
}

func (sp MediumProduct) Price() float64 {
	return sp.productPrice
}

type LargeProduct struct {
	productPrice float64
}

func (sp LargeProduct) Price() float64 {
	return sp.productPrice
}
