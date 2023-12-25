var mem []int

func numDecodings(s string) int {
    mem = make([]int, len(s)+1)
    return calDecodings(0, s)
}

func calDecodings(i int, s string) int {
    if len(s) <= i {
        return 1
    }
    if s[i] == '0' {
        return 0
    }
    if mem[i] != 0 {
        return mem[i]
    }
    oneDigit := calDecodings(i+1, s)
    if i+1 < len(s) && (s[i] == '1' || (s[i] == '2' && s[i+1] <= '6')) {
        oneDigit += calDecodings(i+2, s)
    }
    mem[i] = oneDigit
    return oneDigit
}
