class Solution {
public:
    vector<int> intersect(vector<int>& nums1, vector<int>& nums2) {
        int f1[1001],f2[1002],i,j;
        vector <int> ans;
        memset(f1,0,sizeof f1);
        memset(f2,0,sizeof f2);
        for(i=0;i<nums1.size();i++)
        {
            f1[nums1[i]]++;
        }
        for(i=0;i<nums2.size();i++)
        {
            f2[nums2[i]]++;
        }
        for(i=0;i<=1000;i++)
        {
            for(j=0;j<min(f1[i],f2[i]);j++)
                ans.push_back(i);
        }
        return ans;
    }
};