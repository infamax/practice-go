package warriors

func convertImage(image string) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0)
	for _, ch := range image {
		if ch != '\n' {
			tmp = append(tmp, int(ch)-'0')
		} else {
			res = append(res, tmp)
			tmp = nil
		}
	}
	res = append(res, tmp)
	return res
}

func DFS(matrix [][]int, i, j, n int) {
	if matrix[i][j] == 0 {
		return
	}

	matrix[i][j] = 0

	if i+1 < n {
		DFS(matrix, i+1, j, n)
	}

	if i-1 >= 0 {
		DFS(matrix, i-1, j, n)
	}

	if j-1 >= 0 {
		DFS(matrix, i, j-1, n)
	}

	if j+1 < n {
		DFS(matrix, i, j+1, n)
	}

	if i-1 >= 0 && j-1 >= 0 {
		DFS(matrix, i-1, j-1, n)
	}

	if i-1 >= 0 && j+1 < n {
		DFS(matrix, i-1, j+1, n)
	}

	if i+1 < n && j-1 >= 0 {
		DFS(matrix, i+1, j-1, n)
	}

	if i+1 < n && j+1 < n {
		DFS(matrix, i+1, j+1, n)
	}
}

func Count(image string) int {
	matrix := convertImage(image)
	n := len(matrix)
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				count++
				DFS(matrix, i, j, n)
			}
		}
	}
	return count
}
