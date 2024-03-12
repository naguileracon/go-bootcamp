package service

import (
	"app/internal"
	"app/internal/repository"
	"errors"
	"net/http"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// Create is a method that creates a vehicle
func (s *VehicleDefault) Create(v *internal.Vehicle) (err error) {
	err = s.rp.Create(v)
	if err != nil {
		// extract repository error from err
		var r *repository.VehicleRepositoryError
		// validate if err is a repository error
		switch {
		case errors.As(err, &r):
			err = NewVehicleServiceError(r.Error(), r.HttpStatusCode)
		default:
			err = NewVehicleServiceError("vehicle service error", http.StatusInternalServerError)
			return
		}
	}
	return
}

func (s *VehicleDefault) UpdateMaxSpeed(id int, newMaxSpeed float64) (err error) {
	err = s.rp.UpdateMaxSpeed(id, newMaxSpeed)
	var r *repository.VehicleRepositoryError
	if err != nil {
		switch {
		case errors.As(err, &r):
			err = NewVehicleServiceError(r.Error(), r.HttpStatusCode)
		default:
			err = NewVehicleServiceError("vehicle service error", http.StatusInternalServerError)
		}
	}
	return
}
