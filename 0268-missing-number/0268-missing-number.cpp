class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int flag[10005]={0}, i;
        for(i=0; i<nums.size(); i++) {
            flag[nums[i]] = 1;
        }
        for(i=0; i<nums.size()+1; i++) {
            if(flag[i] == 0){
                break;
            }
        }
        return i;
    }
};