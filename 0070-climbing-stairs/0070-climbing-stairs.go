var dp []int

func climbStairs(n int) int {
    dp = make([]int,46)
    return calc(n)
}

func calc(n int) int {
    if n <= 0 {
        if n == 0 {
            return 1
        }
        return 0
    }
    if dp[n] != 0 {
        return dp[n]
    }
    dp[n] = calc(n-1) + calc(n-2)
    return dp[n]
}