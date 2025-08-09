package main

import "fmt"

type engine struct {}

func (e engine) Start() {
	fmt.Println("Engine Started")
}