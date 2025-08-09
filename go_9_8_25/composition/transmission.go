package main

import "fmt"

type transmission struct {}

func (t transmission) ShiftUp() {
	fmt.Println("Shifting up")
}

func (t transmission) ShiftDown() {
	fmt.Println("Shifting Down")
}