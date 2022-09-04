class Solution {
public:
    vector<vector<int>> matrixReshape(vector<vector<int>>& mat, int r, int c) {
        int rr=mat.size(), cc=mat[0].size();
        if(rr*cc != r*c)
            return mat;
        vector< vector <int> > ans(r, vector <int> (c,0));
        for(int i=0; i<rr; i++){
            for(int j=0; j<cc; j++){
                int pos = i*cc + j;
                ans[pos/c][pos%c] = mat[i][j];
            }
        }
        return ans;
    }
};