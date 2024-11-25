package main

import (
    "fmt"
    "vehicle-management/vehicle"
)

func main() {
    // Create instances of vehicles
    car := vehicle.Car{
        BaseVehicle: vehicle.BaseVehicle{Make: "Toyota", Model: "Corolla", Year: 2020},
        NumDoors:    4,
    }

    bike := vehicle.Bike{
        BaseVehicle: vehicle.BaseVehicle{Make: "Yamaha", Model: "YZF-R3", Year: 2021},
        HasCarrier:  false,
    }

    truck := vehicle.Truck{
        BaseVehicle:  vehicle.BaseVehicle{Make: "Volvo", Model: "FH16", Year: 2019},
        LoadCapacity: 20000,
    }

    // List of vehicles (polymorphism with interface)
    vehicles := []vehicle.Vehicle{car, bike, truck}

    // Perform actions on each vehicle
    for _, v := range vehicles {
        fmt.Println(v.Start())
        fmt.Println(v.Drive(100)) // Example distance
        fmt.Println(v.Stop())
        fmt.Println()
    }

    // Access specific details using structs
    fmt.Printf("Car details: %s with %d doors\n", car.Details(), car.NumDoors)
    fmt.Printf("Bike details: %s, carrier available: %v\n", bike.Details(), bike.HasCarrier)
    fmt.Printf("Truck details: %s with load capacity: %d kg\n", truck.Details(), truck.LoadCapacity)
}
