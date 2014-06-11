package main

import (
	"./game"
	"fmt"
	"os"
	"strings"
)

const Version = "0.5.1"
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

		display.UpdateDisplay(board.M, player.Score, player.HighScore)
		newPoints := 0
		shifted := 0
		os.Stdin.Read(ans)
		var err error
		switch strings.ToLower(string(ans)) {
		case "l":
			// RIGHT
			fallthrough
		case "c":
			// [RIGHT ARROW]
			fallthrough
		case "d":
			// Check  ans byte in case left arrow (0x44) was pressed
			if ans[0] == 68 {
				newPoints, shifted = board.ShiftLeft()
			} else {
				newPoints, shifted = board.ShiftRight()
			}
		case "k":
			// DOWN
			fallthrough
		case "b":
			// [DOWN ARROW]
			fallthrough
		case "s":
			newPoints, shifted = board.ShiftDown()
		case "j":
			// LEFT
			fallthrough
		case "a":
			// Check ans byte in case up arrow (0x41) was pressed
			if ans[0] == 65 {
				newPoints, shifted = board.ShiftUp()
			} else {
				newPoints, shifted = board.ShiftLeft()
			}
		case "i":
			// UP
			fallthrough
		case "w":
			newPoints, shifted = board.ShiftUp()
		case "n":
			// NEW
			fallthrough
		case "r":
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
		case "h":
			fmt.Println("")
			fmt.Print(`
Controls:

UP	w, i, [UP ARROW]
DOWN	s, k, [DOWN ARROW]
LEFT	a*, j, [LEFT ARROW]
RIGHT	d*, l, [RIGHT ARROW]

HELP	h

QUIT	q, [ESC]

RESET		n, r
NEW GAME	n, r
			 
* If experiencing unexpected movements ensure
  your [CAPS LOCK] is not on)`)
			fmt.Println("")
			os.Stdin.Read(ans)
		default:
			if ans[0] == 27 {
				// ESC
				fmt.Printf("Are you sure you want to quit?[n]: ")
				os.Stdin.Read(ans)
				if strings.ToLower(string(ans)) == "y" {
					fmt.Println("")
					return
				}
			}
			//fmt.Printf("You typed string: %s\n Byte: %x", string(ans), ans)
			//os.Stdin.Read(ans)
		}

		if shifted == 0 {
			continue
		} else {
			if err = board.NewCell(); err != nil {
				display.UpdateDisplay(board.M, player.Score, player.HighScore)
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
		if player.Score >= player.HighScore {
			player.HighScore = player.Score
		}
	}
}
