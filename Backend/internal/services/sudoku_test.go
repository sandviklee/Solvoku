package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Unit Testing
func TestBoardPrint(t *testing.T) {
	sudokuTestBoard := [][]int{{5, 3, 0, 0, 7, 0, 0, 0, 0}, {6, 0, 0, 1, 9, 5, 0, 0, 0}, {0, 9, 8, 0, 0, 0, 0, 6, 0}, {8, 0, 0, 0, 6, 0, 0, 0, 3}, {4, 0, 0, 8, 0, 3, 0, 0, 1}, {7, 0, 0, 0, 2, 0, 0, 0, 6}, {0, 6, 0, 0, 0, 0, 2, 8, 0}, {0, 0, 0, 4, 1, 9, 0, 0, 5}, {0, 0, 0, 0, 8, 0, 0, 0, 0}}
	sudokuTestBoardString := "5 3 0 | 0 7 0 | 0 0 0 \n6 0 0 | 1 9 5 | 0 0 0 \n0 9 8 | 0 0 0 | 0 6 0 \n- - - - - - - - - - -\n8 0 0 | 0 6 0 | 0 0 3 \n4 0 0 | 8 0 3 | 0 0 1 \n7 0 0 | 0 2 0 | 0 0 6 \n- - - - - - - - - - -\n0 6 0 | 0 0 0 | 2 8 0 \n0 0 0 | 4 1 9 | 0 0 5 \n0 0 0 | 0 8 0 | 0 0 0 \n"
	board := Board{sudokuTestBoard}

	assert.Equal(t, board.String(), sudokuTestBoardString)

}
