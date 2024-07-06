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

	// Test Initialization of Variables
	sudokuTestVariables := TestVariables
	assert.Equal(suite.T(), sudokuTestVariables, sudokuCtx.variables, "Variables were not initialized correctly")

	// Test Initialization of Domains
	sudokuTestDomains := TestDomains
	assert.Equal(suite.T(), sudokuTestDomains, sudokuCtx.domains, "Domains were not initialized correctly")

	// Test Initialization of Constraints
	sudokuTestConstraints := TestConstraints
	assert.Equal(suite.T(), sudokuTestConstraints, sudokuCtx.constraints, "Constraints were not initialized correctly")
}

func TestSudokuTestSuite(t *testing.T) {
	suite.Run(t, new(SudokuTestSuite))

}
