func missingNumber(nums []int) int {
    var sum int
    n := len(nums)
    for _, v := range nums {
        sum = sum + v
    }
    return ((n+1)*n)/2 - sum
}