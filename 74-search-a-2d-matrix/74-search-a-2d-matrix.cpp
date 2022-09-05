class Solution {
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        int colSize = matrix[0].size(), sz = matrix.size() * matrix[0].size();
        int lo = 0, hi = sz-1, mid;
        while(lo <= hi){
            mid = (lo + hi)>>1;
            if(matrix[mid/colSize][mid%colSize] == target)
                return true;
            else if(matrix[mid/colSize][mid%colSize] < target)
                lo = mid + 1;
            else
                hi = mid - 1;
        }
        return false;
    }
};