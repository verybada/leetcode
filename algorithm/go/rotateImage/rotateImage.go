package rotateimage

func rotate(matrix [][]int) {
	width := len(matrix)
	for i := 0; i < width; i++ {
		for j := i + 1; j < width; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < width; i++ {
		for j := 0; j < width/2; j++ {
			matrix[i][j], matrix[i][width-j-1] = matrix[i][width-j-1], matrix[i][j]
		}
	}
}
