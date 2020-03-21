package main

func main() {
	board := initBoard(8, 8)
	winner := playAI(board)
	endGame(board, winner, false)
}
