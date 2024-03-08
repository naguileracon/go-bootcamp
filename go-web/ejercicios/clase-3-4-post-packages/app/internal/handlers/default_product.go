package handlers

import (
	"app/internal"
	"app/pkg/tools"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
)

func NewDefaultProduct(sv internal.ProductService) *DefaultProduct {
	return &DefaultProduct{
		sv: sv,
	}
}

type DefaultProduct struct {
	sv internal.ProductService
}

type ProductJSON struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ProductRequestBody struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (d *DefaultProduct) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"message": "invalid request body"})
			return
		}

		// - parse to map (dynamic)
		bodyMap := map[string]any{}
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}

		// validate required fields
		if err := tools.CheckFieldExistance(bodyMap, "name", "quantity", "code_value", "expiration", "price"); err != nil {
			var fieldError *tools.FieldError
			if errors.As(err, &fieldError) {
				response.JSON(w, http.StatusBadRequest, map[string]any{"message": fmt.Sprintf("%s is required", fieldError.Field)})
				return
			}
			response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			return
		}

		// validate date structure
		if err := tools.CheckDate(bodyMap["expiration"].(string)); err != nil {
			var dateError *tools.DateError
			if errors.As(err, &dateError) {
				response.JSON(w, http.StatusBadRequest, map[string]any{"message": fmt.Sprintf("%s is not a valid date", dateError.Date)})
				return
			}
			response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			return
		}

		// - parse json to struct (static)
		var body ProductRequestBody
		if err := json.Unmarshal(bytes, &body); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}

		// process
		// - serialize the request body into a task
		product := internal.Product{
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}
		// - save the task
		if err := d.sv.Save(&product); err != nil {
			switch {
			case errors.Is(err, internal.ErrCodeValueAlreadyExists):
				response.JSON(w, http.StatusInternalServerError, map[string]any{"message": internal.ErrCodeValueAlreadyExists.Error()})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		data := ProductJSON{
			ID:          product.ID,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "task created",
			"data":    data,
		})

	}
}

func (d *DefaultProduct) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
		product, err := d.sv.GetById(productId)
		if err != nil {
			if errors.Is(err, internal.ErrProductNotFound) {
				response.JSON(w, http.StatusNotFound, map[string]any{"message": internal.ErrProductNotFound.Error()})
				return
			}
			response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "task created",
			"data":    product,
		})
	}
}
