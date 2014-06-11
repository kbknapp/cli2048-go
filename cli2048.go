package main

import (
	"./gmboard"
	"./gmdisplay"
	"fmt"
	"os"
)

const Version = "0.4.2"
const GameSize = 4
const Os = "Linux"

func main() {
	game := gmboard.NewGameBoard(GameSize)
	display := gmdisplay.NewGameDisplay(Os, GameSize)

	var ans []byte = make([]byte, 1)

	// Return console to normal upon exit
	defer display.CloseDisplay()

	for {

		display.UpdateDisplay(game.M)

		os.Stdin.Read(ans)
		var err error
		switch string(ans) {
		case "l":
			err = game.ShiftRight()
		case "k":
			err = game.ShiftDown()
		case "j":
			err = game.ShiftLeft()
		case "i":
			err = game.ShiftUp()
		case "q":
			return
		}

		if err != nil {
			continue
		} else {
			if err = game.NewCell(); err != nil {
				display.UpdateDisplay(game.M)
				fmt.Printf("\n\n%s!\n\n", err.Error())
				return
			}
		}
	}
}
