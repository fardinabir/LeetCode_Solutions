class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int i=0,j=0,prev,sz=nums1.size()+nums2.size(),med=sz/2;
        vector <int> srt;
        while(i+j<=med && i<nums1.size() && j<nums2.size()){
            if(nums1[i]<nums2[j]){
                srt.push_back(nums1[i++]);
            }
            else{
                srt.push_back(nums2[j++]);
            }
        }
        while(i>=nums1.size() && i+j<=med){
            srt.push_back(nums2[j++]);
        }
        while(j>=nums2.size() && i+j<=med){
            srt.push_back(nums1[i++]);
        }
        if(sz&1){
           return srt.back(); 
        }
        else{
            return (srt[srt.size()-1]+srt[srt.size()-2])/2.0;
        }
    }
};