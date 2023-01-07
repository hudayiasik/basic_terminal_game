package main

import (
	"fmt"
	"time"
)

var x int
var y int
var height int = 10
var width int = 20
var key rune // keyboard input (w, s, a, d, q)

func main() {
	//infinit loop
	for {
		display()      // display the map
		fmt.Scan(&key) // get the keyboard input
		if move() == false {break} // move the player until press q
		time.Sleep(500 * time.Millisecond) // sleep 500ms
	}
}

func move() bool {
	switch key {
	case 'w': //up
		if x > 0 {
			x--
		}
	case 's': //down
		if x < height-1 {
			x++
		}
	case 'a': //left
		if y > 0 {
			y--
		}
	case 'd': //right
		if y < width-1 {
			y++
		}
	case 'q': //quit
		return false
	}
	return true
}

func display() { // display the map
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == x && j == y {
				fmt.Print("O") // player
			} else {
				fmt.Print(".") // empty
			}
		}
		fmt.Println() // new line
	}
}
