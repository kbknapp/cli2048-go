package gmboard

import (
	"errors"
	"fmt"
	"github.com/kbknapp/go/matrix"
	"math/rand"
	"time"
)

const GameSize = 4

type GameBoard struct {
	matrix.Matrix
}

func NewGameBoard() GameBoard {

	gb := GameBoard{matrix.NewMatrix(GameSize)}
	for i := 0; i < len(gb.M); i++ {
		gb.M[i] = 0
	}
	fmt.Printf("%v\n", gb)
	return gb
}

func (gb *GameBoard) NewCell() {
	i := 0
	rand.Seed(time.Now().Unix())
	for {
		i = rand.Intn(len(gb.M))
		if gb.M[i] == 0 {
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
	fmt.Printf("i=%d\nnum=%d\n", i, num)
	gb.M[i] = num
}

func (gb *GameBoard) ShiftUp() error {
	done := false
	moves := 0
	var offLimits [GameSize * GameSize]int
	for {
		for i, row := range gb.Rows {
			for j, cell := range row {
				if i == 0 {
					continue
				}

				currNum := gb.M[cell].(int)

				if currNum == 0 {
					continue
				}
				newIndex := -1
				posIndex := -1
				for ni := i - 1; ni >= 0; ni-- {
					newIndex = gb.Rows[ni][j]
					if gb.M[newIndex] == 0 {
						posIndex = newIndex
						continue
					} else {
						break
					}
				}
				//fmt.Printf("i %d, j %d, cN %d, n_i %d, pI %d, c %d\n", i, j, currNum, newIndex, posIndex, cell)
				//fmt.Printf("%v\n", gb.M)
				if gb.M[newIndex] == 0 {
					gb.M[newIndex] = currNum
					gb.M[cell] = 0
					done = false
					moves++
				} else if currNum == gb.M[newIndex] {
					if offLimits[cell] != 1 {
						gb.M[newIndex] = currNum * 2
						offLimits[newIndex] = 1
						gb.M[cell] = 0
						done = false
						moves++
					}

				} else {
					if posIndex == -1 {
						continue
					}
					gb.M[posIndex] = currNum
					gb.M[cell] = 0
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

func (gb *GameBoard) ShiftDown() error {
	done := false
	moves := 0
	var offLimits [GameSize * GameSize]int
	for {
		for i := gb.Size - 1; i >= 0; i-- {
			for j := gb.Size - 1; j >= 0; j-- {
				if i == gb.Size-1 {
					continue
				}

				currNum := gb.M[gb.Rows[i][j]].(int)

				if currNum == 0 {
					continue
				}
				newIndex := -1
				posIndex := -1
				for ni := i + 1; ni <= gb.Size-1; ni++ {
					newIndex = gb.Rows[ni][j]
					if gb.M[newIndex] == 0 {
						posIndex = newIndex
						continue
					} else {
						break
					}
				}
				//fmt.Printf("i %d, j %d, cN %d, n_i %d, pI %d, c %d\n", i, j, currNum, newIndex, posIndex, gb.M[rows[i][j]])
				//fmt.Printf("%v\n", gb.M)
				if gb.M[newIndex] == 0 {
					gb.M[newIndex] = currNum
					gb.M[gb.Rows[i][j]] = 0
					done = false
					moves++
				} else if currNum == gb.M[newIndex] {
					if offLimits[gb.Rows[i][j]] != 1 {
						gb.M[newIndex] = currNum * 2
						offLimits[newIndex] = 1
						gb.M[gb.Rows[i][j]] = 0
						done = false
						moves++
					}

				} else {
					if posIndex == -1 {
						continue
					}
					gb.M[posIndex] = currNum
					gb.M[gb.Rows[i][j]] = 0
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

func (gb *GameBoard) ShiftLeft() error {
	done := false
	moves := 0
	var offLimits [16]int
	for {
		for i, col := range gb.Cols {
			for j, cell := range col {
				if i == 0 {
					continue
				}

				currNum := gb.M[cell].(int)

				if currNum == 0 {
					continue
				}
				newIndex := -1
				posIndex := -1
				for ni := i - 1; ni >= 0; ni-- {
					newIndex = gb.Cols[ni][j]
					if gb.M[newIndex] == 0 {
						posIndex = newIndex
						continue
					} else {
						break
					}
				}
				//fmt.Printf("i %d, j %d, cN %d, n_i %d, pI %d, c %d\n", i, j, currNum, newIndex, posIndex, cell)
				//fmt.Printf("%v\n", gb.M)
				if gb.M[newIndex] == 0 {
					gb.M[newIndex] = currNum
					gb.M[cell] = 0
					done = false
					moves++
				} else if currNum == gb.M[newIndex] {
					if offLimits[cell] != 1 {
						gb.M[newIndex] = currNum * 2
						offLimits[newIndex] = 1
						gb.M[cell] = 0
						done = false
						moves++
					}

				} else {
					if posIndex == -1 {
						continue
					}
					gb.M[posIndex] = currNum
					gb.M[cell] = 0
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

func (gb *GameBoard) ShiftRight() error {
	done := false
	moves := 0
	var offLimits [GameSize * GameSize]int
	for {
		for i := gb.Size - 1; i >= 0; i-- {
			for j := gb.Size - 1; j >= 0; j-- {
				if i == gb.Size-1 {
					continue
				}

				currNum := gb.M[gb.Cols[i][j]].(int)

				if currNum == 0 {
					continue
				}
				newIndex := -1
				posIndex := -1
				for ni := i + 1; ni <= gb.Size-1; ni++ {
					newIndex = gb.Cols[ni][j]
					if gb.M[newIndex] == 0 {
						posIndex = newIndex
						continue
					} else {
						break
					}
				}
				//fmt.Printf("i %d, j %d, cN %d, n_i %d, pI %d, c %d\n", i, j, currNum, newIndex, posIndex, gb.M[rows[i][j]])
				//fmt.Printf("%v\n", gb.M)
				if gb.M[newIndex] == 0 {
					gb.M[newIndex] = currNum
					gb.M[gb.Cols[i][j]] = 0
					done = false
					moves++
				} else if currNum == gb.M[newIndex] {
					if offLimits[gb.Cols[i][j]] != 1 {
						gb.M[newIndex] = currNum * 2
						offLimits[newIndex] = 1
						gb.M[gb.Cols[i][j]] = 0
						done = false
						moves++
					}

				} else {
					if posIndex == -1 {
						continue
					}
					gb.M[posIndex] = currNum
					gb.M[gb.Cols[i][j]] = 0
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
