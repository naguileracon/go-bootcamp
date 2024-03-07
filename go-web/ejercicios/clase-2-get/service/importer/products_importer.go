package importer

import (
	"clase-2-get/structure"
	"encoding/json"
	"fmt"
	"os"
)

func NewProductsImporter() *ProductsImporterStruct {
	return &ProductsImporterStruct{}

}

type ProductsImporterStruct struct {
	Products []structure.Product
}

func (pi *ProductsImporterStruct) Import(filePath string) (err error) {
	// Read the content of the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Create a variable to store the unmarshalled data
	var products []structure.Product

	// Use the decoder to read and decode the JSON content
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&products)
	if err != nil {
		err = fmt.Errorf("error decoding JSON: %w", err)
		return err
	}

	//fmt.Println("Products:", products)
	pi.Products = products
	return

}

func (pi *ProductsImporterStruct) GetProductById(productId int) (product structure.Product, err error) {
	for _, pr := range pi.Products {
		if pr.Id == productId {
			product = pr
			return
		}
	}
	err = fmt.Errorf("product not found")
	return
}

func (pi *ProductsImporterStruct) GetExpensiveProducts(priceGt float64) (products []structure.Product) {
	products = make([]structure.Product, 0)
	for _, pr := range pi.Products {
		if pr.Price > priceGt {
			products = append(products, pr)
		}
	}
	return
}
