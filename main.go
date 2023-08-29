package main

import (
	"fmt"
	"math/rand"
)

const MAZE_SIZE int = 5

var r = rand.New(rand.NewSource(11111))

type Cell struct {
	x, y int
}

func valid_cell_or_nil(x, y int) *Cell {
	if x < MAZE_SIZE && x >= 0 && y < MAZE_SIZE && y >= 0 {
		return &Cell{x, y}
	}
	return nil
}
func (c *Cell) North() *Cell {
	return valid_cell_or_nil(c.x, c.y+1)
}
func (c *Cell) South() *Cell {
	return valid_cell_or_nil(c.x, c.y-1)
}
func (c *Cell) East() *Cell {
	return valid_cell_or_nil(c.x+1, c.y)
}
func (c *Cell) West() *Cell {
	return valid_cell_or_nil(c.x-1, c.y)

}

type Maze struct {
	visited map[Cell]bool
	path    map[Cell]map[Cell]bool
	current []Cell
}

func makeMaze() Maze {
	maze := Maze{
		visited: make(map[Cell]bool),
		path:    make(map[Cell]map[Cell]bool),
		current: make([]Cell, 0),
	}

	maze.push(Cell{})

	return maze
}

func (m *Maze) push(c Cell) {
	fmt.Println("visiting", c)
	prev := m.currentCell()
	if prev != nil {
		fmt.Printf("connecting %v to %v\n", prev, c)
		if m.path[*prev] == nil {
			m.path[*prev] = make(map[Cell]bool)
		}
		m.path[*prev][c] = true
	}

	m.visited[c] = true
	m.current = append(m.current, c)

	if len(m.visited) > MAZE_SIZE*MAZE_SIZE {
		fmt.Println("all visited")
		return
	}

	next := m.nextUnvisitedNeighbor()
	fmt.Println("next", next)
	if next == nil {
		n := m.popCurrent()
		if len(m.current) == 0 {
			return
		}
		fmt.Println("unwinding current", m.current)
		m.push(n)
	} else {
		m.push(*next)
	}
	return
}

func (m *Maze) currentCell() *Cell {
	if len(m.current) > 0 {
		return &m.current[len(m.current)-1]
	}
	return nil
}

func (m *Maze) popCurrent() (c Cell) {
	fmt.Println("pop current: before", m.current)
	c, m.current = m.current[len(m.current)-2], m.current[:len(m.current)-2]
	fmt.Println("pop current: after", m.current)
	return
}

func (m *Maze) nextUnvisitedNeighbor() *Cell {
	c := m.currentCell()
	neighbors := []*Cell{c.North(), c.South(), c.East(), c.West()}
	r.Shuffle(4, func(i, j int) { neighbors[i], neighbors[j] = neighbors[j], neighbors[i] })
	for _, n := range neighbors {
		fmt.Println("neighbor", n)
		if n == nil {
			fmt.Println("  is nil")
			continue
		}
		if _, ok := m.visited[*n]; ok {
			fmt.Println("  was visited already")
			continue
		}
		fmt.Println("  is valid")
		return n
	}
	fmt.Println("no valid neighbors")
	return nil
}

func main() {
	maze := makeMaze()
	for from, to := range maze.path {
		for t := range to {
			fmt.Println("from", from, "to", t)
		}
	}

}
