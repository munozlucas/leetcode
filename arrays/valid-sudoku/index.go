package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	const boxLen = 9
	row := [boxLen][boxLen]int{}
	col := [boxLen][boxLen]int{}
	box := [boxLen][boxLen]int{}

	for c := 0; c < boxLen; c++ {
		for r := 0; r < boxLen; r++ {
			val := board[r][c]

			if val != '.' {
				valInSet := val - 1
				// check column
				if col[c][valInSet] == 1 {
					return false
				}
				col[c][valInSet] = 1

				// check row
				if row[r][valInSet] == 1 {
					return false
				}
				row[r][valInSet] = 1

				// check box
				boxIndex := (r/3)*3 + c/3 // box index = row * rowSize + col
				if box[boxIndex][valInSet] == 1 {
					return false
				}
				box[boxIndex][valInSet] = 1
			}
		}
	}

	return true
}

func main() {
	validSudoku := [][]byte{
		{5, 3, '.', '.', 7, '.', '.', '.', '.'},
		{6, '.', '.', 1, 9, 5, '.', '.', '.'},
		{'.', 9, 8, '.', '.', '.', '.', 6, '.'},
		{8, '.', '.', '.', 6, '.', '.', '.', 3},
		{4, '.', '.', 8, '.', 3, '.', '.', 1},
		{7, '.', '.', '.', 2, '.', '.', '.', 6},
		{'.', 6, '.', '.', '.', '.', 2, 8, '.'},
		{'.', '.', '.', 4, 1, 9, '.', '.', 5},
		{'.', '.', '.', '.', 8, '.', '.', 7, 9},
	}
	valid := isValidSudoku(validSudoku)

	fmt.Println(valid)

	invalidSudoku := [][]byte{
		{5, 5, '.', '.', 7, '.', '.', '.', '.'},
		{6, '.', '.', 1, 9, 5, '.', '.', '.'},
		{'.', 9, 8, '.', '.', '.', '.', 6, '.'},
		{8, '.', '.', '.', 6, '.', '.', '.', 3},
		{4, '.', '.', 8, '.', 3, '.', '.', 1},
		{7, '.', '.', '.', 2, '.', '.', '.', 6},
		{'.', 6, '.', '.', '.', '.', 2, 8, '.'},
		{'.', '.', '.', 4, 1, 9, '.', 1, 5},
		{'.', '.', '.', '.', 8, '.', '.', 7, 9},
	}
	invalid := isValidSudoku(invalidSudoku)

	fmt.Println(invalid)

}
