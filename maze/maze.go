package maze

import (
	"fmt"
	"math/rand"
)

const OUT_OF_MAZE int = -1

type Surrounding[T interface{}] struct {
	North T
	South T
	East  T
	West  T
}

func NewSurrounding[T interface{}](North T, South T, East T, West T) Surrounding[T] {
	return Surrounding[T]{North, South, East, West}
}

func (s Surrounding[T]) Shuffled(r *rand.Rand) []T {
	arr := []T{s.North, s.South, s.East, s.West}
	r.Shuffle(4, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}

type Cell struct {
	I         int
	X, Y      int
	Neighbors Surrounding[int]
	Paths     Surrounding[bool]
}

func rowMajor(mazeSize, x, y int) int {
	if x < 0 || y < 0 || x >= mazeSize || y >= mazeSize {
		return OUT_OF_MAZE
	}

	return (y * mazeSize) + x
}

func NewCell(mazeSize, x, y int) Cell {
	I := rowMajor(mazeSize, x, y)
	if I == OUT_OF_MAZE {
		panic(fmt.Sprintf("invalid cell coords %v,%v", x, y))
	}

	return Cell{
		I: I,
		X: x,
		Y: y,
		Neighbors: NewSurrounding[int](
			rowMajor(mazeSize, x, y-1),
			rowMajor(mazeSize, x, y+1),
			rowMajor(mazeSize, x+1, y),
			rowMajor(mazeSize, x-1, y),
		),
		Paths: NewSurrounding[bool](false, false, false, false),
	}
}
func (c *Cell) String() string {
	return fmt.Sprintf("[Cell #%02v] %v,%v, nbrs: %v, pth: %v}\n", c.I, c.X, c.Y, &c.Neighbors, &c.Paths)
}
func (s *Surrounding[T]) String() string {
	return fmt.Sprintf("{n:%v, s:%v, e:%v, w:%v}", s.North, s.South, s.East, s.West)
}

func (c *Cell) Connect(to *Cell) {
	// fmt.Printf("connect %v (%v, %v) to %v (%v, %v)\n", c.I, c.X, c.Y, to.I, to.X, to.Y)

	switch to.I {
	case c.Neighbors.North:
		to.Paths.South = true
		c.Paths.North = true
	case c.Neighbors.South:
		to.Paths.North = true
		c.Paths.South = true
	case c.Neighbors.East:
		to.Paths.West = true
		c.Paths.East = true
	case c.Neighbors.West:
		to.Paths.East = true
		c.Paths.West = true
	default:
		panic(fmt.Sprintf("%v,%v is not a neighbor of %v,%v", c.X, c.Y, to.X, to.Y))
	}
}

type Maze struct {
	cells   map[int]*Cell
	size    int
	visited map[int]bool
	stack   []int
	r       *rand.Rand
}

func Generate(size int, seed int64) Maze {
	cells := make(map[int]*Cell, size*size)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			c := NewCell(size, x, y)
			cells[c.I] = &c
		}
	}

	// todo maybe we start in the middle of the maze instead?

	maze := Maze{
		cells:   cells,
		size:    size,
		visited: make(map[int]bool),
		stack:   make([]int, 1),
	}

	if seed > 0 {
		maze.r = rand.New(rand.NewSource(seed))
	} else {
		maze.r = rand.New(rand.NewSource(rand.Int63()))
	}

	maze.visit()

	// for _, cell := range maze.cells {
	// 	fmt.Printf("%v\n", cell)
	// }

	return maze
}

func (m *Maze) visit() {
	c := m.currentCell()

	m.visited[c.I] = true

	next := m.nextUnvisitedNeighbor(c)

	if next == nil {
		m.stack = m.stack[:len(m.stack)-1]
		if len(m.stack) == 0 {
			return
		}
	} else {
		c.Connect(next)
		m.stack = append(m.stack, next.I)
	}
	m.visit()
}

func (m *Maze) currentCell() *Cell {
	return m.cells[m.stack[len(m.stack)-1]]
}

func (m *Maze) IsPath(a, b int) bool {
	aCell := m.cells[a]
	switch b {
	case aCell.Neighbors.North:
		return aCell.Paths.North
	case aCell.Neighbors.South:
		return aCell.Paths.South
	case aCell.Neighbors.East:
		return aCell.Paths.East
	case aCell.Neighbors.West:
		return aCell.Paths.West
	default:
		panic(fmt.Sprintf("%v is not a neighbor of %v", a, b))
	}
}

func (m *Maze) nextUnvisitedNeighbor(c *Cell) *Cell {
	for _, n := range c.Neighbors.Shuffled(m.r) {
		if n == OUT_OF_MAZE {
			continue
		}
		if _, ok := m.visited[n]; ok {
			continue
		}
		return m.cells[n]
	}
	return nil
}
