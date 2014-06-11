package main

import (
	"./game"
	"fmt"
	"os"
)

const Version = "0.4.4"
const boardSize = 4
const Os = "Linux"

func main() {
	board := game.NewGameBoard(boardSize)
	display := game.NewGameDisplay(Os, boardSize)
	player := game.Player{}

	var ans []byte = make([]byte, 1)

	// Return console to normal upon exit
	defer display.CloseDisplay()

	for {

		display.UpdateDisplay(board.M, player.Score)
		newPoints := 0
		os.Stdin.Read(ans)
		var err error
		switch string(ans) {
		case "l":
			newPoints, err = board.ShiftRight()
		case "k":
			newPoints, err = board.ShiftDown()
		case "j":
			newPoints, err = board.ShiftLeft()
		case "i":
			newPoints, err = board.ShiftUp()
		case "q":
			return
		}

		if err != nil {
			continue
		} else {
			if err = board.NewCell(); err != nil {
				display.UpdateDisplay(board.M, player.Score)
				fmt.Printf("\n\n%s!\n\n", err.Error())
				return
			}
		}

		player.Score += newPoints
	}
}
