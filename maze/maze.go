package maze

import (
	"fmt"
	"math/rand"
)

var r = rand.New(rand.NewSource(11111))

type Path map[Cell]map[Cell]bool

type Cell struct {
	I     int
	X, Y  int
	North *int
	South *int
	East  *int
	West  *int
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
		I:     *I,
		X:     x,
		Y:     y,
		North: rowMajor(mazeSize, x, y-1),
		South: rowMajor(mazeSize, x, y+1),
		East:  rowMajor(mazeSize, x-1, y),
		West:  rowMajor(mazeSize, x+1, y),
	}
}
func (c *Cell) String() string {
	safeIntPrint := func(i *int) string {
		if i == nil {
			return "∅"
		}
		return fmt.Sprintf("%d", *i)
	}
	return fmt.Sprintf(
		"{i: %v, x: %v, y: %v, nbrs: [%v, %v, %v, %v]}\n",
		c.I, c.X, c.Y,
		safeIntPrint(c.North), safeIntPrint(c.South), safeIntPrint(c.East), safeIntPrint(c.West),
	)
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
	neighbors := []*int{c.North, c.South, c.East, c.West}
	m.r.Shuffle(4, func(i, j int) { neighbors[i], neighbors[j] = neighbors[j], neighbors[i] })
	for _, n := range neighbors {
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
