package main

import (
	"./gmboard"
	"fmt"
	"os"
	"os/exec"
)

const Version = "0.3"

func main() {
	//fmt.Println("Making Game board...")
	game := gmboard.NewGameBoard()

	//fmt.Println("Making cells...")
	game.NewCell()
	game.NewCell()

	var ans []byte = make([]byte, 1)

	//fmt.Println("Setting ttys...")
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not disp enter chars on screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	for {

		updateDisplay(game)

		//fmt.Printf("cols=%v\nrows=%v\n\n", game.Cols, game.Rows)

		os.Stdin.Read(ans)

		switch string(ans) {
		case "l":
			if err := game.ShiftRight(); err != nil {
				continue
			}
		case "k":
			if err := game.ShiftDown(); err != nil {
				continue
			}
		case "j":
			if err := game.ShiftLeft(); err != nil {
				continue
			}
		case "i":
			if err := game.ShiftUp(); err != nil {
				continue
			}
		case "q":
			return
		}

		game.NewCell()
	}
}

func updateDisplay(gb gmboard.GameBoard) {
	clearTerminal()
	printSep(gb.Size)

	for i := 0; i < len(gb.M); i++ {
		if (i+1)%gb.Size == 0 {
			fmt.Printf("|%s|\n", getCellString(gb.M[i]))
			printSep(gb.Size)
		} else {
			fmt.Printf("|%s", getCellString(gb.M[i]))
		}
	}
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printSep(size int) {
	for i := 0; i < size; i++ {
		fmt.Print("+----")
	}
	fmt.Println("+")
}

func getCellString(v int) string {
	if v == 0 {
		return "    "
	} else if v < 10 {
		return fmt.Sprintf("  %d ", v)
	} else if v < 100 {
		return fmt.Sprintf(" %d ", v)
	} else if v < 1000 {
		return fmt.Sprintf(" %d", v)
	} else {
		return fmt.Sprintf("%d", v)
	}
}
