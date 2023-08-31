package maze_test

import (
	"elmdash/gothru/maze"
	"testing"
)

func TestMaze(t *testing.T) {
	_ = maze.Generate(5, 11111)

	// path := m.Path()
	// if len(path) == 0 {
	// 	t.Fatal("no path generated")
	// }
	// // given the size and seed, we know these should be the right connections
	// // from drawing the maze on paper
	// expectConnection(t, path, maze.Cell{0, 3}, maze.Cell{0, 4})
	// expectConnection(t, path, maze.Cell{2, 2}, maze.Cell{1, 2})
	// expectConnection(t, path, maze.Cell{1, 2}, maze.Cell{1, 1})
	// expectConnection(t, path, maze.Cell{1, 2}, maze.Cell{1, 3})
	// expectConnection(t, path, maze.Cell{2, 1}, maze.Cell{2, 0})
	// expectConnection(t, path, maze.Cell{3, 1}, maze.Cell{3, 2})
	// expectConnection(t, path, maze.Cell{4, 2}, maze.Cell{4, 3})
	// expectConnection(t, path, maze.Cell{4, 0}, maze.Cell{4, 1})
	// expectConnection(t, path, maze.Cell{0, 1}, maze.Cell{0, 2})
	// expectConnection(t, path, maze.Cell{0, 2}, maze.Cell{0, 3})
	// expectConnection(t, path, maze.Cell{0, 4}, maze.Cell{1, 4})
	// expectConnection(t, path, maze.Cell{1, 4}, maze.Cell{2, 4})
	// expectConnection(t, path, maze.Cell{2, 4}, maze.Cell{3, 4})
	// expectConnection(t, path, maze.Cell{2, 0}, maze.Cell{1, 0})
	// expectConnection(t, path, maze.Cell{2, 0}, maze.Cell{3, 0})
	// expectConnection(t, path, maze.Cell{3, 3}, maze.Cell{2, 3})
	// expectConnection(t, path, maze.Cell{2, 3}, maze.Cell{2, 2})
	// expectConnection(t, path, maze.Cell{1, 1}, maze.Cell{2, 1})
	// expectConnection(t, path, maze.Cell{4, 1}, maze.Cell{3, 1})
	// expectConnection(t, path, maze.Cell{3, 2}, maze.Cell{4, 2})
	// expectConnection(t, path, maze.Cell{4, 3}, maze.Cell{4, 4})
	// expectConnection(t, path, maze.Cell{0, 0}, maze.Cell{0, 1})
	// expectConnection(t, path, maze.Cell{3, 4}, maze.Cell{3, 3})
	// expectConnection(t, path, maze.Cell{3, 0}, maze.Cell{4, 0})
}

func expectConnection(t *testing.T, path maze.Path, from, to maze.Cell) {
	if _, ok := path[from]; !ok {
		t.Fatalf("from %v not found", from)
	}
	for t := range path[from] {
		if t == to {
			return
		}
	}
	t.Fatalf("from %v not connected to %v", from, to)
}


func TestMazeImage(t *testing.T) {
	// m := maze.Generate(5, 11111)
	// _ = m.Image()
}