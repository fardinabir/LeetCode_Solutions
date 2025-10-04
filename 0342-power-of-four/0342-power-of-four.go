func isPowerOfFour(n int) bool {
    for n > 3 && n&1 == 0 {
        if n%4 != 0 {
            return false
        }
        n = n >> 2
    }
    return n == 1
}