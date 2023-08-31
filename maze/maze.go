package maze

import (
	"fmt"
	"math/rand"
)

var r = rand.New(rand.NewSource(11111))

type Path map[Cell]map[Cell]bool

type Surrounding[T interface{}] struct {
	North *T
	South *T
	East  *T
	West  *T
}

func NewSurrounding[T interface{}](North *T, South *T, East *T, West *T) Surrounding[T] {
	return Surrounding[T]{North, South, East, West}
}

func (s Surrounding[T]) Shuffled(r *rand.Rand) []*T {
	arr := []*T{s.North, s.South, s.East, s.West}
	r.Shuffle(4, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}

func (s Surrounding[T]) String() string {
	safePrint := func(i *T) string {
		if i == nil {
			return "∅"
		}
		return fmt.Sprintf("%v", *i)
	}
	return fmt.Sprintf(
		"[n: %v, s: %v, e: %v, w: %v]",
		safePrint(s.North),
		safePrint(s.South),
		safePrint(s.East),
		safePrint(s.West),
	)
}

type Cell struct {
	I         int
	X, Y      int
	Neighbors Surrounding[int]
}

func rowMajor(mazeSize, x, y int) *int {
	if x < 0 || y < 0 || x >= mazeSize || y >= mazeSize {
		return nil
	}

	i := (y * mazeSize) + x
	return &i
}

func NewCell(mazeSize, x, y int) Cell {
	I := rowMajor(mazeSize, x, y)
	if I == nil {
		panic(fmt.Sprintf("invalid cell coords %v,%v", x, y))
	}

	return Cell{
		I: *I,
		X: x,
		Y: y,
		Neighbors: NewSurrounding[int](
			rowMajor(mazeSize, x, y-1),
			rowMajor(mazeSize, x, y+1),
			rowMajor(mazeSize, x-1, y),
			rowMajor(mazeSize, x+1, y),
		),
	}
}
func (c *Cell) String() string {
	return fmt.Sprintf("{i: %v, x: %v, y: %v, nbrs: %v}\n", c.I, c.X, c.Y, c.Neighbors)
}

type Maze struct {
	cells   map[int]*Cell
	size    int
	visited map[int]bool
	path    Path
	current []Cell
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

	maze := Maze{
		cells:   cells,
		size:    size,
		visited: make(map[int]bool),
		path:    make(Path),
		current: make([]Cell, 0),
	}

	if seed > 0 {
		maze.r = rand.New(rand.NewSource(seed))
	} else {
		maze.r = rand.New(rand.NewSource(rand.Int63()))
	}

	// maze.push(Cell{})

	return maze
}

func (m *Maze) Path() Path {
	return m.path
}

func (m *Maze) push(c Cell) {
	prev := m.currentCell()
	if prev != nil {
		if m.path[*prev] == nil {
			m.path[*prev] = make(map[Cell]bool)
		}
		m.path[*prev][c] = true
	}

	m.visited[c.I] = true
	m.current = append(m.current, c)

	next := m.nextUnvisitedNeighbor()
	if next == nil {
		n := m.popCurrent()
		if len(m.current) == 0 {
			return
		}
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
	// pops two elements, this works because we prematurely add to current for each visit
	c, m.current = m.current[len(m.current)-2], m.current[:len(m.current)-2]
	return
}

func (m *Maze) nextUnvisitedNeighbor() *Cell {
	c := m.currentCell()
	for _, n := range c.Neighbors.Shuffled(m.r) {
		if n == nil {
			continue
		}
		if _, ok := m.visited[*n]; ok {
			continue
		}
		return m.cells[*n]
	}
	return nil
}
