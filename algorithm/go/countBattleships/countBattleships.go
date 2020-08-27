package countBattleships

func countBattleships(board [][]byte) int {
	count := 0

	y_len := len(board[0])
	x_len := len(board)

	for y := 0; y < y_len; y++ {
		for x := 0; x < x_len; x++ {
			if board[x][y] != 'X' {
				continue
			}

			right_x := x + 1
			down_y := y + 1
			if right_x < x_len {
				if board[right_x][y] == 'X' {
					continue
				}
			}

			if down_y < y_len {
				if board[x][down_y] == 'X' {
					continue
				}
			}

			count++
		}
	}

	return count
}
