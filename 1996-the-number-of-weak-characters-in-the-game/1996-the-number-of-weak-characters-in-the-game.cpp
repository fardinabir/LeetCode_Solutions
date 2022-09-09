class Solution {
public:
    
    int numberOfWeakCharacters(vector<vector<int>>& prop) {
        sort(prop.begin(),prop.end(), [](vector<int> &a, vector<int> &b){
        if(a[0] == b[0])
            return a[1]>b[1];
        return a[0]<b[0];
    });
        
        int ans = 0, mx = prop[prop.size()-1][1];
        for(int i = (int) prop.size()-2; i>=0; i--){
            if(prop[i][1] < mx)
                ans++;
            else
                mx = prop[i][1];
        }
        return ans;
    }
};