package game

import (
	"errors"
	"github.com/kbknapp/gok/math"
	"math/rand"
	"time"
)

type GameBoard struct {
	math.Matrix
}

func NewGameBoard(size int) GameBoard {

	gb := GameBoard{math.NewMatrix(size), make([]int, size*size)}
	for i := 0; i < len(gb.M); i++ {
		gb.M[i] = 0
	}

	gb.NewCell()
	gb.NewCell()

	return gb
}

func (gb *GameBoard) Reset() {
	for i, _ := range gb.M {
		gb.M[i] = 0
	}
	gb.NewCell()
	gb.NewCell()
}

func (gb *GameBoard) NewCell() error {
	if gb.isFull() {
		return gb.movesLeft()
	}
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

	return gb.movesLeft()
}

func (gb *GameBoard) shiftIndices(indices [][]int) (int, int) {
	done := false
	moves := 0
	p := 0
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
						p += gb.M[newIndex]
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
	return p, moves
}

func (gb *GameBoard) shiftIndicesRev(indices [][]int) (int, int) {
	done := false
	moves := 0
	p := 0
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
						p += gb.M[newIndex]
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
	return p, moves
}

func (gb *GameBoard) ShiftUp() (int, int) {
	return gb.shiftIndices(gb.Rows)
}
func (gb *GameBoard) ShiftDown() (int, int) {
	return gb.shiftIndicesRev(gb.Rows)
}

func (gb *GameBoard) ShiftLeft() (int, int) {
	return gb.shiftIndices(gb.Cols)
}

func (gb *GameBoard) ShiftRight() (int, int) {
	return gb.shiftIndicesRev(gb.Cols)
}

func (gb *GameBoard) movesLeft() error {
	a := 0
	b := gb.M[1]
	for i := 0; i < len(gb.M)-1; i++ {
		if (i+1)%gb.Size == 0 {
			b = gb.M[i+2]
			continue
		}
		a = gb.M[i]
		if a == b {
			return nil
		}
		if i >= len(gb.M)-2 {
			break
		}
		b = gb.M[i+2]
	}
	a = 0
	b = gb.M[gb.Size]
	for i, col := range gb.Cols {
		for j, cell := range col {
			a = gb.M[cell]
			if a == b {
				return nil
			}
			if j >= 2 {
				if i < gb.Size-1 {
					b = gb.M[gb.Cols[i+1][1]]
				}
				break
			}
			b = gb.M[col[j+2]]
		}
	}
	if gb.isFull() {
		return errors.New("Game over")
	} else {
		return nil
	}
}

func (gb *GameBoard) isFull() bool {
	for i := 0; i < len(gb.M); i++ {
		if gb.M[i] == 0 {
			return false
		}
	}
	return true
}
