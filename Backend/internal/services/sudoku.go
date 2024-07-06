// Package services
// This is a sudoku interface for the CSP
package services

import (
	"errors"
	"strconv"
)

type Board struct {
	numbers [][]int
}

type variable struct {
	a, b int
}

type sudokuCtx struct {
	board       Board
	variables   []variable
	domains     map[variable][]int
	constraints map[variable][]variable
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

// Transforms the sudoku solution into a Board
func TransformSolutionToBoard(solution map[variable]int) (board Board) {
	var numbers [][]int

	for i := 0; i < 9; i++ {
		numbers = append(numbers, make([]int, 9))
	}

	for key, value := range solution {
		numbers[key.a][key.b] = value
	}
	board.SetNumbers(numbers)
	return
}

func InitSudokuCtx(board Board) (ctx sudokuCtx, err error) {
	ctx.board = board

	if len(board.numbers) <= 0 {
		return ctx, errors.New("board is empty")
	}

	ctx.initSudokuVariables()
	ctx.initSudokuDomains()
	ctx.initSudokuConstraints()

	return
}

// Initializes Variables for a Sudoku Board and sets it in the SudokuCtx
// Basically a Setter...
func (ctx *sudokuCtx) initSudokuVariables() {
	variables := make([]variable, 0)
	maxCells := 9
	for i := 0; i < maxCells; i++ {
		for j := 0; j < maxCells; j++ {
			variables = append(variables, variable{i, j})
		}
	}

	ctx.variables = variables
}

// Initializes Domains for Sudoku Board and sets it in the SudokuCtx
// Depends on initialization of variable beforehand
func (ctx *sudokuCtx) initSudokuDomains() {
	domains := make(map[variable][]int)

	for _, position := range ctx.variables {
		domains[position] = make([]int, 0)
	}

	for i, row := range ctx.board.numbers {
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

	ctx.domains = domains
}

// Initializes the Constraints for Sudoku Board and sets it in the SudokuCtx
func (ctx *sudokuCtx) initSudokuConstraints() {
	constraints := make(map[variable][]variable)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			addConstraints(variable{i, j}, constraints)
		}
	}

	ctx.constraints = constraints

}

// Helper function to add the constraints
func addConstraints(v variable, constraints map[variable][]variable) {
	vA := v.a
	vB := v.b
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
