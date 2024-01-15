func findWinners(matches [][]int) [][]int {
    mapOfLosers := make(map[int]int)
    mapOfWinners := make(map[int]int)
    for _, match := range matches {
        mapOfWinners[match[0]]++
        mapOfLosers[match[1]]++
    }
    ans := make([][]int, 2)
    for key, val := range mapOfLosers {
        if val == 1 {
            ans[1] = append(ans[1], key)
        }
    }
    for key, _ := range mapOfWinners {
        if mapOfLosers[key] == 0 {
            ans[0] = append(ans[0], key)
        }
    }
    sort.Ints(ans[0])
    sort.Ints(ans[1])
    return ans
}