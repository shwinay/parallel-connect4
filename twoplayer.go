package main

import (
	"fmt"
)

func playTwoPlayer(board [][]int) int {
	player := PLAYER1
	otherPlayer := -1
	if player == PLAYER1 {
		otherPlayer = PLAYER2
	} else {
		otherPlayer = PLAYER1
	}
	for !(boardFull(board) || isWin(board, otherPlayer)) {
		printBoard(board)
		fmt.Println()
		fmt.Println("Player", player, "enter column: ")
		var col int
		_, err := fmt.Scanf("%d", &col)
		for err != nil || col < 0 || col >= len(board[0]) {
			fmt.Println("Try again: ")
			_, err = fmt.Scanf("\n%d", &col)
		}

		if !insertPiece(board, col, player) {
			return BLANK
		}
		if player == PLAYER1 {
			player = PLAYER2
			otherPlayer = PLAYER1
		} else {
			player = PLAYER1
			otherPlayer = PLAYER2
		}
		fmt.Println()
	}
	if boardFull(board) {
		return BLANK
	}
	return otherPlayer
}

func endGame(board [][]int, player int) {
	printBoard(board)
	fmt.Println()
	if player == BLANK {
		fmt.Println("TIE!")
		fmt.Println("Winner: everyone :)")
	} else {
		fmt.Println("CONNECT 4!")
		fmt.Println("Winner: Player", player)
	}
}
