package main

import (
	"fmt"
)

type Player struct {
	Name     string
	Age      int
	Height   float64
	Weight   float64
	Position string
}

type Runner interface {
	Run()
}

type Walker interface {
	Walk()
}

type Shooter interface {
	Shoot()
}

func main() {
	player := &Player{}
	var runner Runner = player
	var walker Walker = player

	runner.Run()
	walker.Walk()
}

func (player *Player) Run() {
	fmt.Println("Player is running")
}

func (player *Player) Walk() {
	fmt.Println("Player is walking")
}
