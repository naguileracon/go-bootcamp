package main

import (
	"two/products"
)

func main() {
	// small product
	smallProduct, err := products.GetProduct("small", 6000.5)
	if err != nil {
		println(err.Error())
		return
	}
	println(smallProduct.Price())
	// medium product
	mediumProduct, err := products.GetProduct("medium", 6000.5)
	if err != nil {
		println(err.Error())
		return
	}
	println(mediumProduct.Price())
	// large product
	largeProduct, err := products.GetProduct("large", 6000.5)
	if err != nil {
		println(err.Error())
		return
	}
	println(largeProduct.Price())
	// not found product
	product, err := products.GetProduct("great", 6000.5)
	if err != nil {
		println(err.Error())
		return
	}
	println(product.Price())
}
