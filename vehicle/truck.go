package vehicle

import "fmt"

// Truck struct, composed of BaseVehicle
type Truck struct {
    BaseVehicle // Composition
    LoadCapacity int // In kilograms
}

// Implement Vehicle interface for Truck
func (t Truck) Start() string {
    return fmt.Sprintf("Truck %s is starting.", t.Details())
}

func (t Truck) Stop() string {
    return fmt.Sprintf("Truck %s is stopping.", t.Details())
}

func (t Truck) Drive(distance int) string {
    return fmt.Sprintf("Truck %s transported goods %d kilometers.", t.Details(), distance)
}
