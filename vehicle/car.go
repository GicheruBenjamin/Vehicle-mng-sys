package vehicle

import "fmt"

// Car struct, composed of BaseVehicle
type Car struct {
    BaseVehicle // Composition
    NumDoors    int
}

// Implement Vehicle interface for Car
func (c Car) Start() string {
    return fmt.Sprintf("Car %s is starting.", c.Details())
}

func (c Car) Stop() string {
    return fmt.Sprintf("Car %s is stopping.", c.Details())
}

func (c Car) Drive(distance int) string {
    return fmt.Sprintf("Car %s drove %d kilometers.", c.Details(), distance)
}
