package controller

import (
	"clase-2-get/service/importer"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

const FilePath = "products.json"

func NewControllerProduct(productsImporter *importer.ProductsImporterStruct) (controller *ControllerProduct, err error) {
	err = productsImporter.Import(FilePath)
	if err != nil {
		err = fmt.Errorf("error importing products: %w", err)
		return
	}
	controller = &ControllerProduct{*productsImporter}
	return
}

type ControllerProduct struct {
	ProductsImporter importer.ProductsImporterStruct
}

func (cp *ControllerProduct) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products := cp.ProductsImporter.Products
		if len(products) == 0 {
			code := http.StatusNotFound
			w.WriteHeader(code)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("products not found")
			return
		}
		code := http.StatusOK
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

func (cp *ControllerProduct) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productId, _ := strconv.Atoi(chi.URLParam(r, "productId"))
		product, err := cp.ProductsImporter.GetProductById(productId)
		if err != nil {
			code := http.StatusNotFound
			w.WriteHeader(code)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("product with id " + strconv.Itoa(productId) + " not found")
			return
		}
		code := http.StatusOK
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}
}

func (cp *ControllerProduct) GetExpensiveProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		priceGt, _ := strconv.ParseFloat(r.URL.Query().Get("priceGt"), 64)
		products := cp.ProductsImporter.GetExpensiveProducts(priceGt)
		code := http.StatusOK
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

func (cp *ControllerProduct) GetPing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := http.StatusOK
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("pong")
	}

}
