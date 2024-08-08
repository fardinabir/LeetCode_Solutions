class Solution {
public:
    int thirdMax(vector<int>& nums) {
        long long int first, second, third, cnt = 0;
        first = second = third = -2147483649;
        for(int i=0; i<nums.size(); i++){
            if(nums[i]>first){
                third = second;
                second = first;
                first = nums[i];
            } else if(nums[i]>second && nums[i]<first){
                third = second;
                second = nums[i];
            } else if(nums[i]>third && nums[i]<second){
                third = nums[i];
            }
        }
        if(third == -2147483649){
            return first;
        }
        return third;
    }
};