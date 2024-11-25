package vehicle

import "fmt"

// Bike struct, composed of BaseVehicle
type Bike struct {
    BaseVehicle // Composition
    HasCarrier  bool
}

// Implement Vehicle interface for Bike
func (b Bike) Start() string {
    return fmt.Sprintf("Bike %s is starting.", b.Details())
}

func (b Bike) Stop() string {
    return fmt.Sprintf("Bike %s is stopping.", b.Details())
}

func (b Bike) Drive(distance int) string {
    return fmt.Sprintf("Bike %s rode %d kilometers.", b.Details(), distance)
}
