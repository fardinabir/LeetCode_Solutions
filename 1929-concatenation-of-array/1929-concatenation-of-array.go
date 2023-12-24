func getConcatenation(nums []int) []int {
    lenght := len(nums)
    for i := 0; i < lenght; i++ {
        nums = append(nums, nums[i%lenght]) 
    }
    return nums
}