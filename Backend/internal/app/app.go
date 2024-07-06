// Package app implements all the internal features into exportable functions.
package app

import (
	"fmt"
	"github.com/sandviklee/solvoku/internal/services"
)

// Initialize initializes the main program.
func Run() {
	csp := services.CSP{}
	board := services.Board{}
	boardNumbers := [][]int{{0, 0, 0, 0, 0, 0, 0, 0, 9}, {3, 7, 0, 1, 0, 0, 0, 0, 0},
		{0, 2, 9, 6, 0, 0, 0, 1, 0}, {0, 1, 0, 7, 4, 0, 0, 0, 0}, {7, 5, 0, 0, 0, 0, 0, 6, 3},
		{0, 0, 0, 0, 3, 2, 0, 7, 0}, {0, 6, 0, 0, 0, 7, 5, 2, 0}, {0, 0, 0, 0, 0, 8, 0, 9, 7},
		{4, 0, 0, 0, 0, 0, 0, 0, 0}}
	board.SetNumbers(boardNumbers)
	sudokuCtx, _ := services.InitSudokuCtx(board)
	csp.InitCSPWithSudoku(sudokuCtx)
	solution := services.Solve(csp)
	boardSolution := services.TransformSolutionToBoard(solution)
	fmt.Println(boardSolution.String())
}
