package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SudokuTestSuite struct {
	suite.Suite
	sudokuBoard Board
	sudokuCtx   SudokuCtx
}

func (suite *SudokuTestSuite) SetupTest() {
	suite.sudokuBoard.numbers = [][]int{{5, 3, 0, 0, 7, 0, 0, 0, 0}, {6, 0, 0, 1, 9, 5, 0, 0, 0}, {0, 9, 8, 0, 0, 0, 0, 6, 0}, {8, 0, 0, 0, 6, 0, 0, 0, 3}, {4, 0, 0, 8, 0, 3, 0, 0, 1}, {7, 0, 0, 0, 2, 0, 0, 0, 6}, {0, 6, 0, 0, 0, 0, 2, 8, 0}, {0, 0, 0, 4, 1, 9, 0, 0, 5}, {0, 0, 0, 0, 8, 0, 0, 0, 0}}
	suite.sudokuCtx.board = suite.sudokuBoard
}

func (suite *SudokuTestSuite) TestBoardPrint() {
	sudokuTestBoardString := "5 3 0 | 0 7 0 | 0 0 0 \n6 0 0 | 1 9 5 | 0 0 0 \n0 9 8 | 0 0 0 | 0 6 0 \n- - - - - - - - - - -\n8 0 0 | 0 6 0 | 0 0 3 \n4 0 0 | 8 0 3 | 0 0 1 \n7 0 0 | 0 2 0 | 0 0 6 \n- - - - - - - - - - -\n0 6 0 | 0 0 0 | 2 8 0 \n0 0 0 | 4 1 9 | 0 0 5 \n0 0 0 | 0 8 0 | 0 0 0 \n"
	board := Board{}
	assert.Equal(suite.T(), board.String(), sudokuTestBoardString)
}

func (suite *SudokuTestSuite) TestInitSudokuCtx() {

}

func TestSudokuTestSuite(t *testing.T) {
	suite.Run(t, new(SudokuTestSuite))
}
