package problem

func setZeroes(matrix [][]int) {
	// 时间复杂度O(m*n)
	// 空间复杂度O(m*n)
	if len(matrix) == 0 {
		return
	}
	row, col := len(matrix), len(matrix[0])
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 && !visited[i][j] {
				visited[i][j] = true
				setValue(matrix, i, col, j, row, visited)
			}
		}
	}
}

func setValue(matrix [][]int, i, col, j, row int, visited [][]bool) {
	for k := 0; k < col; k++ {
		if matrix[i][k] != 0 {
			matrix[i][k] = 0
			visited[i][k] = true
		}
	}
	for k := 0; k < row; k++ {
		if matrix[k][j] != 0 {
			matrix[k][j] = 0
			visited[k][j] = true
		}
	}
}
