package internal

import (
	"net/http"
	"strconv"
)

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (e *errorResponse) Error() string {
	return e.Status + ": " + e.Message
}

var (
	ErrCodeValueAlreadyExists = &errorResponse{Status: strconv.Itoa(http.StatusConflict), Message: "code_value already exists"}
	ErrProductNotFound        = &errorResponse{Status: strconv.Itoa(http.StatusNotFound), Message: "product not found"}
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
	Update(product Product) (err error)
	UpdatePartial(id int, fields map[string]any) (err error)
	Delete(id int) (err error)
}

type ProductService interface {
	Save(product *Product) (err error)
	GetById(productId int) (product Product, err error)
	Update(product Product) (err error)
	UpdatePartial(id int, fields map[string]any) (err error)
	Delete(id int) (err error)
}
