package marubatsh

import "fmt"




type TicTacToeBoard struct {
    Board [][]string
    n int
}

func New(n int)*TicTacToeBoard {

    board := [][]string{}

    for i := 0; i < n; i++ {
        row := []string{}
        for j := 0; j < n; j++ {
            row = append(row, "")
        }
        board = append(board, row)
    }
    return &TicTacToeBoard{
        Board: board,
        n: n,
    }
}

func (b *TicTacToeBoard) Put(x int, y int, value string) {
    b.Board[x][y] = value
}

func (b *TicTacToeBoard) Get(x int, y int) string {
    return b.Board[x][y]
}


func(b *TicTacToeBoard) IsFinished()bool {
    return b.IsFinishedByRow() || b.IsFinishedByCol() || b.IsFinishedByDiagonal()
}

func (b *TicTacToeBoard) IsFinishedByRow() bool {
    for i := 0; i < b.n; i++ {
        if b.Board[i][0] == "" {
            continue
        }
        for j := 1; j < b.n; j++ {
            if b.Board[i][j] != b.Board[i][0] {
                break
            }
            if j == b.n - 1 {
                return true
            }
        }
    }
    return false
}

func (b *TicTacToeBoard) IsFinishedByCol() bool {
    for i := 0; i < b.n; i++ {
        if b.Board[0][i] == "" {
            continue
        }
        for j := 1; j < b.n; j++ {
            if b.Board[j][i] != b.Board[0][i] {
                break
            }
            if j == b.n - 1 {
                return true
            }
        }
    }
    return false
}

func (b *TicTacToeBoard) IsFinishedByDiagonal() bool {
    topLeft := b.Board[0][0]
    if topLeft != "" {
        for i := 1; i < b.n; i++ {
            if b.Board[i][i] != topLeft {
                break
            }
            if i == b.n - 1 {
                return true
            }
        }
    }
    topRight := b.Board[0][b.n - 1]
    if topRight != "" {
        for i := 1; i < b.n; i++ {
            if b.Board[i][b.n - 1 - i] != topRight {
                break
            }
            if i == b.n - 1 {
                return true
            }
        }
    }
    return false
}

func TicTacToe() {
    board := New(3)
    board.Put(0, 0, "X")
    board.Put(0, 1, "X")
    board.Put(0, 2, "X")
    fmt.Println(board.IsFinished())
}