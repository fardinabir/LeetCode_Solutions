class Solution {
public:
    int minimumOperations(vector<int>& nums) {
        unordered_set <int> ust;
        for(int i=0;i<nums.size();i++){
            if(nums[i])
                ust.insert(nums[i]);
        }
        return ust.size();
    }
};