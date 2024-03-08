package internal

import "errors"

var (
	ErrCodeValueAlreadyExists = errors.New("code_value already exists")
	ErrProductNotFound        = errors.New("product not found")
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ProductRepository interface {
	Save(product *Product) (err error)
	GetById(productId int) (product Product, err error)
}

type ProductService interface {
	Save(product *Product) (err error)
	GetById(productId int) (product Product, err error)
}
