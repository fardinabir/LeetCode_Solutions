class Solution {
public:
    vector<string> summaryRanges(vector<int>& nums) {
        vector <string> vc;
        int l=0;
        for(int i=0; i<nums.size(); i++) {
            if(i == nums.size() - 1 || nums[i] != nums[i+1] - 1){
                if(l == i){
                    vc.push_back(to_string(nums[i]));
                } else {
                    vc.push_back(to_string(nums[l])+"->"+to_string(nums[i]));
                }
                l = i+1;
            }       
        }
        return vc;
    }
};