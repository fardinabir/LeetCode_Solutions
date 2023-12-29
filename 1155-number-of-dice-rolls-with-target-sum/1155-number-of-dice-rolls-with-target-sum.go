const mod = 1000000007

func numRollsToTarget(n, k, target int) int {
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, target+1)
    }

    dp[0][0] = 1

    for i := 1; i <= n; i++ {
        for j := 1; j <= target; j++ {
            for t := 1; t <= k; t++ {
                if j-t >= 0 {
                    dp[i][j] = (dp[i][j] + dp[i-1][j-t]) % mod
                }
            }
        }
    }

    return dp[n][target]
}
