package gmboard

import (
	"errors"
	"math/rand"
)

const size = 4

type GameBoard struct {
	[size]int
	rows [size][size]int 
	cols [size][size]int
}

func NewGameBoard() *GameBoard {
	gb = GameBoard{}
	r := 0
	c := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			gb.rows[i][j] = r
			r++
			gb.cols[i][j] = c
			c += size
		}
		c = size + i
	}

	return &gb
}

func (gb *GameBoard) newCell() {
	i := 0
	rand.Seed(time.Now().Unix())
	for {
		i = rand.Intn(len(gb))
		if gb[i] == 0 {
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
	gb[i] = num
}

func (gb *GameBoard) ShiftUp() error {
	done := false
	moves := 0
	var offLimits [16]int
	for {
		for i, row := range gb.rows {
			for j, cell := range row {
				if i == 0 {
					continue
				}

				currNum := gb[cell]

				if currNum == 0 {
					continue
				}
				new_index := -1
				posIndex := -1
				for ni := i - 1; ni >= 0; ni-- {
					new_index = gb.rows[ni][j]
					if square[new_index] == 0 {
						posIndex = new_index
						continue
					} else {
						break
					}
				}
				//fmt.Printf("i %d, j %d, cN %d, n_i %d, pI %d, c %d\n", i, j, currNum, new_index, posIndex, cell)
				//fmt.Printf("%v\n", square)
				if gb[new_index] == 0 {
					gb[new_index] = currNum
					gb[cell] = 0
					done = false
					moves++
				} else if currNum == gb[new_index] {
					if offLimits[cell] != 1 {
						gb[new_index] = currNum * 2
						offLimits[new_index] = 1
						gb[cell] = 0
						done = false
						moves++
					}

				} else {
					if posIndex == -1 {
						continue
					}
					gb[posIndex] = currNum
					gb[cell] = 0
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
