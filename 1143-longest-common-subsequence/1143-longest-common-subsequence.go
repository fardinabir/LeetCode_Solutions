var dp [1000][1000]int

func initializeDp() {
    for i := 0; i < len(dp); i++ {
        for j := 0; j < len(dp[i]); j++ {
            dp[i][j] = -1
        }
    }
}

func longestCommonSubsequence(text1 string, text2 string) int {
    initializeDp()
    return findLCS(&text1, &text2, len(text1)-1, len(text2)-1)
}

func findLCS(str1, str2 *string, ind1, ind2 int) int {
    if ind1 == -1 || ind2 == -1 {
        return 0
    }
    if dp[ind1][ind2] != -1 {
        return dp[ind1][ind2]
    }
    if (*str1)[ind1] == (*str2)[ind2] {
        return 1+findLCS(str1,str2, ind1-1, ind2-1)
    }
    dp[ind1][ind2] = max(findLCS(str1,str2, ind1-1, ind2),findLCS(str1,str2, ind1, ind2-1))
    return dp[ind1][ind2]
}