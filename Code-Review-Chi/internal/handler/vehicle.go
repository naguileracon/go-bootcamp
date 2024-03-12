package handler

import (
	"app/internal"
	"app/internal/service"
	"app/platform"
	"encoding/json"
	"errors"
	"github.com/bootcamp-go/web/response"
	"io"
	"net/http"
)

// VehicleRequestBody is a struct that represents a vehicle in JSON format
type VehicleRequestBody struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// reading body
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"message": "invalid request vehicle body"})
			return
		}

		// -parse to map (dynamic)
		bodyMap := map[string]any{}
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request vehicle body",
			})
			return
		}

		// validate required fields
		requiredFields := []string{"id", "brand", "model", "registration", "color", "year", "passengers", "max_speed",
			"fuel_type", "transmission", "weight", "height", "width"}
		if err := platform.ValidateRequiredFields(bodyMap, requiredFields...); err != nil {
			response.JSON(w, err.HttpStatusCode, map[string]any{
				"message": err.ErrorMessage,
			})
			return
		}

		// parse to Vehicle struct
		var reqBody VehicleRequestBody
		err = json.Unmarshal(bytes, &reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		// save
		vehicle := internal.Vehicle{
			Id: reqBody.ID,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           reqBody.Brand,
				Model:           reqBody.Model,
				Registration:    reqBody.Registration,
				Color:           reqBody.Color,
				FabricationYear: reqBody.FabricationYear,
				Capacity:        reqBody.Capacity,
				MaxSpeed:        reqBody.MaxSpeed,
				FuelType:        reqBody.FuelType,
				Transmission:    reqBody.Transmission,
				Weight:          reqBody.Weight,
				Dimensions: internal.Dimensions{
					Height: reqBody.Height,
					Length: reqBody.Length,
					Width:  reqBody.Width,
				},
			},
		}
		if err := h.sv.Create(&vehicle); err != nil {
			var vehicleServiceError *service.VehicleServiceError
			switch {
			case errors.As(err, &vehicleServiceError):
				response.JSON(w, vehicleServiceError.HttpStatusCode, map[string]any{
					"message": vehicleServiceError.ErrorMessage,
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "internal server error",
				})
			}
			return
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    vehicle,
		})
	}
}
