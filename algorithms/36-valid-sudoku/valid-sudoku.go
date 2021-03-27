package _6_valid_sudoku

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			// 检查横排
			for n := 0; n < i; n++ {
				if board[n][j] == board[i][j] {
					return false
				}
			}
			// 检查竖排
			for m := 0; m < j; m++ {
				if board[i][m] == board[i][j] {
					return false
				}
			}
			// 检查同一个九宫格
			iStart := (i / 3) * 3
			jStart := (j / 3) * 3
			for n := iStart; n < iStart+3; n++ {
				for m := jStart; m < jStart+3; m++ {
					if n == i && m == j {
						continue
					}
					if board[n][m] == board[i][j] {
						return false
					}
				}
			}
		}
	}
	return true
}
