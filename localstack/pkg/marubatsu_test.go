package marubatsh

import (
	"testing"
)

func TestTicTacToeBoard_IsFinishedByCol(t *testing.T) {
    board := New(3)
    // "o" "x" "x"
    // "o" "x" "x"
    // "o" "x" "x"
    board.Put(0, 0, "o")
    board.Put(1, 0, "o")
    board.Put(2, 0, "o")
    if !board.IsFinishedByCol() {
        t.Fatal("IsFinishedByCol should be true")
    }
}

func TestTicTacToeBoard_IsFinishedByRow(t *testing.T) {
    board := New(3)
    // "o" "o" "o"
    // "x" "x" "x"
    // "x" "x" "x"
    board.Put(0, 0, "o")
    board.Put(0, 1, "o")
    board.Put(0, 2, "o")
    if !board.IsFinishedByRow() {
        t.Fatal("IsFinishedByRow should be true")
    }
}


func TestTicTacToeBoard_IsFinishedByDiagonal(t *testing.T) {
    board := New(3)
    // "o" "x" "x"
    // "x" "o" "x"
    // "x" "x" "o"
    board.Put(0, 0, "o")
    board.Put(1, 1, "o")
    board.Put(2, 2, "o")
    if !board.IsFinishedByDiagonal() {
        t.Fatal("IsFinishedByDiagonal should be true")
    }

    board = New(3)
    // "x" "x" "o"
    // "x" "o" "x"
    // "o" "x" "x"
    board.Put(0, 2, "o")
    board.Put(1, 1, "o")
    board.Put(2, 0, "o")
    if !board.IsFinishedByDiagonal() {
        t.Fatal("IsFinishedByDiagonal should be true")
    }
    
}