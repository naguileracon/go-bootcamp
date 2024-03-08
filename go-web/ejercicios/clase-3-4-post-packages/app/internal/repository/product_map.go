package repository

import (
	"app/internal"
)

func NewProductMap(db map[int]internal.Product, lastID int) *ProductMap {
	if db == nil {
		db = make(map[int]internal.Product)
	}

	return &ProductMap{
		products: db,
		lastId:   lastID,
	}
}

type ProductMap struct {
	products map[int]internal.Product
	lastId   int
}

func (pm *ProductMap) Save(product *internal.Product) (err error) {
	// validate if code_value already exists
	for _, p := range pm.products {
		if p.CodeValue == product.CodeValue {
			err = internal.ErrCodeValueAlreadyExists
			return
		}
	}
	product.ID = pm.lastId
	pm.lastId++
	pm.products[product.ID] = *product
	return
}

func (pm *ProductMap) GetById(productId int) (product internal.Product, err error) {
	product, ok := pm.products[productId]
	if !ok {
		err = internal.ErrProductNotFound
		return
	}
	return
}
