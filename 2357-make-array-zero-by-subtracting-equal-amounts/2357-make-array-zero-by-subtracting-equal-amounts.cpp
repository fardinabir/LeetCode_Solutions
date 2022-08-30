class Solution {
public:
    int minimumOperations(vector<int>& nums) {
        sort(nums.begin(),nums.end());
        int carry=0,s=0,i=0;
        for(i=0;i<nums.size();i++){
            if(nums[i]!=0 && nums[i]>carry){
                carry+=(nums[i]-carry);
                s++;
            }
        }
        return s;
    }
};