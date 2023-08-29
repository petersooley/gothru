package maze

import "math/rand"

var r = rand.New(rand.NewSource(11111))

type Path map[Cell]map[Cell]bool

type Cell struct {
	X, Y int
}

func valid_cell_or_nil(mazeSize, x, y int) *Cell {
	if x < mazeSize && x >= 0 && y < mazeSize && y >= 0 {
		return &Cell{x, y}
	}
	return nil
}
func (c *Cell) North(mazeSize int) *Cell {
	return valid_cell_or_nil(mazeSize, c.X, c.Y+1)
}
func (c *Cell) South(mazeSize int) *Cell {
	return valid_cell_or_nil(mazeSize, c.X, c.Y-1)
}
func (c *Cell) East(mazeSize int) *Cell {
	return valid_cell_or_nil(mazeSize, c.X+1, c.Y)
}
func (c *Cell) West(mazeSize int) *Cell {
	return valid_cell_or_nil(mazeSize, c.X-1, c.Y)

}

type Maze struct {
	size    int
	visited map[Cell]bool
	path    Path
	current []Cell
	r       *rand.Rand
}

func Generate(size int, seed int64) Maze {
	maze := Maze{
		size:    size,
		visited: make(map[Cell]bool),
		path:    make(Path),
		current: make([]Cell, 0),
	}

	if seed > 0 {
		maze.r = rand.New(rand.NewSource(seed))
	} else {
		maze.r = rand.New(rand.NewSource(rand.Int63()))
	}

	maze.push(Cell{})

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

	m.visited[c] = true
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
	neighbors := []*Cell{c.North(m.size), c.South(m.size), c.East(m.size), c.West(m.size)}
	m.r.Shuffle(4, func(i, j int) { neighbors[i], neighbors[j] = neighbors[j], neighbors[i] })
	for _, n := range neighbors {
		if n == nil {
			continue
		}
		if _, ok := m.visited[*n]; ok {
			continue
		}
		return n
	}
	return nil
}
