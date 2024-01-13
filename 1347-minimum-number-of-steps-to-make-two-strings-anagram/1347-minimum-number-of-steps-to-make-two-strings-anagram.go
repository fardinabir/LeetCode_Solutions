func minSteps(s string, t string) int {
    mps := make(map[rune]int)
    mpt := make(map[rune]int)
    
    for i:= 0; i<len(s); i++ {
        mps[rune(s[i])]++
        mpt[rune(t[i])]++
    }
    
    var cnt int
    for key, val := range mpt {
        if val2,ok := mps[key]; ok {
            cnt += min(val, val2)
        }
    }
    return len(s) - cnt
}

func AbsInt(x int) int {
	return int(math.Abs(float64(x)))
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}