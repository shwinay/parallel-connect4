package main

import (
	"fmt"
	"math"
	"math/rand"
)

func stateHeuristic(board [][]int, player int) int {
	numThrees := 0
	checkRow := func(board [][]int, row, player int) int {
		threes := 0
		count := 0
		for i := 0; i < len(board[row]); i++ {
			if board[row][i] == player {
				count++
			} else {
				count = 0
			}
			if count >= 3 {
				threes++
				count = 0
			}
		}
		return threes
	}
	checkCol := func(board [][]int, col, player int) int {
		threes := 0
		count := 0
		for i := 0; i < len(board); i++ {
			if board[i][col] == player {
				count++
			} else {
				count = 0
			}
			if count >= 3 {
				threes++
				count = 0
			}
		}
		return threes
	}
	// checkUpperLeftDiagonal := func(board [][]int, lowerLeftI, lowerLeftJ, player int) int {
	// 	threes := 0
	// 	count := 0
	// 	for i, j := lowerLeftI, lowerLeftJ; i >= 0 && j < len(board[0]); i, j = i-1, j+1 {
	// 		if board[i][j] == player {
	// 			count++
	// 		} else {
	// 			count = 0
	// 		}
	// 		if count >= 3 {
	// 			threes++
	// 			count = 0
	// 		}
	// 	}
	// 	return threes
	// }
	// checkLowerRightDiagonal := func(board [][]int, upperLeftI, upperLeftJ, player int) int {
	// 	threes := 0
	// 	count := 0
	// 	for i, j := upperLeftI, upperLeftJ; i < len(board) && j < len(board[0]); i, j = i+1, j+1 {
	// 		if board[i][j] == player {
	// 			count++
	// 		} else {
	// 			count = 0
	// 		}
	// 		if count >= 3 {
	// 			threes++
	// 			count = 0
	// 		}
	// 	}
	// 	return threes
	// }

	//check rows
	for i := 0; i < len(board); i++ {
		numThrees += checkRow(board, i, player)
	}
	//check cols
	for i := 0; i < len(board[0]); i++ {
		numThrees += checkCol(board, i, player)
	}
	// //check upper left diagonal, starting from bottom right corner
	// for j := len(board[0]) - 1; j >= 0; j-- {
	// 	numThrees += checkUpperLeftDiagonal(board, len(board)-1, j, player)
	// }
	// for i := len(board) - 2; i >= 0; i-- {
	// 	numThrees += checkUpperLeftDiagonal(board, i, 0, player)
	// }

	// //check bottom right diagonal, starting from top right corner
	// for j := len(board[0]) - 1; j >= 0; j-- {
	// 	numThrees += checkLowerRightDiagonal(board, 0, j, player)
	// }
	// for i := 1; i < len(board); i++ {
	// 	numThrees += checkLowerRightDiagonal(board, i, 0, player)
	// }

	return numThrees
}

//minimax AI - returns (maxMove, maxValue)
func sMinimax(board [][]int, player, depth int) (int, int) {

	otherPlayer := BLANK
	if player == PLAYER1 {
		otherPlayer = PLAYER2
	} else {
		otherPlayer = PLAYER1
	}

	//terminal states
	if isWin(board, player) { //win
		//fmt.Println("found winning state for player", player)
		return 0, math.MaxInt32
	} else if isWin(board, otherPlayer) { //win for other player
		//fmt.Println("BAD: found winning state for player", player)
		return 0, math.MinInt32
	} else if boardFull(board) { //draw
		//fmt.Println("found board full state..")
		return 0, 0
	} else if depth <= 0 { //static evaluation function, depth limit reached
		//fmt.Println("hit depth limit, state heuristic:", stateHeuristic(board, player))
		return -1, stateHeuristic(board, player)
	}

	if player == PLAYER2 { //maximizing agent - AI
		bestMove := -1
		maxVal := math.MinInt32
		//simulate dropping piece into column i
		for i := 0; i < len(board[0]); i++ {
			newBoard := copyBoard(board)
			if !insertPiece(newBoard, i, player) {
				continue
			}
			_, val := sMinimax(newBoard, otherPlayer, depth)
			if val >= maxVal {
				bestMove = i
				maxVal = val
			}
		}
		return bestMove, maxVal
	} else if player == PLAYER1 { //minimizing agent - human player
		bestMove := -1
		minVal := math.MaxInt32
		//simulate dropping piece into column i
		for i := 0; i < len(board[0]); i++ {
			newBoard := copyBoard(board)
			if !insertPiece(newBoard, i, player) {
				continue
			}
			_, val := sMinimax(newBoard, otherPlayer, depth-1)
			if val <= minVal {
				bestMove = i
				minVal = val
			}
		}
		return bestMove, minVal
	} else {
		fmt.Println("not minimizing or maximizing agent.. this should never happen")
		return -1, -1
	}
}

func randomAI(board [][]int) int {
	return rand.Intn(len(board[0]))
}

func playAI(board [][]int) int {
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
		if player == PLAYER1 {
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
			player = PLAYER2
			otherPlayer = PLAYER1
		} else {
			aiCol, _ := sMinimax(board, PLAYER2, 4)
			fmt.Println("AI dropped piece in column", aiCol)
			if !insertPiece(board, aiCol, player) {
				return BLANK
			}
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
