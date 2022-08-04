class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        int i,s;
        s = nums.size();
        unordered_map <int,bool> ump;
        for(i=0;i<s;i++){
            if(ump.find(nums[i])!=ump.end())
                return true;
            ump[nums[i]]=true;
        }
        return false;
    }
};