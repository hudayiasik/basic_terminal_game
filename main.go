package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math/rand"
	"os"
)

type snake struct {
	x, y int
}

type game struct {
	snake     *list.List
	food      *food
	height    int
	width     int
	running   bool
	direction int // 0-up, 1-down, 2-left, 3-right

}
type food struct {
	x, y int
}

func (g *game) initFood() {
	// Initialize the food struct with random x and y positions within the grid
	g.food = &food{rand.Intn(g.height), rand.Intn(g.width)}
}

func (g *game) checkFood() {
	// check if snake's head is on the food
	if g.snake.Front().Value.(*snake).x == g.food.x && g.snake.Front().Value.(*snake).y == g.food.y {
		// Add a new element to the snake's body
		g.snake.PushBack(&snake{g.snake.Back().Value.(*snake).x, g.snake.Back().Value.(*snake).y})
		// Generate a new food at a random position
		g.initFood()
	}
}

var key string // keyboard input (w, s, a, d, q)

func (g *game) setDirection(key rune) {
	if key == ' ' {
		return
	}
	switch key {
	case 'w': //up
		if g.direction != 1 {
			g.direction = 0
		}

	case 's': //down
		if g.direction != 0 {
			g.direction = 1
		}

	case 'a': //left
		if g.direction != 3 {
			g.direction = 2
		}
	case 'd': //right
		if g.direction != 2 {
			g.direction = 3
		}
	case 'q': //quit
		g.running = false
	}

}

func (g *game) moveTails() {
	// Iterate through the elements in the list, starting from the second element
	for e := g.snake.Front().Next(); e != nil; e = e.Next() {
		// Set the current element's x and y to the x and y of the previous element
		e.Value.(*snake).x = e.Prev().Value.(*snake).x
		e.Value.(*snake).y = e.Prev().Value.(*snake).y
	}
}

func (g *game) update() {

	dir := g.direction

	g.moveTails()
	switch dir {
	case 0: //up
		// game.snake.Front
		if g.snake.Front().Value.(*snake).x > 0 {
			g.snake.Front().Value.(*snake).x -= 1
		}
	case 1: //down
		if g.snake.Front().Value.(*snake).x < g.height-1 {
			g.snake.Front().Value.(*snake).x += 1
		}
	case 2: //left
		if g.snake.Front().Value.(*snake).y > 0 {
			g.snake.Front().Value.(*snake).y -= 1
		}

	case 3: //right
		if g.snake.Front().Value.(*snake).y < g.width-1 {
			g.snake.Front().Value.(*snake).y += 1
		}
	}

}
func elementInList(l *list.List, x int, y int) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value.(*snake).x == x && e.Value.(*snake).y == y {
			return true
		}
	}
	return false
}
func (g *game) display() {
	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			if elementInList(g.snake, i, j) {
				fmt.Print("O") // snake
			} else if i == g.food.x && j == g.food.y {
				fmt.Print("X") // food
			} else {
				fmt.Print(".") // empty
			}
		}
		fmt.Println() // new line
	}
}

func readInput() rune {
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
	s := snake{0, 3}
	f := food{0, 0}
	g := game{list.New(), &f, 10, 10, true, 2}
	g.snake.PushBack(&s) // head
	g.initFood()

	for g.running {
		g.display()
		key := readInput()
		g.setDirection(key)
		g.checkFood()
		g.update()
	}
}
