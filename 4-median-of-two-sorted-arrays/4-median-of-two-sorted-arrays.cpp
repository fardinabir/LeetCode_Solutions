class Solution {
public:
double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
    if(nums1.size()<nums2.size())
            return findMedianSortedArrays(nums2,nums1);
    int s1=nums1.size(),s2=nums2.size(),left1,left2,right1,right2;
    int lo=0,hi=s2,mid2,mid1;
    while(lo<=hi){
        mid2 = (lo+hi)/2;
        mid1 = (s1+s2+1)/2-mid2;
        
        left2 = mid2-1>=0? nums2[mid2-1] : INT_MIN;
        left1 = mid1-1>=0? nums1[mid1-1] : INT_MIN;
        right2 = mid2<s2? nums2[mid2] : INT_MAX;
        right1 = mid1<s1? nums1[mid1] : INT_MAX;
        
        if(left2<=right1 && left1<=right2){
            return (s1+s2)&1? max(left1,left2) : (max(left1,left2)+min(right1,right2))/2.0;
        }
        else if(left1>right2)
            lo = mid2+1;
        else
            hi = mid2-1;
    }
    return -1;
    }
};