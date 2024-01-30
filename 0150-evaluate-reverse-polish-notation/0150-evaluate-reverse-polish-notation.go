func evalRPN(tokens []string) int {
    var stack []int
    for _, token := range tokens {
        if token == "+" {
            b := stack[len(stack)-1]
            a := stack[len(stack)-2]
            stack[len(stack)-2] = a+b
            stack = stack[:len(stack)-1]
        } else if token == "-" {
            b := stack[len(stack)-1]
            a := stack[len(stack)-2]
            stack[len(stack)-2] = a-b
            stack = stack[:len(stack)-1]
        } else if token == "*" {
            b := stack[len(stack)-1]
            a := stack[len(stack)-2]
            stack[len(stack)-2] = a*b
            stack = stack[:len(stack)-1]
        } else if token == "/" {
            b := stack[len(stack)-1]
            a := stack[len(stack)-2]
            stack[len(stack)-2] = a/b
            stack = stack[:len(stack)-1]
        } else {
            val, _ := strconv.Atoi(token)
            stack = append(stack, val)
        }
    }
    fmt.Println(stack)
    return stack[0]
}