package main

func main() {
	board := initBoard(7, 6)
	winner := playTwoPlayer(board)
	endGame(board, winner)
}
