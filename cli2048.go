package main

import (
	"./gmboard"
	"fmt"
	"os"
	"os/exec"
)

const Version = "0.3"

func main() {
	game := gmboard.NewGameBoard()

	game.NewCell()
	game.NewCell()

	var ans []byte = make([]byte, 1)

	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not disp enter chars on screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	// Return console to normal upon exit
	defer exec.Command("stty", "-F", "/dev/tty", "-cbreak").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	for {

		updateDisplay(game)

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
			if err = game.NewCell(); err != nil {
				fmt.Printf("\n\n%s!\n\n", err.Error())
				return
			}
		} else {
			game.NewCell()
		}
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
