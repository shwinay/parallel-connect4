package main

import (
	"fmt"
	"math"
	"math/rand"
)

//MaxDepth : maximum depth of minimax search
const MaxDepth = 6

func stateHeuristic(board [][]int, player int) int {
	score := 0

	checkWindow := func(window []int, player int) int {
		score, playerCount, otherPlayerCount, emptyCount := 0, 0, 0, 0
		for i := 0; i < len(window); i++ {
			if window[i] == player {
				playerCount++
			} else if window[i] == BLANK {
				emptyCount++
			} else {
				otherPlayerCount++
			}
		}
		if playerCount == 4 {
			score = 100
		} else if playerCount == 3 && emptyCount == 1 {
			score = 5
		} else if playerCount == 2 && emptyCount == 2 {
			score = 2
		}
		if otherPlayerCount == 3 && emptyCount == 1 {
			score -= 4
		}
		return score
	}

	//score horizontal
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i])-3; j++ {
			window := make([]int, 4)
			for k := j; k < j+4; k++ {
				window[k-j] = board[i][k]
			}
			score += checkWindow(window, player)
		}
	}

	//score vertical
	for j := 0; j < len(board[0]); j++ {
		for i := 0; i < len(board)-3; i++ {
			window := make([]int, 4)
			for k := i; k < i+4; k++ {
				window[k-i] = board[k][j]
			}
			score += checkWindow(window, player)
		}
	}

	//score diagonals
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {

			//check positive sloped line
			window := make([]int, 0)
			for diagI, diagJ := i, j; diagI >= 0 && diagJ < len(board[0]); diagI, diagJ = diagI-1, diagJ+1 {
				window = append(window, board[diagI][diagJ])
			}
			if len(window) == 4 {
				score += checkWindow(window, player)
			}
			window = make([]int, 0)
			for diagI, diagJ := i, j; diagI < len(board) && diagJ < len(board[0]); diagI, diagJ = diagI+1, diagJ+1 {
				window = append(window, board[diagI][diagJ])
			}
			if len(window) == 4 {
				score += checkWindow(window, player)
			}
		}
	}

	return score
}

//minimax AI - returns (maxMove, maxValue)
func sMinimax(board [][]int, player, depth int) (int, int) {

	//terminal states
	if isWin(board, PLAYER2) {
		return 0, math.MaxInt32
	} else if isWin(board, PLAYER1) {
		return 0, math.MinInt32
	} else if boardFull(board) { //draw
		return 0, 0
	} else if depth <= 0 { //static evaluation function, depth limit reached
		return -1, stateHeuristic(board, player)
	}

	if player == PLAYER2 { //maximizing agent - AI
		bestMove := rand.Intn(len(board[0]))
		maxVal := math.MinInt32
		//simulate dropping piece into column i
		for i := 0; i < len(board[0]); i++ {
			newBoard := copyBoard(board)
			if !insertPiece(newBoard, i, player) {
				continue
			}
			_, val := sMinimax(newBoard, PLAYER1, depth-1)

			if val > maxVal {
				bestMove = i
				maxVal = val
			}
		}
		return bestMove, maxVal
	} else if player == PLAYER1 { //minimizing agent - human player
		bestMove := rand.Intn(len(board[0]))
		minVal := math.MaxInt32
		//simulate dropping piece into column i
		for i := 0; i < len(board[0]); i++ {
			newBoard := copyBoard(board)
			if !insertPiece(newBoard, i, player) {
				continue
			}
			_, val := sMinimax(newBoard, PLAYER2, depth-1)
			if val < minVal {
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
	otherPlayer := PLAYER2
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
			aiCol, _ := sMinimax(board, PLAYER2, MaxDepth)
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
