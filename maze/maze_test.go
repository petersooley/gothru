package maze_test

import (
	"elmdash/gothru/maze"
	"testing"
)

func TestMaze(t *testing.T) {
	m := maze.Generate(5, 11111)

	// given the size and seed, we know these should be the right connections
	// from drawing the maze on paper
	expectConnection(t, &m, 0, 5)
	expectConnection(t, &m, 5, 10)
	expectConnection(t, &m, 10, 11)
	expectConnection(t, &m, 11, 6)
	expectConnection(t, &m, 6, 1)
	expectConnection(t, &m, 1, 2)
	expectConnection(t, &m, 2, 3)
	expectConnection(t, &m, 3, 8)
	expectConnection(t, &m, 8, 7)
	expectConnection(t, &m, 7, 12)
	expectConnection(t, &m, 12, 13)
	expectConnection(t, &m, 13, 18)
	expectConnection(t, &m, 18, 19)
	expectConnection(t, &m, 19, 14)
	expectConnection(t, &m, 14, 9)
	expectConnection(t, &m, 9, 4)
	expectConnection(t, &m, 19, 24)
	expectConnection(t, &m, 24, 23)
	expectConnection(t, &m, 23, 22)
	expectConnection(t, &m, 22, 21)
	expectConnection(t, &m, 21, 20)
	expectConnection(t, &m, 20, 15)
	expectConnection(t, &m, 15, 16)
	expectConnection(t, &m, 16, 17)
}

func expectConnection(t *testing.T, m *maze.Maze, from, to int) {
	if !m.IsPath(from, to) {
		t.Fatalf("%v not connected to %v", from, to)
	}
}

func TestMazeImage(t *testing.T) {
	// m := maze.Generate(5, 11111)
	// _ = m.Image()
}
