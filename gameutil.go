package main

import (
	"flag"
	"fmt"

	"github.com/logrusorgru/aurora"
)

//BLANK blank board space
const BLANK = 0

//PLAYER1 player 1 piece
const PLAYER1 = 1

//PLAYER2 player 2 piece
const PLAYER2 = 2

var au aurora.Aurora

func initBoard(width, height int) [][]int {
	var colors = flag.Bool("colors", true, "enable or disable colors")
	au = aurora.NewAurora(*colors)
	board := make([][]int, height)
	for i := 0; i < len(board); i++ {
		board[i] = make([]int, width)
	}
	return board
}

func printBoard(board [][]int) {
	//print row numbers
	for i := 0; i < len(board[0]); i++ {
		fmt.Print("", au.Blue(i), " ")
	}
	fmt.Println()
	fmt.Println()

	//print board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			piece := board[i][j]
			if piece == BLANK {
				fmt.Print("-", " ")
			} else if piece == PLAYER1 {
				fmt.Print("", au.Red("o"), " ")
			} else if piece == PLAYER2 {
				fmt.Print("", au.Yellow("o"), " ")
			}
		}
		fmt.Println()
	}
}

func insertPiece(board [][]int, col, piece int) bool {
	row := len(board) - 1
	for row >= 0 && board[row][col] != 0 {
		row--
	}
	if row < 0 {
		return false
	}

	board[row][col] = piece
	return true
}

func boardFull(board [][]int) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == BLANK {
				return false
			}
		}
	}
	return true
}

func isWin(board [][]int, player int) bool {
	checkRow := func(board [][]int, row, player int) bool {
		count := 0
		for i := 0; i < len(board[row]); i++ {
			if board[row][i] == player {
				count++
			} else {
				count = 0
			}
			if count >= 4 {
				return true
			}
		}
		return false
	}
	checkCol := func(board [][]int, col, player int) bool {
		count := 0
		for i := 0; i < len(board); i++ {
			if board[i][col] == player {
				count++
			} else {
				count = 0
			}
			if count >= 4 {
				return true
			}
		}
		return false
	}
	checkDiagonals := func(board [][]int, row, col, player int) bool {
		//check bottom left
		count := 0
		for i, j := row, col; i < len(board) && j >= 0; i, j = i+1, j-1 {
			if board[i][j] == player {
				count++
			} else {
				count = 0
			}
			if count >= 4 {
				return true
			}
		}
		//check bottom right
		count = 0
		for i, j := row, col; i < len(board) && j < len(board[0]); i, j = i+1, j+1 {
			if board[i][j] == player {
				count++
			} else {
				count = 0
			}
			if count >= 4 {
				return true
			}
		}
		return false
	}

	//check rows
	for i := 0; i < len(board); i++ {
		if checkRow(board, i, player) {
			return true
		}
	}

	//check cols
	for i := 0; i < len(board[0]); i++ {
		if checkCol(board, i, player) {
			return true
		}
	}

	//TODO: CHECK DIAGONALS
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if checkDiagonals(board, i, j, player) {
				return true
			}
		}
	}

	return false
}

func copyBoard(board [][]int) [][]int {
	newBoard := make([][]int, len(board))
	for i := 0; i < len(newBoard); i++ {
		newBoard[i] = make([]int, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			newBoard[i][j] = board[i][j]
		}
	}

	return newBoard
}
