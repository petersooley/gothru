package main

import (
	"elmdash/gothru/maze"
	"fmt"
)

func main() {
	maze := maze.Generate(5, 11111)
	for from, to := range maze.Path() {
		for t := range to {
			fmt.Printf("expectConnection(t, path, maze.Cell{%v, %v}, maze.Cell{%v, %v})\n", from.X, from.Y, t.X, t.Y)
		}
	}
	fmt.Println("hello")
}
