func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			mapping := make(map[int]int)
			mapping[0] = 1
			sum := 0

			for k := 0; k < m; k++ {
				sum += matrix[k][j]
				if i > 0 {
					sum -= matrix[k][i-1]
				}

				if val, exists := mapping[sum-target]; exists {
					count += val
				}

				if _, exists := mapping[sum]; !exists {
					mapping[sum] = 0
				}

				mapping[sum]++
			}
		}
	}

	return count
}
