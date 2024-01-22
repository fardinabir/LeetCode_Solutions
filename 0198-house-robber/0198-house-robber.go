var dp [101]int

func rob(nums []int) int {
    for i := range dp {
        dp[i] = -1
    }
    return calMaxMoney(&nums, 0)
}

func calMaxMoney(nums *[]int, ind int) int{
    if ind >= len(*nums) {
        return 0
    }
    if dp[ind] != -1 {
        return dp[ind]
    }
    dp[ind] = max(calMaxMoney(nums, ind+1), calMaxMoney(nums, ind+2) + (*nums)[ind])
    return dp[ind]
}