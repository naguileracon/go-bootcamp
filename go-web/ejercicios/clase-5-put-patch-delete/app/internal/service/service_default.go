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

func (sd *ServiceDefault) Update(product internal.Product) (err error) {
	err = sd.repository.Update(product)
	return
}

func (sd *ServiceDefault) UpdatePartial(id int, fields map[string]any) (err error) {
	err = sd.repository.UpdatePartial(id, fields)
	return
}

func (sd *ServiceDefault) Delete(id int) (err error) {
	err = sd.repository.Delete(id)
	return
}
