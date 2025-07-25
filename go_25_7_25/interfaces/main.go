package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name          string
	Color 		  string
	NumberOfTeeth int
}


func main() {
	dog := Dog{
		Name: "Soonpy",
		Breed: "Pumarian",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name: "Guj",
		Color: "Black",
		NumberOfTeeth: 22,
	}

	PrintInfo(&gorilla)


}

func PrintInfo(a Animal)  {
	log.Println("This animal says", a.Says(), "and it is having", a.NumberOfLegs())
}

func (d *Dog) Says() string {
	return "Bowf"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Wooo"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}