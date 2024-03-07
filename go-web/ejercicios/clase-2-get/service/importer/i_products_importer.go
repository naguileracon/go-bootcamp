package importer

import "clase-2-get/structure"

type ProductsImporter interface {
	Import(filePath string) (err error)
	GetProductById(productId int) (product structure.Product, err error)
}
