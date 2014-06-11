package game

import (
	"fmt"
	"os"
	"os/exec"
)

type GameDisplay struct {
	Os   string
	Size int
}

func NewGameDisplay(os string, size int) GameDisplay {

	switch os {
	case "Linux":
		// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not disp enter chars on screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	case "Windows":
		//
	}
	return GameDisplay{os, size}
}

func (gd *GameDisplay) CloseDisplay() {
	switch gd.Os {
	case "Linux":
		exec.Command("stty", "-F", "/dev/tty", "-cbreak").Run()
		exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	case "Windows":
		//
	}

}

func (gd *GameDisplay) UpdateDisplay(seq []int, score int) {
	gd.clearTerminal()

	fmt.Printf("Score:\t%d\n", score)
	gd.printSep()

	for i := 0; i < len(seq); i++ {
		if (i+1)%gd.Size == 0 {
			fmt.Printf("|%s|\n", gd.getCellString(seq[i]))
			gd.printSep()
		} else {
			fmt.Printf("|%s", gd.getCellString(seq[i]))
		}
	}
}

func (gd *GameDisplay) clearTerminal() {
	switch gd.Os {
	case "Linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "Windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (gd *GameDisplay) printSep() {
	for i := 0; i < gd.Size; i++ {
		fmt.Print("+----")
	}
	fmt.Println("+")
}

func (gd *GameDisplay) getCellString(v int) string {
	switch gd.Os {
	case "Linux":
		return gd.formatCellNix(v)
	case "Windows":
		fallthrough
	default:
		return gd.formatCellWin(v)
	}

}

func (gd *GameDisplay) formatCellNix(v int) string {
	s := ""
	switch v {
	case 2:
		s = fmt.Sprintf("\x1b[0;31m  %d \x1b[0m", v)
	case 4:
		s = fmt.Sprintf("\x1b[1;31m  %d \x1b[0m", v)
	case 8:
		s = fmt.Sprintf("\x1b[0;33m  %d \x1b[0m", v)
	case 16:
		fallthrough
	case 32:
		s = fmt.Sprintf("\x1b[1;33m %d \x1b[0m", v)
	case 64:
		s = fmt.Sprintf("\x1b[0;34m %d \x1b[0m", v)
	case 128:
		fallthrough
	case 256:
		s = fmt.Sprintf("\x1b[1;34m %d\x1b[0m", v)
	case 512:
		s = fmt.Sprintf("\x1b[0;32m %d\x1b[0m", v)
	case 1024:
		fallthrough
	case 2048:
		s = fmt.Sprintf("\x1b[1;32m%d\x1b[0m", v)
	default:
		s = "    "
	}
	return s
}

func (gd *GameDisplay) formatCellWin(v int) string {
	s := ""
	switch v {
	case 2:
		fallthrough
	case 4:
		fallthrough
	case 8:
		s = fmt.Sprintf("  %d ", v)
	case 16:
		fallthrough
	case 32:
		fallthrough
	case 64:
		s = fmt.Sprintf(" %d ", v)
	case 128:
		fallthrough
	case 256:
		fallthrough
	case 512:
		s = fmt.Sprintf(" %d", v)
	case 1024:
		fallthrough
	case 2048:
		s = fmt.Sprintf("%d", v)
	default:
		s = "    "
	}
	return s
}
