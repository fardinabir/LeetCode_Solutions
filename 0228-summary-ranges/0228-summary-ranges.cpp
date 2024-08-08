class Solution {
public:
    vector<string> summaryRanges(vector<int>& nums) {
        vector <string> vc;
        int l=0;
        if(nums.empty()){
            return vc;
        }
        for(int i=0; i<nums.size(); i++) {
            if(i == nums.size() - 1 || nums[i] != nums[i+1] - 1){
                if(l == i){
                    vc.push_back(to_string(nums[i]));
                } else {
                    string st, r;
                    st = to_string(nums[l]);
                    r = to_string(nums[i]);
                    vc.push_back(st+"->"+r);
                }
                l = i+1;
            }       
        }
        return vc;
    }
};