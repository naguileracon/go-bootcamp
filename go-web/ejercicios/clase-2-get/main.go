package main

import (
	"clase-2-get/controller"
	"clase-2-get/service/importer"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	productImporter := importer.NewProductsImporter()
	productController, _ := controller.NewControllerProduct(productImporter)
	r := chi.NewRouter()
	r.HandleFunc("/ping", productController.GetPing())
	r.HandleFunc("/products", productController.GetProducts())
	r.HandleFunc("/products/{productId}", productController.GetProductById())
	r.HandleFunc("/products/search", productController.GetExpensiveProducts())
	http.ListenAndServe(":8080", r)
}
