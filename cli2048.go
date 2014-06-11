package main

import (
	"./gmboard"
	"./gmdisplay"
	"./gmplayer"
	"fmt"
	"os"
)

const Version = "0.4.2"
const GameSize = 4
const Os = "Linux"

func main() {
	game := gmboard.NewGameBoard(GameSize)
	display := gmdisplay.NewGameDisplay(Os, GameSize)
	player := gmplayer.Player{}

	var ans []byte = make([]byte, 1)

	// Return console to normal upon exit
	defer display.CloseDisplay()

	for {

		display.UpdateDisplay(game.M, player.Score)
		newPoints := 0
		os.Stdin.Read(ans)
		var err error
		switch string(ans) {
		case "l":
			newPoints, err = game.ShiftRight()
		case "k":
			newPoints, err = game.ShiftDown()
		case "j":
			newPoints, err = game.ShiftLeft()
		case "i":
			newPoints, err = game.ShiftUp()
		case "q":
			return
		}

		if err != nil {
			continue
		} else {
			if err = game.NewCell(); err != nil {
				display.UpdateDisplay(game.M, player.Score)
				fmt.Printf("\n\n%s!\n\n", err.Error())
				return
			}
		}

		player.Score += newPoints
	}
}
