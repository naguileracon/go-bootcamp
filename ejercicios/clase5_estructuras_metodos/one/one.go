package main

import (
	"errors"
	"fmt"
)

func main() {
	// inserting product
	productToSave := Product{
		ID:          3,
		Name:        "TV",
		Price:       "$1000",
		Description: "Panasonic TV",
		Category:    "Electronics",
	}
	err := productToSave.Save()
	if err != nil {
		println(err.Error())
		return
	}
	// getting all products
	Product.GetAll(productToSave)
	// finding product by id
	id := 3
	product, error := GetById(id)
	if error != "" {
		fmt.Printf("product with id %d not found", id)
		return
	}
	fmt.Println(product)
}

var Products = []Product{
	{
		ID:          1,
		Name:        "Laptop",
		Price:       "$1000",
		Description: "MacBook Pro",
		Category:    "Electronics",
	},
	{
		ID:          2,
		Name:        "Headphones",
		Price:       "$50",
		Description: "Sony Headphones",
		Category:    "Electronics",
	},
}

type Product struct {
	ID          int
	Name        string
	Price       string
	Description string
	Category    string
}

func (p Product) Save() (err error) {
	for _, product := range Products {
		if product.ID == p.ID {
			err = errors.New("product already exists")
			return
		}
	}
	Products = append(Products, p)
	return
}

func (p Product) GetAll() {
	fmt.Println(Products)
}

func GetById(id int) (result Product, err string) {
	for _, product := range Products {
		if product.ID == id {
			result = product
			return
		}
	}
	err = "product not found"
	return
}
