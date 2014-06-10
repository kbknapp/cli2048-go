package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const Version = "0.1"
const size = 4

var square [16]int

var rows [][]int = [][]int{
	[]int{0, 1, 2, 3},
	[]int{4, 5, 6, 7},
	[]int{8, 9, 10, 11},
	[]int{12, 13, 14, 15},
}
var cols [][]int = [][]int{
	[]int{0, 4, 8, 12},
	[]int{1, 5, 9, 13},
	[]int{2, 6, 10, 14},
	[]int{3, 7, 11, 15},
}

func main() {
	initCli2048()
	for i := 0; i < len(square); i++ {
		square[i] = 0
	}
	newCell()
	newCell()

	ans := ""
	for {

		updateDisplay()

		fmt.Print("Move: ")
		fmt.Scanf("%v", &ans)

		switch ans {
		case "l":
			if err := moveRight(); err != nil {
				continue
			}
		case "k":
			if err := moveDown(); err != nil {
				continue
			}
		case "j":
			moveLeft()
		case "i":
			if err := moveUp(); err != nil {
				continue
			}
		case "q":
			break
		}

		newCell()
	}
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
		if num%2 == 0 && num != 0 {
			break
		}
	}
	//fmt.Printf("i=%d\nnum=%d\n", i, num)
	square[i] = num
}

func updateDisplay() {
	clearTerminal()
	printSep()

	for i := 0; i < size*size; i++ {
		if (i+1)%size == 0 {
			fmt.Printf("|%s|\n", getCellString(square[i]))
			printSep()
		} else {
			fmt.Printf("|%s", getCellString(square[i]))
		}
	}
}

func clearTerminal() {
	cmd := exec.Command("clear")
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

func moveUp() error {
	done := false
	moves := 0
	var offLimits [16]int
	for {
		for i, row := range rows {
			for j, cell := range row {
				if i == 0 {
					continue
				}

				currNum := square[cell]

				if currNum == 0 {
					continue
				}
				new_index := -1
				posIndex := -1
				for ni := i - 1; ni >= 0; ni-- {
					new_index = rows[ni][j]
					if square[new_index] == 0 {
						posIndex = new_index
						continue
					} else {
						break
					}
				}
				//fmt.Printf("i %d, j %d, cN %d, n_i %d, pI %d, c %d\n", i, j, currNum, new_index, posIndex, cell)
				//fmt.Printf("%v\n", square)
				if square[new_index] == 0 {
					square[new_index] = currNum
					square[cell] = 0
					done = false
					moves++
				} else if currNum == square[new_index] {
					if offLimits[cell] != 1 {
						square[new_index] = currNum * 2
						offLimits[new_index] = 1
						square[cell] = 0
						done = false
						moves++
					}

				} else {
					if posIndex == -1 {
						continue
					}
					square[posIndex] = currNum
					square[cell] = 0
					done = false
					moves++
				}
			}
		}
		if done {
			break
		}
		done = true
	}
	if moves > 0 {
		return nil
	}
	return errors.New("No moves")
}

func moveDown() error {
	done := false
	moves := 0
	var offLimits [16]int
	for {
		for i := size - 1; i >= 0; i-- {
			for j := size - 1; j >= 0; j-- {
				if i == size-1 {
					continue
				}

				currNum := square[rows[i][j]]

				if currNum == 0 {
					continue
				}
				new_index := -1
				posIndex := -1
				for ni := i + 1; ni <= size-1; ni++ {
					new_index = rows[ni][j]
					if square[new_index] == 0 {
						posIndex = new_index
						continue
					} else {
						break
					}
				}
				//fmt.Printf("i %d, j %d, cN %d, n_i %d, pI %d, c %d\n", i, j, currNum, new_index, posIndex, square[rows[i][j]])
				//fmt.Printf("%v\n", square)
				if square[new_index] == 0 {
					square[new_index] = currNum
					square[rows[i][j]] = 0
					done = false
					moves++
				} else if currNum == square[new_index] {
					if offLimits[rows[i][j]] != 1 {
						square[new_index] = currNum * 2
						offLimits[new_index] = 1
						square[rows[i][j]] = 0
						done = false
						moves++
					}

				} else {
					if posIndex == -1 {
						continue
					}
					square[posIndex] = currNum
					square[rows[i][j]] = 0
					done = false
					moves++
				}
			}
		}
		if done {
			break
		}
		done = true
	}
	if moves > 0 {
		return nil
	}
	return errors.New("No moves")
}

func moveLeft() error {
	done := false
	moves := 0
	var offLimits [16]int
	for {
		for i, col := range cols {
			for j, cell := range col {
				if i == 0 {
					continue
				}

				currNum := square[cell]

				if currNum == 0 {
					continue
				}
				new_index := -1
				posIndex := -1
				for ni := i - 1; ni >= 0; ni-- {
					new_index = cols[ni][j]
					if square[new_index] == 0 {
						posIndex = new_index
						continue
					} else {
						break
					}
				}
				//fmt.Printf("i %d, j %d, cN %d, n_i %d, pI %d, c %d\n", i, j, currNum, new_index, posIndex, cell)
				//fmt.Printf("%v\n", square)
				if square[new_index] == 0 {
					square[new_index] = currNum
					square[cell] = 0
					done = false
					moves++
				} else if currNum == square[new_index] {
					if offLimits[cell] != 1 {
						square[new_index] = currNum * 2
						offLimits[new_index] = 1
						square[cell] = 0
						done = false
						moves++
					}

				} else {
					if posIndex == -1 {
						continue
					}
					square[posIndex] = currNum
					square[cell] = 0
					done = false
					moves++
				}
			}
		}
		if done {
			break
		}
		done = true
	}
	if moves > 0 {
		return nil
	}
	return errors.New("No moves")
}

func moveRight() error {
	done := false
	moves := 0
	var offLimits [16]int
	for {
		for i := size - 1; i >= 0; i-- {
			for j := size - 1; j >= 0; j-- {
				if i == size-1 {
					continue
				}

				currNum := square[cols[i][j]]

				if currNum == 0 {
					continue
				}
				new_index := -1
				posIndex := -1
				for ni := i + 1; ni <= size-1; ni++ {
					new_index = cols[ni][j]
					if square[new_index] == 0 {
						posIndex = new_index
						continue
					} else {
						break
					}
				}
				//fmt.Printf("i %d, j %d, cN %d, n_i %d, pI %d, c %d\n", i, j, currNum, new_index, posIndex, square[rows[i][j]])
				//fmt.Printf("%v\n", square)
				if square[new_index] == 0 {
					square[new_index] = currNum
					square[cols[i][j]] = 0
					done = false
					moves++
				} else if currNum == square[new_index] {
					if offLimits[cols[i][j]] != 1 {
						square[new_index] = currNum * 2
						offLimits[new_index] = 1
						square[cols[i][j]] = 0
						done = false
						moves++
					}

				} else {
					if posIndex == -1 {
						continue
					}
					square[posIndex] = currNum
					square[cols[i][j]] = 0
					done = false
					moves++
				}
			}
		}
		if done {
			break
		}
		done = true
	}
	if moves > 0 {
		return nil
	}
	return errors.New("No moves")
}
