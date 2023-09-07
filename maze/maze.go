package maze

import (
	"fmt"
	"math/rand"
)

var r = rand.New(rand.NewSource(11111))

type Path map[Cell]map[Cell]bool

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
			rowMajor(mazeSize, x-1, y),
			rowMajor(mazeSize, x+1, y),
		),
		Paths: NewSurrounding[bool](false, false, false, false),
	}
}
func (c *Cell) String() string {
	return fmt.Sprintf("[%v] %v,%v, nbrs: %v}\n", c.I, c.X, c.Y, c.Neighbors)
}

func (c *Cell) Connect(to *Cell) {
	switch to.I {
	case c.Neighbors.North:
		c.Paths.North = true
	case c.Neighbors.South:
		c.Paths.South = true
	case c.Neighbors.East:
		c.Paths.East = true
	case c.Neighbors.West:
		c.Paths.West = true
	default:
		panic(fmt.Sprintf("%v,%v is not a neighbor of %v,%v", c.X, c.Y, to.X, to.Y))
	}
}

type Maze struct {
	cells   map[int]*Cell
	size    int
	visited map[int]bool
	path    Path
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

	fmt.Println(cells)

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

	return maze
}

func (m *Maze) Path() Path {
	return m.path
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
		m.stack = append(m.stack, c.I)
	}
	m.visit()
	return
}

func (m *Maze) currentCell() *Cell {
	return m.cells[m.stack[len(m.stack)-1]]
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
