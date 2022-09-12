class Solution {
public:
    void rotate(vector<int>& nums, int k) {
        int tmp,r=k%nums.size();
        for(int i=0; i<(nums.size()-r)/2; i++)
            swap(nums[i], nums[nums.size()-r-i-1]);
        for(int i=0; i<r/2; i++)
            swap(nums[i+nums.size()-r],nums[nums.size()-i-1]);
        for(int i=0; i<nums.size()/2; i++)
            swap(nums[i], nums[nums.size()-i-1]);
    }
};