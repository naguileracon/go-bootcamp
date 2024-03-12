package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Create is a method that creates a vehicle
	Create(v *Vehicle) (err error)
	// UpdateMaxSpeed is a method that updates the max speed of a vehicle
	UpdateMaxSpeed(id int, newMaxSpeed float64) (err error)
}
