class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int ind = 0;
        for(int i=0; i<nums.size(); i++){
            if(nums[i]!=nums[ind])
                nums[++ind] = nums[i];
        }
        return ind+1;
    }
};