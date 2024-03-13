package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Create is a method that creates a vehicle
	Create(v Vehicle) (vehicle Vehicle, err error)
	// UpdateMaxSpeed is a method that updates the max speed of a vehicle
	UpdateMaxSpeed(id int, newMaxSpeed float64) (vehicle Vehicle, err error)
	// CreateMultiple is a method that creates multiple vehicles
	CreateMultiple(v []Vehicle) (vehicles []Vehicle, err error)
	// GetByDimensions is a method that returns a map of vehicles by dimensions
	GetByDimensions(maxWidth float64, minWidth float64, maxHeight float64, minHeight float64) (vehicles []Vehicle, err error)
	// Delete is a method that deletes a vehicle
	Delete(id int) (err error)
	// GetAverageSpeedByBrand is a method that returns the average speed of by brand
	GetAverageSpeedByBrand(brand string) (averageSpeed float64, err error)
	// GetByBrandAndRangeOfYears is a method that returns a slice of vehicles by brand and range of years
	GetByBrandAndRangeOfYears(brand string, minYear int, maxYear int) (vehicles []Vehicle, err error)
	// GetByColorAndYear is a method that returns a slice of vehicles by color and year
	GetByColorAndYear(color string, year int) (vehicles []Vehicle, err error)
	// GetByFuelType is a method that returns a slice of vehicles by fuel type
	GetByFuelType(fuelType string) (vehicles []Vehicle, err error)
	// GetByTransmission is a method that returns a slice of vehicles by transmission
	GetByTransmission(transmission string) (vehicles []Vehicle, err error)
	// UpdateFuel is a method that updates the fuel of a vehicle
	UpdateFuel(id int, newFuel string) (vehicle Vehicle, err error)
	// GetAverageCapacityByBrand is a method that returns the average capacity of by brand
	GetAverageCapacityByBrand(brand string) (averageCapacity float64, err error)
	//GetByRangeOfWeight is a method that returns a slice of vehicles by range of weight
	GetByRangeOfWeight(minWeight float64, maxWeight float64) (vehicles []Vehicle, err error)
}
