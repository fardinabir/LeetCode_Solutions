class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
        int x,i,c=1;
        vector <int> vc;
        for(i=digits.size()-1;i>=0;i--)
        {
            if(c==1)
            {
                digits[i]+=c;
                c=digits[i]/10;
                digits[i]=digits[i]%10;
            }
            if(c==0)
                break;
        }
        if(c==1)
            vc.push_back(1);
        for(i=0;i<digits.size();i++)
            vc.push_back(digits[i]);
        return vc;
    }
};