package products

import "fmt"

func GetProduct(productType string, price float64) (IProduct, error) {
	if productType == "small" {
		return SmallProduct{price}, nil
	}
	if productType == "medium" {
		return MediumProduct{price}, nil
	}
	if productType == "large" {
		return LargeProduct{price}, nil
	}
	return nil, fmt.Errorf("Wrong product type")
}
