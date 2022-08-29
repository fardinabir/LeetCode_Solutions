class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int mx=INT_MIN,c=0;
        for(int i=0;i<nums.size();i++){
            c+=nums[i];
            mx=max(c,mx);
            if(c<0)
                c=0;
        }
        return mx;
    }
};