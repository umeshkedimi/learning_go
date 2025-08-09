package main

import "fmt"

type Convertable struct {
	Engine
	Transmission
	SteeringWheel
}

func (c Convertable) ConvertUp() {
	fmt.Println("Coverting Up")
}