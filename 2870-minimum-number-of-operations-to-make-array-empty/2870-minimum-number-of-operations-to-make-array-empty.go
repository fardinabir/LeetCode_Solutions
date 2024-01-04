func minOperations(nums []int) int {
    numMap := map[int]int{}
    for _, x := range nums {
        numMap[x]++
    }
    var cnt int
    for _, value := range numMap {
        if value < 2 {
            return -1
        }
        cnt += value/3
        if value%3 != 0 {
            cnt++
        }
    }
    return cnt
}