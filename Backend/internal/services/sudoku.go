// Package services
// This is a sudoku interface for the CSP
package services

import (
	"errors"
	"strconv"
)

// Public

type Board struct {
	numbers [][]int
}

type SudokuCtx struct {
	board       Board
	variables   []variable
	domains     map[variable][]int
	constraints int
}

// SetNumbers Setter for Board
func (board *Board) SetNumbers(numbers [][]int) {
	board.numbers = numbers
}

// Convert to String function for the Board struct to represent a Sudoku Board.
func (board *Board) String() (boardString string) {
	i := 0
	for i < len(board.numbers) {
		j := 0
		if i%3 == 0 && i != 0 {
			boardString += "- - - - - - - - - - -\n"
		}
		for j < len(board.numbers[i]) {
			if j%3 == 0 && j != 0 {
				boardString += "| "
			}
			boardString += strconv.Itoa(board.numbers[i][j]) + " "
			j++
		}

		boardString += "\n"
		i++
	}
	return
}

func InitSudokuCtx(board Board) (sudokuCtx SudokuCtx, err error) {
	sudokuCtx.board = board

	if len(board.numbers) <= 0 {
		return sudokuCtx, errors.New("board is empty")
	}

	sudokuCtx.initSudokuVariables()
	sudokuCtx.initSudokuDomains()

	return
}

// Private

type variable struct {
	a, b interface{}
}

// Initializes Variables for a Sudoku Board and sets it in the SudokuCtx
// Basically a Setter...
func (sudokuCtx *SudokuCtx) initSudokuVariables() {
	variables := make([]variable, 0)
	maxCells := 9
	for i := 0; i < maxCells; i++ {
		for j := 0; j < maxCells; j++ {
			variables = append(variables, variable{i, j})
		}
	}

	sudokuCtx.variables = variables
}

// Initializes Domains for Sudoku Board and sets it in the SudokuCtx
// Depends on initialization of variable beforehand
func (sudokuCtx *SudokuCtx) initSudokuDomains() {
	domains := make(map[variable][]int)

	for _, position := range sudokuCtx.variables {
		domains[position] = make([]int, 0)
	}

	for i, row := range sudokuCtx.board.numbers {
		for j, cell := range row {
			position := variable{i, j}

			if cell != 0 {
				domains[position] = append(domains[position], cell)
				continue
			}

			for k := 1; k < 10; k++ {
				domains[position] = append(domains[position], k)
			}

		}
	}

	sudokuCtx.domains = domains
}

// Initializes the Constraints for Sudoku Board and sets it in the SudokuCtx
func (sudokuCtx *SudokuCtx) initSudokuConstraints(variables []variable) {
	constraints := make(map[variable][]variable)

	for i := range variables {
		addConstraints(i, constraints)
	}

}

func addConstraints(v variable, constraints map[variable][]variable) {
	vA := v.a.(int)
	vB := v.b.(int)
	for i := 0; i < 9; i++ {
		if i != vA {
			constraints[v] = append(constraints[v], variable{i, vB})
		}
		if i != vB {
			constraints[v] = append(constraints[v], variable{vA, i})
		}
	}
	subI, subJ := vA/3, vB/3
	for i := subI * 3; i < (subI+1)*3; i++ {
		for j := subJ * 3; j < (subJ+1)*3; j++ {
			vN := variable{i, j}
			if v != vN {
				constraints[v] = append(constraints[v], vN)
			}
		}
	}

}
