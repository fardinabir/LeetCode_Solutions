const mod = int(1e9) + 7

func kInversePairs(n int, k int) int {
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, k+1)
    }

    for i := 0; i <= n; i++ {
        dp[i][0] = 1
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= k; j++{
            sub := 0
            if j >= i {
                sub = dp[i-1][j-i]
            }
            dp[i][j] = ((dp[i-1][j] + dp[i][j-1]) + mod - sub) % mod;
        }
    }
    return dp[n][k]
}