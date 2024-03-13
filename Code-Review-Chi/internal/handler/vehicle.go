package handler

import (
	"app/internal"
	"app/internal/service"
	"app/platform"
	"encoding/json"
	"errors"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
	"strings"
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
		if _, err := h.sv.Create(vehicle); err != nil {
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

		vehicleJSON := VehicleJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.Brand,
			Model:           vehicle.Model,
			Registration:    vehicle.Registration,
			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			Height:          vehicle.Height,
			Length:          vehicle.Length,
			Width:           vehicle.Width,
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "vehicle created successfully",
			"data":    vehicleJSON,
		})
	}
}

func (h *VehicleDefault) UpdateMaxSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid id")
			return
		}

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
		requiredFields := []string{"max_speed"}
		if err := platform.ValidateRequiredFields(bodyMap, requiredFields...); err != nil {
			response.JSON(w, err.HttpStatusCode, map[string]any{
				"message": err.ErrorMessage,
			})
			return
		}
		// process
		maxSpeed, ok := bodyMap["max_speed"].(float64)
		if !ok {
			response.Error(w, http.StatusBadRequest, "Invalid max_speed")
			return
		}
		vehicle, err := h.sv.UpdateMaxSpeed(id, maxSpeed)
		if err != nil {
			var VehicleServiceError *service.VehicleServiceError
			switch {
			case errors.As(err, &VehicleServiceError):
				response.JSON(w, VehicleServiceError.HttpStatusCode, map[string]any{
					"message": VehicleServiceError.ErrorMessage,
				})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"message": "internal server error",
				})
			}
			return
		}
		vehicleJSON := VehicleJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.Brand,
			Model:           vehicle.Model,
			Registration:    vehicle.Registration,
			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			Height:          vehicle.Height,
			Length:          vehicle.Length,
			Width:           vehicle.Width,
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "max speed updated successfully",
			"data":    vehicleJSON,
		})
	}
}

func (h *VehicleDefault) CreateMultiple() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// reading body
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"message": "invalid request vehicle body"})
			return
		}

		// -parse to slice of map (dynamic)
		var bodyMap []map[string]any
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request vehicle body",
			})
			return
		}

		// validate require fields
		requiredFields := []string{"id", "brand", "model", "registration", "color", "year", "passengers", "max_speed",
			"fuel_type", "transmission", "weight", "height", "width"}
		for _, value := range bodyMap {
			if err := platform.ValidateRequiredFields(value, requiredFields...); err != nil {
				response.JSON(w, err.HttpStatusCode, map[string]any{
					"message": err.ErrorMessage,
				})
				return
			}
		}

		// parse to slice of Vehicle struct
		var vehiclesRequestBody []VehicleRequestBody
		err = json.Unmarshal(bytes, &vehiclesRequestBody)
		if err != nil || len(vehiclesRequestBody) == 0 {
			response.Error(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		// save
		var vehicles []internal.Vehicle
		for _, value := range vehiclesRequestBody {
			vehicles = append(vehicles, internal.Vehicle{
				Id: value.ID,
				VehicleAttributes: internal.VehicleAttributes{
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
					Dimensions: internal.Dimensions{
						Height: value.Height,
						Length: value.Length,
						Width:  value.Width,
					},
				},
			})
		}
		vehicles, err = h.sv.CreateMultiple(vehicles)

		if err != nil {
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

		var vehiclesJSON []VehicleJSON
		for _, vehicle := range vehicles {
			vehiclesJSON = append(vehiclesJSON, VehicleJSON{
				ID:              vehicle.Id,
				Brand:           vehicle.Brand,
				Model:           vehicle.Model,
				Registration:    vehicle.Registration,
				Color:           vehicle.Color,
				FabricationYear: vehicle.FabricationYear,
				Capacity:        vehicle.Capacity,
				MaxSpeed:        vehicle.MaxSpeed,
				FuelType:        vehicle.FuelType,
				Transmission:    vehicle.Transmission,
				Weight:          vehicle.Weight,
				Height:          vehicle.Height,
				Length:          vehicle.Length,
				Width:           vehicle.Width,
			})
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "vehicles created successfully",
			"data":    vehiclesJSON,
		})
	}
}

func (h *VehicleDefault) GetByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// read query param
		height := r.URL.Query().Get("height")
		if height == "" {
			response.Error(w, http.StatusBadRequest, "Invalid height")
			return
		}
		width := r.URL.Query().Get("width")
		if width == "" {
			response.Error(w, http.StatusBadRequest, "Invalid width")
			return
		}
		maxHeight, err := strconv.ParseFloat(strings.Split(height, "-")[1], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid max length")
			return
		}
		minHeight, err := strconv.ParseFloat(strings.Split(height, "-")[0], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid min length")
			return
		}
		maxWidth, err := strconv.ParseFloat(strings.Split(width, "-")[1], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid max width")
			return
		}
		minWidth, err := strconv.ParseFloat(strings.Split(width, "-")[0], 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid min width")
			return
		}

		//process
		vehicles, err := h.sv.GetByDimensions(maxWidth, minWidth, maxHeight, minHeight)
		if err != nil {
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
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "vehicles found successfully",
			"data":    vehicles,
		})
	}
}

func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// read url param
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid id")
			return
		}

		// process
		err = h.sv.Delete(id)
		if err != nil {
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
		response.JSON(w, http.StatusNoContent, nil)
	}
}

func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request
		brand := chi.URLParam(r, "brand")

		//process
		averageSpeed, err := h.sv.GetAverageSpeedByBrand(brand)
		if err != nil {
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
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "average speed found successfully",
			"data":    averageSpeed,
		})
	}
}

