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
	M []int
}

func NewGameBoard() GameBoard {

	gb := GameBoard{matrix.NewMatrix(GameSize), make([]int, GameSize*GameSize)}
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
			if num == 4 {
				w := rand.Intn(4)
				if (w % 2) != 0 {
					break
				}
			} else {
				break
			}
		}
	}

	gb.M[i] = num
}

func shiftIndices(gb *GameBoard, indices [][]int) error {
	done := false
	moves := 0
	offLimits := make([]bool, len(gb.M))
	for {
		for i, seq := range indices {
			for j, cell := range seq {
				if i == 0 {
					continue
				}

				currNum := gb.M[cell]

				if currNum == 0 {
					continue
				}

				newIndex := -1
				posIndex := -1
				for ni := i - 1; ni >= 0; ni-- {
					newIndex = indices[ni][j]
					if gb.M[newIndex] == 0 {
						posIndex = newIndex
						continue
					} else {
						break
					}
				}

				if gb.M[newIndex] == 0 {
					gb.M[newIndex] = currNum
					gb.M[cell] = 0
					done = false
					moves++
				} else if currNum == gb.M[newIndex] {
					if !offLimits[cell] {
						gb.M[newIndex] = currNum * 2
						offLimits[newIndex] = true
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

func shiftIndicesRev(gb *GameBoard, indices [][]int) error {
	done := false
	moves := 0
	offLimits := make([]bool, len(gb.M))
	for {
		for i := gb.Size - 1; i >= 0; i-- {
			for j := gb.Size - 1; j >= 0; j-- {
				if i == gb.Size-1 {
					continue
				}

				currNum := gb.M[indices[i][j]]

				if currNum == 0 {
					continue
				}

				newIndex := -1
				posIndex := -1
				for ni := i + 1; ni <= gb.Size-1; ni++ {
					newIndex = indices[ni][j]
					if gb.M[newIndex] == 0 {
						posIndex = newIndex
						continue
					} else {
						break
					}
				}

				if gb.M[newIndex] == 0 {
					gb.M[newIndex] = currNum
					gb.M[indices[i][j]] = 0
					done = false
					moves++
				} else if currNum == gb.M[newIndex] {
					if !offLimits[indices[i][j]] {
						gb.M[newIndex] = currNum * 2
						offLimits[newIndex] = true
						gb.M[indices[i][j]] = 0
						done = false
						moves++
					}

				} else {
					if posIndex == -1 {
						continue
					}
					gb.M[posIndex] = currNum
					gb.M[indices[i][j]] = 0
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

func (gb *GameBoard) ShiftUp() error {
	return shiftIndices(gb, gb.Rows)
}
func (gb *GameBoard) ShiftDown() error {
	return shiftIndicesRev(gb, gb.Rows)
}

func (gb *GameBoard) ShiftLeft() error {
	return shiftIndices(gb, gb.Cols)
}

func (gb *GameBoard) ShiftRight() error {
	return shiftIndicesRev(gb, gb.Cols)
}
