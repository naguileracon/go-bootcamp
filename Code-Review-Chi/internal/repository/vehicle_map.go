package repository

import (
	"app/internal"
	"net/http"
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
func (r *VehicleMap) Create(v *internal.Vehicle) (err error) {
	if r.vehicleAlreadyExists(v.Id) {
		err = NewVehicleRepositoryError("vehicle already exists", http.StatusConflict)
		return
	}
	r.db[v.Id] = *v
	return
}

// UpdateMaxSpeed is a method that updates the max speed of a vehicle
func (r *VehicleMap) UpdateMaxSpeed(id int, newMaxSpeed float64) (err error) {
	if !r.vehicleAlreadyExists(id) {
		err = NewVehicleRepositoryError("vehicle does not exist", http.StatusNotFound)
		return
	}
	vehicle := r.db[id]
	vehicle.MaxSpeed = newMaxSpeed
	r.db[id] = vehicle
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
