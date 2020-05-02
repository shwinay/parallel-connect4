package main

import (
	"fmt"
	"math"
	"math/rand"
)

func pAlphaBeta(board [][]int, player, depth int) (int, int) {
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

		resultChan := make(chan Result)
		numBranches := 0
		//simulate dropping piece into column i
		for i := 0; i < len(board[0]); i++ {

			newBoard := copyBoard(board)
			if !insertPiece(newBoard, i, player) {
				continue
			}
			go pBranchAlphaBeta(newBoard, PLAYER1, depth-1, math.MinInt32, math.MaxInt32, resultChan, i)
			numBranches++
		}

		for i := 0; i < numBranches; i++ {
			res := <-resultChan
			if res.val > maxVal {
				bestMove = res.col
				maxVal = res.val
			}
		}

		return bestMove, maxVal
	} else {
		fmt.Println("not minimizing or maximizing agent.. this should never happen")
		return -1, -1
	}
}

func pBranchAlphaBeta(board [][]int, player, depth, alpha, beta int, resultChan chan Result, col int) {

	//terminal states
	if isWin(board, PLAYER2) {
		resultChan <- Result{0, math.MaxInt32, col}
	} else if isWin(board, PLAYER1) {
		resultChan <- Result{0, math.MinInt32, col}
	} else if boardFull(board) { //draw
		resultChan <- Result{0, 0, col}
	} else if depth <= 0 { //static evaluation function, depth limit reached
		resultChan <- Result{-1, stateHeuristic(board, player), col}
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
			_, val := sAlphaBeta(newBoard, PLAYER1, depth-1, alpha, beta)

			if val > maxVal {
				bestMove = i
				maxVal = val
			}
		}
		resultChan <- Result{bestMove, maxVal, col}
	} else if player == PLAYER1 { //minimizing agent - human player
		bestMove := rand.Intn(len(board[0]))
		minVal := math.MaxInt32
		//simulate dropping piece into column i
		for i := 0; i < len(board[0]); i++ {
			newBoard := copyBoard(board)
			if !insertPiece(newBoard, i, player) {
				continue
			}
			_, val := sAlphaBeta(newBoard, PLAYER2, depth-1, alpha, beta)
			if val < minVal {
				bestMove = i
				minVal = val
			}
		}
		resultChan <- Result{bestMove, minVal, col}
	} else {
		fmt.Println("not minimizing or maximizing agent.. this should never happen")
		resultChan <- Result{-1, -1, -1}
	}
}
