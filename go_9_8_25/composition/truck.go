package main

import "fmt"

type Truck struct {
	Engine
	Transmission
	SteeringWheel
}

func (t Truck) FourWheelDriver() {
	fmt.Println("Toggeling 4WD...")
}