func (h *VehicleDefault) GetByBrandAndRangeOfYears() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")
		if brand == "" {
			response.Error(w, http.StatusBadRequest, "Brand cannot be empty")
			return
		}
		startYear := chi.URLParam(r, "start_year")
		if startYear == "" {
			response.Error(w, http.StatusBadRequest, "Start year cannot be empty")
			return
		}
		endYear := chi.URLParam(r, "end_year")
		if endYear == "" {
			response.Error(w, http.StatusBadRequest, "End year cannot be empty")
			return
		}
		startYearInt, err := strconv.Atoi(startYear)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Start year must be a number")
			return
		}
		endYearInt, err := strconv.Atoi(endYear)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "End year must be a number")
			return
		}

		// process
		vehicles, err := h.sv.GetByBrandAndRangeOfYears(brand, startYearInt, endYearInt)
		if err != nil {
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
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "vehicles found successfully",
			"data":    vehicles,
		})
	}
}

func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		color := chi.URLParam(r, "color")
		if color == "" {
			response.Error(w, http.StatusBadRequest, "Color cannot be empty")
			return
		}
		year := chi.URLParam(r, "year")
		if year == "" {
			response.Error(w, http.StatusBadRequest, "Year cannot be empty")
			return
		}
		yearInt, err := strconv.Atoi(year)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Year must be a number")
			return
		}

		// process
		vehicles, err := h.sv.GetByColorAndYear(color, yearInt)
		if err != nil {
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
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "vehicles found successfully",
			"data":    vehicles,
		})
	}
}

func (h *VehicleDefault) GetByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		fuelType := chi.URLParam(r, "type")
		if fuelType == "" {
			response.Error(w, http.StatusBadRequest, "Fuel type cannot be empty")
			return
		}

		// process
		vehicles, err := h.sv.GetByFuelType(fuelType)
		if err != nil {
			var serviceError *service.VehicleServiceError
			switch {
			case errors.As(err, &serviceError):
				response.Error(w, serviceError.HttpStatusCode, serviceError.ErrorMessage)
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "vehicles found successfully",
			"data":    vehicles,
		})
	}
}

func (h *VehicleDefault) GetByTransmission() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transmission := chi.URLParam(r, "transmission")
		if transmission == "" {
			response.Error(w, http.StatusBadRequest, "Transmission cannot be empty")
			return
		}

		// process
		vehicles, err := h.sv.GetByTransmission(transmission)
		if err != nil {
			var serviceError *service.VehicleServiceError
			switch {
			case errors.As(err, &serviceError):
				response.Error(w, serviceError.HttpStatusCode, serviceError.ErrorMessage)
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "vehicles found successfully",
			"data":    vehicles,
		})
	}
}

func (h *VehicleDefault) UpdateFuel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id := chi.URLParam(r, "id")
		if id == "" {
			response.Error(w, http.StatusBadRequest, "Id cannot be empty")
			return
		}
		idInt, err := strconv.Atoi(id)

		// reading body
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"message": "invalid request body"})
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
		requiredFields := []string{"fuel_type"}
		if err := platform.ValidateRequiredFields(bodyMap, requiredFields...); err != nil {
			response.JSON(w, err.HttpStatusCode, map[string]any{
				"message": err.ErrorMessage,
			})
			return
		}

		//process
		fuelTypeString, ok := bodyMap["fuel_type"].(string)
		if !ok {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "fuel type must be a string",
			})
			return
		}

		vehicle, err := h.sv.UpdateFuel(idInt, fuelTypeString)
		if err != nil {
			var serviceError *service.VehicleServiceError
			switch {
			case errors.As(err, &serviceError):
				response.Error(w, serviceError.HttpStatusCode, serviceError.ErrorMessage)
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "fuel type updated successfully",
			"data":    vehicle,
		})

	}
}

func (h *VehicleDefault) GetAverageCapacityByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		brand := chi.URLParam(r, "brand")
		if brand == "" {
			response.Error(w, http.StatusBadRequest, "Brand cannot be empty")
			return
		}

		//process
		averageCapacity, err := h.sv.GetAverageCapacityByBrand(brand)
		if err != nil {
			var serviceError *service.VehicleServiceError
			switch {
			case errors.As(err, &serviceError):
				response.Error(w, serviceError.HttpStatusCode, serviceError.ErrorMessage)
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "average capacity found successfully",
			"data":    averageCapacity,
		})
	}
}

func (h *VehicleDefault) GetByRangeOfWeight() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		minWeight := r.URL.Query().Get("min")
		if minWeight == "" {
			response.Error(w, http.StatusBadRequest, "Min weight cannot be empty")
			return
		}
		minWeightFloat, err := strconv.ParseFloat(minWeight, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Min weight must be a number")
			return
		}
		maxWeight := r.URL.Query().Get("max")
		if maxWeight == "" {
			response.Error(w, http.StatusBadRequest, "Max weight cannot be empty")
			return
		}
		maxWeightFloat, err := strconv.ParseFloat(maxWeight, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Max weight must be a number")
			return
		}

		// process
		vehicles, err := h.sv.GetByRangeOfWeight(minWeightFloat, maxWeightFloat)
		if err != nil {
			var serviceError *service.VehicleServiceError
			switch {
			case errors.As(err, &serviceError):
				response.Error(w, serviceError.HttpStatusCode, serviceError.ErrorMessage)
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "vehicles found successfully",
			"data":    vehicles,
		})
	}
}
