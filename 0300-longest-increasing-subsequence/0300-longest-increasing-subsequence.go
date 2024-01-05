func lengthOfLIS(nums []int) int {
    lisCnt := make([]int, len(nums))
    
    maxLen := 0
    for i, _ := range nums {
        for j:=0; j<i; j++ {
            if nums[j] < nums[i] {
                lisCnt[i] = max(lisCnt[i], lisCnt[j] + 1)
            }
        }
        maxLen = max(maxLen, lisCnt[i])
    }
    return maxLen + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}