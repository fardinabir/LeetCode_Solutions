func findErrorNums(nums []int) []int {
    var ans []int
    freq := make([]int, len(nums))
    for _, x := range nums {
        freq[x-1]++
    }
    var dup, zero int
    for i, x := range freq {
        if x == 2 {
            dup = i+1
        } else if x == 0 {
            zero = i+1
        }
    }
    return append(ans, dup, zero)
}