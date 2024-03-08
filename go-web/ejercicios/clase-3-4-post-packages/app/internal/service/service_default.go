package service

import (
	"app/internal"
)

func NewServiceDefault(repository internal.ProductRepository) *ServiceDefault {
	return &ServiceDefault{
		repository: repository,
	}
}

type ServiceDefault struct {
	repository internal.ProductRepository
}

func (sd *ServiceDefault) Save(product *internal.Product) (err error) {
	err = sd.repository.Save(product)
	return
}

func (sd *ServiceDefault) GetById(productId int) (product internal.Product, err error) {
	product, err = sd.repository.GetById(productId)
	if err != nil {
		return
	}
	return
}
