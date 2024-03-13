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
func (s *VehicleDefault) Create(v internal.Vehicle) (vehicle internal.Vehicle, err error) {
	vehicle, err = s.rp.Create(v)
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

// CreateMultiple is a method that creates multiple vehicles
func (s *VehicleDefault) CreateMultiple(v []internal.Vehicle) (vehicles []internal.Vehicle, err error) {
	vehicles, err = s.rp.CreateMultiple(v)
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

func (s *VehicleDefault) UpdateMaxSpeed(id int, newMaxSpeed float64) (vehicle internal.Vehicle, err error) {
	vehicle, err = s.rp.UpdateMaxSpeed(id, newMaxSpeed)
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

func (s *VehicleDefault) GetByDimensions(maxWidth float64, minWidth float64, maxHeight float64, minHeight float64) (vehicles []internal.Vehicle, err error) {
	vehicles, err = s.rp.GetByDimensions(maxWidth, minWidth, maxHeight, minHeight)
	var r *repository.VehicleRepositoryError
	if err != nil {
		switch {
		case errors.As(err, &r):
			err = NewVehicleServiceError(r.Error(), r.HttpStatusCode)
		default:
			err = NewVehicleServiceError("vehicle service error", http.StatusInternalServerError)
		}
		return
	}
	return
}

func (s *VehicleDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	var r *repository.VehicleRepositoryError
	if err != nil {
		switch {
		case errors.As(err, &r):
			err = NewVehicleServiceError(r.Error(), r.HttpStatusCode)
		default:
			err = NewVehicleServiceError("vehicle service error", http.StatusInternalServerError)
		}
		return
	}
	return
}

func (s *VehicleDefault) GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error) {
	averageSpeed, err = s.rp.GetAverageSpeedByBrand(brand)
	if err != nil {
		var r *repository.VehicleRepositoryError
		switch {
		case errors.As(err, &r):

			err = NewVehicleServiceError(r.Error(), r.HttpStatusCode)
		default:
			err = NewVehicleServiceError("vehicle service error", http.StatusInternalServerError)
		}
	}
	return
}

func (s *VehicleDefault) GetByBrandAndRangeOfYears(brand string, minYear int, maxYear int) (vehicles []internal.Vehicle, err error) {
	vehicles, err = s.rp.GetByBrandAndRangeOfYears(brand, minYear, maxYear)
	if err != nil {
		var r *repository.VehicleRepositoryError
		switch {
		case errors.As(err, &r):
			err = NewVehicleServiceError(r.Error(), r.HttpStatusCode)
		default:
			err = NewVehicleServiceError("vehicle service error", http.StatusInternalServerError)
		}
		return
	}
	return
}

func (s *VehicleDefault) GetByFuelType(fuelType string) (vehicles []internal.Vehicle, err error) {
	vehicles, err = s.rp.GetByFuelType(fuelType)
	if err != nil {
		var r *repository.VehicleRepositoryError
		switch {
		case errors.As(err, &r):
			err = NewVehicleServiceError(r.Error(), r.HttpStatusCode)
		default:
			err = NewVehicleServiceError("vehicle service error", http.StatusInternalServerError)
		}
		return
	}
	return
}

func (s *VehicleDefault) GetByColorAndYear(color string, year int) (vehicles []internal.Vehicle, err error) {
	vehicles, err = s.rp.GetByColorAndYear(color, year)
	if err != nil {
		var r *repository.VehicleRepositoryError
		switch {
		case errors.As(err, &r):
			err = NewVehicleServiceError(r.Error(), r.HttpStatusCode)
		default:
			err = NewVehicleServiceError("vehicle service error", http.StatusInternalServerError)
		}
		return
	}
	return
}

func (s *VehicleDefault) GetByTransmission(transmission string) (vehicles []internal.Vehicle, err error) {
	vehicles, err = s.rp.GetByTransmission(transmission)
	if err != nil {
		var repositoryError *repository.VehicleRepositoryError
		switch {
		case errors.As(err, &repositoryError):
			err = NewVehicleServiceError(repositoryError.Error(), repositoryError.HttpStatusCode)
		default:
			err = NewVehicleServiceError("vehicle service error", http.StatusInternalServerError)
		}
		return
	}
	return
}

func (s *VehicleDefault) GetAverageCapacityByBrand(brand string) (averageCapacity float64, err error) {
	averageCapacity, err = s.rp.GetAverageCapacityByBrand(brand)
	if err != nil {
		var repoError *repository.VehicleRepositoryError
		switch {
		case errors.As(err, &repoError):
			err = NewVehicleServiceError(repoError.Error(), repoError.HttpStatusCode)
		default:
			err = NewVehicleServiceError("vehicle service error", http.StatusInternalServerError)
		}
		return
	}
	return
}

func (s *VehicleDefault) UpdateFuel(id int, newFuel string) (vehicle internal.Vehicle, err error) {
	vehicle, err = s.rp.UpdateFuel(id, newFuel)
	if err != nil {
		var repoError *repository.VehicleRepositoryError
		switch {
		case errors.As(err, &repoError):
			err = NewVehicleServiceError(repoError.Error(), repoError.HttpStatusCode)
		default:
			err = NewVehicleServiceError("vehicle service error", http.StatusInternalServerError)
		}
		return
	}
	return
}

func (s *VehicleDefault) GetByRangeOfWeight(minWeight float64, maxWeight float64) (vehicles []internal.Vehicle, err error) {
	vehicles, err = s.rp.GetByRangeOfWeight(minWeight, maxWeight)
	if err != nil {
		var repoError *repository.VehicleRepositoryError
		switch {
		case errors.As(err, &repoError):
			err = NewVehicleServiceError(repoError.Error(), repoError.HttpStatusCode)
		default:
			err = NewVehicleServiceError("vehicle service error", http.StatusInternalServerError)
		}
		return
	}
	return
}
