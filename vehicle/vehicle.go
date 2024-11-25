package vehicle

import "fmt"

// Vehicle is the interface for all vehicles
type Vehicle interface {
    Start() string
    Stop() string
    Drive(distance int) string
}

// BaseVehicle struct contains common fields
type BaseVehicle struct {
    Make  string
    Model string
    Year  int
}

// BaseVehicle methods
func (b BaseVehicle) Details() string {
    return fmt.Sprintf("%d %s %s", b.Year, b.Make, b.Model)
}
