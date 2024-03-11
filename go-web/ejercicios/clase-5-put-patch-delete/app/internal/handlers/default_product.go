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

func (d *DefaultProduct) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

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
			ID:          id,
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}
		// - save the task
		if err := d.sv.Update(product); err != nil {
			switch {
			case errors.Is(err, internal.ErrCodeValueAlreadyExists):
				response.JSON(w, http.StatusInternalServerError, map[string]any{"message": internal.ErrCodeValueAlreadyExists.Error()})
			case errors.Is(err, internal.ErrProductNotFound):
				response.JSON(w, http.StatusNotFound, map[string]any{"message": internal.ErrProductNotFound.Error()})
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
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product updated",
			"data":    data,
		})

	}
}

func (d *DefaultProduct) UpdatePartial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

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

		if title, ok := bodyMap["name"]; ok {
			_, ok := title.(string)
			if !ok {
				response.Text(w, http.StatusBadRequest, "name is invalid")
				return
			}
		}

		// the same for quantity
		if quantity, ok := bodyMap["quantity"]; ok {
			_, ok := quantity.(float64)
			if !ok {
				response.Text(w, http.StatusBadRequest, "quantity is invalid")
				return
			}
			bodyMap["quantity"] = int(quantity.(float64))
		}

		if codeValue, ok := bodyMap["code_value"]; ok {
			_, ok := codeValue.(string)
			if !ok {
				response.Text(w, http.StatusBadRequest, "code_value is invalid")
				return
			}
		}

		if expiration, ok := bodyMap["expiration"]; ok {
			_, ok := expiration.(string)
			if !ok {
				response.Text(w, http.StatusBadRequest, "expiration is invalid")
				return
			}
			if err := tools.CheckDate(bodyMap["expiration"].(string)); err != nil {
				var dateError *tools.DateError
				if errors.As(err, &dateError) {
					response.JSON(w, http.StatusBadRequest, map[string]any{"message": fmt.Sprintf("%s is not a valid date", dateError.Date)})
					return
				}
				response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "internal server error"})
				return
			}
		}

		if price, ok := bodyMap["price"]; ok {
			_, ok := price.(float64)
			if !ok {
				response.Text(w, http.StatusBadRequest, "price is required")
				return
			}
		}

		if err := d.sv.UpdatePartial(id, bodyMap); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			case errors.Is(err, internal.ErrCodeValueAlreadyExists):
				response.Text(w, http.StatusBadRequest, "code value already exists")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		response.Text(w, http.StatusOK, "product updated")
	}
}

func (d *DefaultProduct) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		if err := d.sv.Delete(id); err != nil {
			switch {
			case errors.Is(err, internal.ErrProductNotFound):
				response.Text(w, http.StatusNotFound, "product not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		response.Text(w, http.StatusOK, "product deleted")
	}
}
