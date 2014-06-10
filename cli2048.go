package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const Version = "0.1"
const size = 4

//type SquareRows [4][4]int
//type SquareCols [4][4]int

var square [16]int
//var cols SquareCols
//var rows SquareRows
var rows [][]int = [][]int{
	[]int{1,2,3,4},
	[]int{5,6,7,8},
	[]int{9,10,11,12},
	[]int{13,14,15,16},
}
var cols [][]int = [][]int{
	[]int {0, 4, 8, 12},
	[]int {1, 5, 9, 13},
	[]int {2, 6, 10, 14},
	[]int {3, 7, 11, 15},
}

func main() {
	initCli2048()
	for i := 0; i < len(square); i++ {
		square[i] = 0
	}
	newCell()
	newCell()

	updateDisplay()
}

func initCli2048() {
	//
}

func newCell() {
	i := 0
	rand.Seed(time.Now().Unix())
	for {
		i = rand.Intn(len(square))
		if square[i] == 0 {
			break
		}
	}
	num := 2
	for {
		num = rand.Intn(5)
		if num % 2 == 0 && num != 0 {
			break
		}
	}
	//fmt.Printf("i=%d\nnum=%d\n", i, num)
	square[i] = num
}

func updateDisplay() {
	clearTerminal()
	printSep()
    
    for i := 0; i < size * size; i++ {
    	if (i+1) % size == 0 {
	    	fmt.Printf("|%s|\n", getCellString(square[i]))
	        printSep()
	    } else {
	    	fmt.Printf("|%s", getCellString(square[i]))
	    }
	}
}

func clearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printSep() {
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
