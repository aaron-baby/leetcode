package union_find

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	cnt := numIslands(grid)
	want := 1
	if cnt != want {
		t.Errorf("got %v want %v given", cnt, want)
	}

	grid = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	cnt = numIslands(grid)
	want = 3
	if cnt != want {
		t.Errorf("got %v want %v given", cnt, want)
	}

	grid = [][]byte{}
	cnt = numIslands(grid)
	want = 0
	if cnt != want {
		t.Errorf("got %v want %v given", cnt, want)
	}
}
