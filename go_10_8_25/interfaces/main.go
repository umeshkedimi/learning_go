package main

import (
	"fmt"
	"math/rand"
)


type FootbalPlayer struct {
	stamina int
	power int
}

func (f FootbalPlayer) kickBall() {
	shot := f.stamina + f.power
	fmt.Println("I am kicking the ball", shot)
}

func main() {
	team := make([]FootbalPlayer, 11)
	for i := 0; i < len(team); i++ {
		team[i] = FootbalPlayer{
			stamina: rand.Intn(10),
			power : rand.Intn(10),
		}
	}
	for i := 0; i < len(team); i ++ {
		team[i].kickBall()
	}
}