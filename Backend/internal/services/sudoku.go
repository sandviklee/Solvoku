// Package services
// This is a sudoku interface for the CSP
package services

import (
	"errors"
	mapset "github.com/deckarep/golang-set/v2"
	"strconv"
)

// Public

type Board struct {
	numbers [][]int
}

type SudokuCtx struct {
	board       Board
	variables   []variable
	domains     map[variable]mapset.Set[int]
	constraints int
}

// Convert to String function for the Board struct to represent a Sudoku Board.
func (board Board) String() (boardString string) {
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

func InitSudokuCtx(board *Board) (sudokuCtx *SudokuCtx) {
	sudokuCtx.board = *board
	_, err := initSudokuVariables(sudokuCtx)

	if err != nil {
		return nil
	}

	_, err = initSudokuDomains(sudokuCtx)

	if err != nil {
		return nil
	}

	return
}

// Private

type variable struct {
	a, b interface{}
}

// Initializes Variables for a Sudoku Board and sets it in the SudokuCtx and returns
func initSudokuVariables(sudokuCtx *SudokuCtx) (variables []variable, err error) {
	if len(sudokuCtx.board.numbers) == 0 {
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
	if len(sudokuCtx.board.numbers) == 0 {
		return domains, errors.New("SudokuCtx not initialized.")
	}

	if sudokuCtx.variables == nil {
		return domains, errors.New("Variables not initialized")
	}

	for _, position := range sudokuCtx.variables {
		domains[position] = mapset.NewSet[int]()
	}

	for i, row := range sudokuCtx.board.numbers {
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
