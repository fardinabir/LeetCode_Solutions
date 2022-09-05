class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        int freq[26] = {0};
        for(int i=0; i<magazine.size(); i++){
            freq[magazine[i]-97]++;
        }
        for(int i=0; i<ransomNote.size(); i++){
            freq[ransomNote[i]-97]--;
            if(freq[ransomNote[i]-97]<0)
                return false;
        }
        return true;
    }
};