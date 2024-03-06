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

func (mp MediumProduct) Price() float64 {
	price := mp.productPrice + (mp.productPrice + 0.03)
	return price
}

type LargeProduct struct {
	productPrice float64
}

func (lp LargeProduct) Price() float64 {
	price := lp.productPrice + (lp.productPrice + 0.06) + float64(2500)
	return price
}
