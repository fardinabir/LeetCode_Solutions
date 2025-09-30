func findDuplicate(nums []int) int {
    n := len(nums)
    for i := range nums {
        for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }
    return nums[n-1]

}