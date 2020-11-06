package backtracking

import "testing"

func TestExist(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	got := exist(board, word)
	want := true
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	word = "ABCB"
	got = exist(board, word)
	want = false
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	board = [][]byte{{'C','A','A'},{'A','A','A'},{'B','C','D'}}
	word = "AAB"
	got = exist(board, word)
	want = true
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
