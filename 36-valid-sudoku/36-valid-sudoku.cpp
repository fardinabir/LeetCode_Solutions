class Solution {
public:
    bool isValidSudoku(vector<vector<char>>& board) {
        unordered_map <int,bool> row[9],col[9],box[9];
        for(int i=0; i<9; i++){
            for(int j=0; j<9; j++){
                if(board[i][j]!='.'){
                    int val = board[i][j];
                    if(row[i][val] || col[j][val] || box[(i/3)*3 + j/3][val])
                        return false;
                    row[i][val] = col[j][val] = box[(i/3)*3 + j/3][val] = true;
                }
            }
        }
        return true;
    }
};