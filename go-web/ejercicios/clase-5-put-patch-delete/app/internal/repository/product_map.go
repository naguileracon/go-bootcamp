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

func (pm *ProductMap) Update(product internal.Product) (err error) {
	_, ok := pm.products[product.ID]
	if !ok {
		err = internal.ErrProductNotFound
		return
	}
	// validate if code_value already exists
	for _, p := range pm.products {
		if p.CodeValue == product.CodeValue && p.ID != product.ID {
			err = internal.ErrCodeValueAlreadyExists
			return
		}
	}
	pm.products[product.ID] = product
	return
}

func (pm *ProductMap) UpdatePartial(id int, fields map[string]any) (err error) {
	product, ok := pm.products[id]
	if !ok {
		err = internal.ErrProductNotFound
		return
	}
	for field, value := range fields {
		switch field {
		case "name":
			product.Name = value.(string)
		case "quantity":
			product.Quantity = value.(int)
		case "code_value":
			for _, p := range pm.products {
				if p.CodeValue == value && p.ID != product.ID {
					err = internal.ErrCodeValueAlreadyExists
					return
				}
			}
			product.CodeValue = value.(string)
		case "is_published":
			product.IsPublished = value.(bool)
		case "expiration":
			product.Expiration = value.(string)
		case "price":
			product.Price = value.(float64)
		}
	}
	pm.products[id] = product
	return
}

func (pm *ProductMap) Delete(id int) (err error) {
	_, ok := pm.products[id]
	if !ok {
		err = internal.ErrProductNotFound
		return
	}
	delete(pm.products, id)
	return
}
