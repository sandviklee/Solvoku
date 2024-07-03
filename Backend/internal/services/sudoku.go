// Package services
// This is a sudoku interface for the CSP
package services

import (
	"errors"
	mapset "github.com/deckarep/golang-set/v2"
)

// Public

type Board struct {
	numbers [9][9]int
}

type SudokuCtx struct {
	Board       Board
	variables   []variable
	domains     map[variable]mapset.Set[int]
	constraints int
}

func InitSudokuCtx(board *Board) (sudokuCtx *SudokuCtx) {
	sudokuCtx.Board = *board
	return
}

// Private

type variable struct {
	a, b interface{}
}

// Initializes Variables for a Sudoku Board and sets it in the SudokuCtx and returns
func initSudokuVariables(sudokuCtx *SudokuCtx) (variables []variable, err error) {
	if len(sudokuCtx.Board.numbers) == 0 {
		return variables, errors.New("SudokuCtx not initialized.")
	}

	maxCells := 9
	for i := 0; i < maxCells; i++ {
		for j := 0; j < maxCells; j++ {
			variables = append(variables, variable{i, j})
		}
	}

	sudokuCtx.variables = variables
	return
}

// Initializes Domains for Sudoku Board and sets it in the SudokuCtx and returns
// Depends on initialization of variable beforehand
func initSudokuDomains(sudokuCtx *SudokuCtx) (domains map[variable]mapset.Set[int], err error) {
	if len(sudokuCtx.Board.numbers) == 0 {
		return domains, errors.New("SudokuCtx not initialized.")
	}

	if sudokuCtx.variables == nil {
		return domains, errors.New("Variables not initialized")
	}

	for _, position := range sudokuCtx.variables {
		domains[position] = mapset.NewSet[int]()
	}

	for i, row := range sudokuCtx.Board.numbers {
		for j, cell := range row {
			position := variable{i, j}

			if cell != 0 {
				domains[position].Add(cell)
				continue
			}

			for k := 1; k < 10; k++ {
				domains[position].Add(k)
			}

		}
	}

	sudokuCtx.domains = domains
	return

}
