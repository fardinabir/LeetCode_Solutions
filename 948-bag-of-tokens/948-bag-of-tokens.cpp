class Solution {
public:
    int bagOfTokensScore(vector<int>& tokens, int power) {
        sort(tokens.begin(),tokens.end());
        int score=0, i=0, j=tokens.size()-1;
        while(i<=j){
            if(power>=tokens[i]){
                power -= tokens[i++], score++;
            }
            else if(score>0 && j>i+1){
                    power += tokens[j--], score--;
            }
            else
                break;
        }
        return score;
    }
};