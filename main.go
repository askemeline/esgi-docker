package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Player is a representation of a player
type Player struct {
	name string
    x int
	y int
	vie int
}

func newPlayer(_name string, _x int, _y int, _vie int) *Player{
	p := Player{
		name :_name,
		x : _x,
		y: _y,
		vie:_vie,
	}
    return &p
}

func (p* Player) attack(p2* Player, damange int){
	p2.vie -= damange
}

func (p* Player) moveUp(p2* Player, mvUp int){
	p2.x += mvUp
}

func (p* Player) moveDown(p2* Player, mvDown int){
	p2.y -= mvDown
}


func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func main() {
	p1 := Player{name:"tata", x:80, y: 20, vie:99}
	p2 := Player{name:"toto", x:90, y: 20, vie:69}
	p1.attack(&p2, 30)
	p1.moveDown(&p2, 30)
	p1.moveUp(&p2, 70)
    rand.Seed(time.Now().UTC().UnixNano())
	switch randInt(1, 3) {
	case 1:
		 fmt.Println("name :", p2.name, "positionX :", p2.x, ", positionY :", p2.x,", Vie :", p2.vie)
	case 2:
		 fmt.Println("name :", p1.name, "positionX :", p1.x, ", positionY :", p1.x,", Vie :", p1.vie)
	}
	// fmt.Println(random())
}