package repository

import (
	"app/internal"
	"fmt"
	"net/http"
	"strings"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// Create is a method that creates a vehicle
func (r *VehicleMap) Create(v internal.Vehicle) (vehicle internal.Vehicle, err error) {
	if r.vehicleAlreadyExists(v.Id) {
		err = NewVehicleRepositoryError(fmt.Sprintf("Vehicle with id: %d already exists", v.Id), http.StatusConflict)
		return
	}
	r.db[v.Id] = v
	vehicle = v
	return
}

// CreateMultiple is a method that creates multiple vehicles
func (r *VehicleMap) CreateMultiple(v []internal.Vehicle) (vehicles []internal.Vehicle, err error) {
	// creating a copy of the original slice
	copyOriginalSlice := make(map[int]internal.Vehicle, len(r.db))
	for key, value := range r.db {
		copyOriginalSlice[key] = value
	}
	for _, value := range v {
		var vehicle internal.Vehicle
		vehicle, err = r.Create(value)
		vehicles = append(vehicles, vehicle)
		if err != nil {
			r.db = copyOriginalSlice
			return
		}
	}
	return
}

// UpdateMaxSpeed is a method that updates the max speed of a vehicle
func (r *VehicleMap) UpdateMaxSpeed(id int, newMaxSpeed float64) (vehicle internal.Vehicle, err error) {
	if !r.vehicleAlreadyExists(id) {
		err = NewVehicleRepositoryError("vehicle does not exist", http.StatusNotFound)
		return
	}
	vehicle = r.db[id]
	vehicle.MaxSpeed = newMaxSpeed
	r.db[id] = vehicle
	return
}

// GetByDimensions is a method that returns a map of vehicles by dimensions
func (r *VehicleMap) GetByDimensions(maxWidth float64, minWidth float64, maxHeight float64, minHeight float64) (vehicles []internal.Vehicle, err error) {
	println(maxWidth, minWidth, maxHeight, minHeight)
	for _, value := range r.db {
		fmt.Printf("value width: %f, value height: %f\n", value.Width, value.Height)
		// print function parameters
		fmt.Printf("maxWidth: %f, minWidth: %f, maxHeight: %f, minHeight: %f\n", maxWidth, minWidth, maxHeight, minHeight)
		if value.Width <= maxWidth && value.Width >= minWidth && value.Height <= maxHeight && value.Height >= minHeight {
			vehicles = append(vehicles, value)
		}
	}
	if len(vehicles) == 0 {
		err = NewVehicleRepositoryError("no vehicles found with the desired dimensions", http.StatusNotFound)
		return
	}
	return
}

func (r *VehicleMap) Delete(id int) (err error) {
	if !r.vehicleAlreadyExists(id) {
		err = NewVehicleRepositoryError("vehicle does not exist", http.StatusNotFound)
		return
	}
	delete(r.db, id)
	return
}

func (r *VehicleMap) GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	for _, value := range r.db {
		if strings.ToLower(value.Brand) == strings.ToLower(brand) {
			averageSpeed += value.MaxSpeed
		}
	}
	if averageSpeed == 0 {
		err = NewVehicleRepositoryError("no vehicles found with the desired brand", http.StatusNotFound)
		return
	}
	averageSpeed = averageSpeed / float64(len(r.db))
	return
}

func (r *VehicleMap) GetByBrandAndRangeOfYears(brand string, minYear int, maxYear int) (vehicles []internal.Vehicle, err error) {
	for _, value := range r.db {
		if strings.ToLower(value.Brand) == strings.ToLower(brand) && value.FabricationYear >= minYear && value.FabricationYear <= maxYear {
			vehicles = append(vehicles, value)
		}
	}
	if len(vehicles) == 0 {
		err = NewVehicleRepositoryError("no vehicles found with the desired brand and range of years", http.StatusNotFound)
		return
	}
	return
}

func (r *VehicleMap) GetByColorAndYear(color string, year int) (vehicles []internal.Vehicle, err error) {
	for _, value := range r.db {
		if strings.ToLower(value.Color) == strings.ToLower(color) && value.FabricationYear == year {
			vehicles = append(vehicles, value)
		}
	}
	if len(vehicles) == 0 {
		err = NewVehicleRepositoryError("no vehicles found with the desired color and year", http.StatusNotFound)
		return
	}
	return
}

func (r *VehicleMap) GetByFuelType(fuelType string) (vehicles []internal.Vehicle, err error) {
	for _, value := range r.db {
		if strings.ToLower(value.FuelType) == strings.ToLower(fuelType) {
			vehicles = append(vehicles, value)
		}
	}
	if len(vehicles) == 0 {
		err = NewVehicleRepositoryError("no vehicles found with the desired fuel type", http.StatusNotFound)
		return
	}
	return
}

func (r *VehicleMap) GetByTransmission(transmission string) (vehicles []internal.Vehicle, err error) {
	for _, value := range r.db {
		if strings.ToLower(transmission) == strings.ToLower(value.Transmission) {
			vehicles = append(vehicles, value)
		}
	}
	if len(vehicles) == 0 {
		err = NewVehicleRepositoryError("no vehicles found with the desired transmission", http.StatusNotFound)
		return
	}
	return
}

func (r *VehicleMap) GetAverageCapacityByBrand(brand string) (averageCapacity float64, err error) {
	for _, value := range r.db {
		if strings.ToLower(value.Brand) == strings.ToLower(brand) {
			averageCapacity += float64(value.Capacity)
		}
	}
	if averageCapacity == 0 {
		err = NewVehicleRepositoryError("no vehicles found with the desired brand", http.StatusNotFound)
		return
	}
	averageCapacity = averageCapacity / float64(len(r.db))
	return
}

func (r *VehicleMap) UpdateFuel(id int, newFuel string) (vehicle internal.Vehicle, err error) {
	if !r.vehicleAlreadyExists(id) {
		err = NewVehicleRepositoryError("vehicle does not exist", http.StatusNotFound)
		return
	}
	vehicle = r.db[id]
	vehicle.FuelType = newFuel
	r.db[id] = vehicle
	return
}

func (r *VehicleMap) GetByRangeOfWeight(minWeight float64, maxWeight float64) (vehicles []internal.Vehicle, err error) {
	for _, value := range r.db {
		if value.Weight >= minWeight && value.Weight <= maxWeight {
			vehicles = append(vehicles, value)
		}
	}
	if len(vehicles) == 0 {
		err = NewVehicleRepositoryError("no vehicles found with the desired range of weight", http.StatusNotFound)
		return
	}
	return
}

// vehicleAlreadyExists is a method that returns true if the vehicle already exists
// A vehicle already exists if the id is equal to the id of any vehicle in the db
func (r *VehicleMap) vehicleAlreadyExists(id int) bool {
	for _, value := range r.db {
		if value.Id == id {
			return true
		}
	}
	return false
}
