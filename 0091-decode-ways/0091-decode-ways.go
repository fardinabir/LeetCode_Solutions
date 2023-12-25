var mem map[string]int

func numDecodings(s string) int {
    mem = make(map[string]int)
    return calDecodings(0, s)
}

func calDecodings(i int, s string) int {
    if len(s) <= i {
        // fmt.Println(cons)
        return 1
    }
    if s[i] == '0' {
        return 0
    }
    value, ok := mem[s[i:]]
    if ok {
        return value
    }
    var nextDigit, oneDigit int
    if s[i] != '0' {
        // fmt.Println("-----", s[:i+1])
        oneDigit = calDecodings(i+1, s)
    }
    if i+1 < len(s)  {
        num, _ := strconv.Atoi(s[i:i+2])
        if num < 27 && num > 0 {
        // fmt.Println("-----", cons)
        nextDigit = calDecodings(i+2, s)
        }
    }
    mem[s[i:]] = oneDigit + nextDigit
    return oneDigit + nextDigit
}