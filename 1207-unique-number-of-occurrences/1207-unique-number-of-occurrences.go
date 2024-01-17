func uniqueOccurrences(arr []int) bool {
    mpFreq := make(map[int]int)
    mpOccur := make(map[int]bool)
    for _, val := range arr {
        mpFreq[val]++
    }
    
    for _, val := range mpFreq {
        if mpOccur[val] {
            return false
        }
        mpOccur[val] = true
    }
    return true
}