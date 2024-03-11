package handlers

import (
	"app/internal/auth"
	"app/internal/product/repository"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestProducts_GetProducts_HandlerFunction(t *testing.T) {
	t.Run("should return a list of products", func(t *testing.T) {
		// arrange
		tokenString := "123456"
		authToken := auth.NewAuthTokenBasic(tokenString)
		productAttributes := map[int]repository.ProductAttributesMap{
			0: {
				Name:        "product 1",
				Quantity:    10,
				CodeValue:   "code 1",
				IsPublished: true,
				Expiration:  time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				Price:       100.0,
			},
		}
		rpMap := repository.NewRepositoryProductMap(productAttributes, 0, "")
		rp := repository.RepositoryProduct(rpMap)
		handler := NewHandlerProducts(rp, authToken)
		handlerFunction := handler.Get()
		// act
		req := httptest.NewRequest("GET", "/products", nil)
		req.Header.Set("Token", tokenString)
		res := httptest.NewRecorder()
		handlerFunction(res, req)
		// assert
		expectedResponse := map[string]any{
			"message": "products",
			"data": []ProductJSON{
				{
					Id:          0,
					Name:        productAttributes[0].Name,
					Quantity:    productAttributes[0].Quantity,
					CodeValue:   productAttributes[0].CodeValue,
					IsPublished: productAttributes[0].IsPublished,
					Expiration:  productAttributes[0].Expiration.Format("2006-01-02"),
					Price:       productAttributes[0].Price,
				},
			},
		}
		expectedResponseJSON, err := json.Marshal(expectedResponse)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json; charset=utf-8", res.Header().Get("Content-Type"))
		require.Equal(t, expectedResponseJSON, res.Body.Bytes())
	})

	t.Run("should get a product by id", func(t *testing.T) {
		// arrange
		tokenString := "123456"
		authToken := auth.NewAuthTokenBasic(tokenString)
		productAttributes := map[int]repository.ProductAttributesMap{
			0: {
				Name:        "product 1",
				Quantity:    10,
				CodeValue:   "code 1",
				IsPublished: true,
				Expiration:  time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				Price:       100,
			},
		}
		rpMap := repository.NewRepositoryProductMap(productAttributes, 0, "")
		rp := repository.RepositoryProduct(rpMap)
		handler := NewHandlerProducts(rp, authToken)
		handlerFunction := handler.GetByID()
		// act
		req := httptest.NewRequest("GET", "/products/0", nil)
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "0")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
		req.Header.Set("Token", tokenString)
		res := httptest.NewRecorder()
		handlerFunction(res, req)
		// assert
		expectedResponse := map[string]any{
			"message": "product",
			"data": ProductJSON{
				Id:          0,
				Name:        productAttributes[0].Name,
				Quantity:    productAttributes[0].Quantity,
				CodeValue:   productAttributes[0].CodeValue,
				IsPublished: productAttributes[0].IsPublished,
				Expiration:  productAttributes[0].Expiration.Format("2006-01-02"),
				Price:       productAttributes[0].Price,
			},
		}
		expectedResponseJSON, err := json.Marshal(expectedResponse)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json; charset=utf-8", res.Header().Get("Content-Type"))
		require.JSONEq(t, string(expectedResponseJSON), res.Body.String())
	})

	t.Run("should create a product", func(t *testing.T) {
		// arrange
		tokenString := "123456"
		authToken := auth.NewAuthTokenBasic(tokenString)

		rpMap := repository.NewRepositoryProductMap(make(map[int]repository.ProductAttributesMap), -1, "")
		rp := repository.RepositoryProduct(rpMap)
		handler := NewHandlerProducts(rp, authToken)
		handlerFunction := handler.Create()
		// act
		exp := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		// 30 days
		req := httptest.NewRequest("POST", "/products", strings.NewReader(
			fmt.Sprintf(`{"name": "product 1", "quantity": 1, "code_value": "code 1", "is_published": true, "expiration": "%s", "price": 1.1}`, exp.Format("2006-01-02"))))
		req.Header.Set("Token", tokenString)
		res := httptest.NewRecorder()
		handlerFunction(res, req)
		// assert
		require.Equal(t, http.StatusCreated, res.Code)
		require.Equal(t, "application/json; charset=utf-8", res.Header().Get("Content-Type"))
		expectedResponse := map[string]any{
			"message": "product created",
			"data": ProductJSON{
				Id:          0,
				Name:        "product 1",
				Quantity:    1,
				CodeValue:   "code 1",
				IsPublished: true,
				Expiration:  exp.Format(time.DateOnly),
				Price:       1.1,
			},
		}
		expectedResponseJSON, err := json.Marshal(expectedResponse)
		require.NoError(t, err)
		require.Equal(t, string(expectedResponseJSON), res.Body.String())
	})
}
