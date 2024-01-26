const M = 1000000007

func findPaths(m, n, N, i, j int) int {
	memo := make([][][]int, m)
	for x := 0; x < m; x++ {
		memo[x] = make([][]int, n)
		for y := 0; y < n; y++ {
			memo[x][y] = make([]int, N+1)
			for z := 0; z <= N; z++ {
				memo[x][y][z] = -1
			}
		}
	}
	return findPathsHelper(m, n, N, i, j, memo)
}

func findPathsHelper(m, n, N, i, j int, memo [][][]int) int {
	if i == m || j == n || i < 0 || j < 0 {
		return 1
	}
	if N == 0 {
		return 0
	}
	if memo[i][j][N] >= 0 {
		return memo[i][j][N]
	}
	memo[i][j][N] = (
		(findPathsHelper(m, n, N-1, i-1, j, memo) + findPathsHelper(m, n, N-1, i+1, j, memo)) % M +
        (findPathsHelper(m, n, N-1, i, j-1, memo) + findPathsHelper(m, n, N-1, i, j+1, memo)) % M) % M
	return memo[i][j][N]
}
