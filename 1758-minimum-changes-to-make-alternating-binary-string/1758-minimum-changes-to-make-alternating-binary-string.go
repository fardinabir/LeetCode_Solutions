func minOperations(s string) int {
    cntEven, cntOdd := 0, 0
    
    for i, x := range s {
        if i%2 == 1 {
            if x == '1' {
                cntOdd++
            } else {
                cntEven++
            }
        } else {
            if x == '0' {
                cntOdd++
            } else {
                cntEven++
            }
        }
    }
    return Min(cntOdd, cntEven)
}

func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}