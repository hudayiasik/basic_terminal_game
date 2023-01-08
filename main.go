package main

import (
	"bufio"
	"fmt"
	"os"
)

type game struct {
	x, y    int
	height  int
	width   int
	running bool
}

var key string // keyboard input (w, s, a, d, q)

func (g *game) update(key rune) {
	if key == ' ' {
		return
	}

	switch key {
	case 'w': //up
		if g.x > 0 {
			g.x--
		}
	case 's': //down
		if g.x < g.height-1 {
			g.x++
		}
	case 'a': //left
		if g.y > 0 {
			g.y--
		}
	case 'd': //right
		if g.y < g.width-1 {
			g.y++
		}
	case 'q': //quit
		g.running = false
	}
}

func (g *game) display() { // display the map
	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			if i == g.x && j == g.y {
				fmt.Print("O") // player
			} else {
				fmt.Print(".") // empty
			}
		}
		fmt.Println() // new line
	}
}
func read_input() rune {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		key = scanner.Text()
		if key != "" && key[0] != ' ' {
			break
		}
	}
	return rune(key[0])
}

func main() {
	g := game{0, 0, 10, 20, true} // x,y,height,width,running
	for g.running {
		g.display()
		key := read_input()
		g.update(key)
	}
}
