class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        int i, j = -1;
        for(i = 0; i < nums.size(); i++){
            if(nums[i] && j>=0)
                swap(nums[i],nums[j++]);
            else if(nums[i]==0 && j < 0)
                j = i;
        }
    }
};