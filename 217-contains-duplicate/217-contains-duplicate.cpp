class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        unordered_map <int,bool> ump;
        for(int i=0;i<nums.size();i++){
            if(ump.find(nums[i])!=ump.end())
                return true;
            ump[nums[i]]=true;
        }
        return false;
    }
};