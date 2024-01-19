var dp [101][101]int

func memSet(dp *[101][101]int){
    for i := 0; i < 100; i++ {
        for j := 0; j < 100; j++ {
            (*dp)[i][j] = -math.MaxInt
        }
    }
}

func minFallingPathSum(matrix [][]int) int {
    mn := math.MaxInt
    memSet(&dp)
    for i, _ := range matrix[0] {
        mn = min(mn, calFallingPath(&matrix, 0, i))
    }
    return mn
}

func calFallingPath(matrix *[][]int, row, col int) int {
    if row == len(*matrix)-1 {
        return (*matrix)[row][col]
    }
    if dp[row][col] != -math.MaxInt {
        return dp[row][col]
    }
    p1 := math.MaxInt
    p2 := math.MaxInt
    p3 := math.MaxInt
    if row+1 < len(*matrix) && col > 0 {
        p1 = calFallingPath(matrix, row+1, col-1)
    }
    if row+1 < len(*matrix) {
        p2 = calFallingPath(matrix, row+1, col)
    }
    if row+1 < len(*matrix) && col+1 < len((*matrix)[0]) {
        p3 = calFallingPath(matrix, row+1, col+1)
    }
    dp[row][col] = min(min(p1, p2), p3) + (*matrix)[row][col]
    return dp[row][col]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}