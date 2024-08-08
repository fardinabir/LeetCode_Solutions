class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        int i, j = 0;
        for(i = 0; i < nums.size(); i++){
            if(nums[i])
                swap(nums[i],nums[j++]);
        }
    }
};