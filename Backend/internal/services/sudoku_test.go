package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SudokuTestSuite struct {
	suite.Suite
	sudokuBoard Board
}

func (suite *SudokuTestSuite) SetupTest() {
	suite.sudokuBoard.numbers = [][]int{{5, 3, 0, 0, 7, 0, 0, 0, 0}, {6, 0, 0, 1, 9, 5, 0, 0, 0}, {0, 9, 8, 0, 0, 0, 0, 6, 0}, {8, 0, 0, 0, 6, 0, 0, 0, 3}, {4, 0, 0, 8, 0, 3, 0, 0, 1}, {7, 0, 0, 0, 2, 0, 0, 0, 6}, {0, 6, 0, 0, 0, 0, 2, 8, 0}, {0, 0, 0, 4, 1, 9, 0, 0, 5}, {0, 0, 0, 0, 8, 0, 0, 0, 0}}
}

func (suite *SudokuTestSuite) TestBoardPrint() {
	sudokuTestBoardString := "5 3 0 | 0 7 0 | 0 0 0 \n6 0 0 | 1 9 5 | 0 0 0 \n0 9 8 | 0 0 0 | 0 6 0 \n- - - - - - - - - - -\n8 0 0 | 0 6 0 | 0 0 3 \n4 0 0 | 8 0 3 | 0 0 1 \n7 0 0 | 0 2 0 | 0 0 6 \n- - - - - - - - - - -\n0 6 0 | 0 0 0 | 2 8 0 \n0 0 0 | 4 1 9 | 0 0 5 \n0 0 0 | 0 8 0 | 0 0 0 \n"
	board := suite.sudokuBoard
	assert.Equal(suite.T(), sudokuTestBoardString, board.String())
}

func (suite *SudokuTestSuite) TestInitSudokuCtx() {
	board := suite.sudokuBoard
	sudokuCtx, _ := InitSudokuCtx(board)

	sudokuTestVarables := []variable{
		{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}, {0, 8},
		{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}, {1, 8},
		{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4}, {2, 5}, {2, 6}, {2, 7}, {2, 8},
		{3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4}, {3, 5}, {3, 6}, {3, 7}, {3, 8},
		{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {4, 5}, {4, 6}, {4, 7}, {4, 8},
		{5, 0}, {5, 1}, {5, 2}, {5, 3}, {5, 4}, {5, 5}, {5, 6}, {5, 7}, {5, 8},
		{6, 0}, {6, 1}, {6, 2}, {6, 3}, {6, 4}, {6, 5}, {6, 6}, {6, 7}, {6, 8},
		{7, 0}, {7, 1}, {7, 2}, {7, 3}, {7, 4}, {7, 5}, {7, 6}, {7, 7}, {7, 8},
		{8, 0}, {8, 1}, {8, 2}, {8, 3}, {8, 4}, {8, 5}, {8, 6}, {8, 7}, {8, 8},
	}

	sudokuTestConstraints := map[variable][]int{{0, 0}: {5}}

	// Test Initialization of Variables
	assert.Equal(suite.T(), sudokuTestVarables, sudokuCtx.variables, "Variables were not initialized correctly")
}

func TestSudokuTestSuite(t *testing.T) {
	suite.Run(t, new(SudokuTestSuite))
}
