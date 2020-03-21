package main

func main() {
	board := initBoard(7, 6)
	winner := playAI(board)
	endGame(board, winner, false)
}
