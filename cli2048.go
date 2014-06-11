package main

import (
	"./game"
	"fmt"
	"os"
	"strings"
)

const Version = "0.4.7"
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
		shifted := 0
		os.Stdin.Read(ans)
		var err error
		switch strings.ToLower(string(ans)) {
		case "l":
			fallthrough
		case "d":
			newPoints, shifted = board.ShiftRight()
		case "k":
			fallthrough
		case "s":
			newPoints, shifted = board.ShiftDown()
		case "j":
			fallthrough
		case "a":
			newPoints, shifted = board.ShiftLeft()
		case "i":
			fallthrough
		case "w":
			newPoints, shifted = board.ShiftUp()
		case "n":
			fmt.Printf("Are you sure you want to reset the game?[n]: ")
			os.Stdin.Read(ans)
			if strings.ToLower(string(ans)) == "y" {
				board.Reset()
				player.Score = 0
			}
		case "q":
			fmt.Printf("Are you sure you want to quit?[n]: ")
			os.Stdin.Read(ans)
			if strings.ToLower(string(ans)) == "y" {
				fmt.Println("")
				return
			}
		}

		if shifted == 0 {
			continue
		} else {
			if err = board.NewCell(); err != nil {
				display.UpdateDisplay(board.M, player.Score)
				fmt.Printf("\n\n%s!\n\n", err.Error())
				fmt.Printf("Play again?[y]: ")
				os.Stdin.Read(ans)
				if strings.ToLower(string(ans)) == "n" {
					fmt.Println("")
					return
				} else {
					board.Reset()
					player.Score = 0
				}
			}
		}

		player.Score += newPoints
	}
}
