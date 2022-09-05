class Solution {
public:
    bool isAnagram(string s, string t) {
        if(s.size()!=t.size())
            return false;
        int flag[26]={0};
        for(int i=0;i<s.size();i++)
            flag[s[i]-'a']++;
        for(int i=0;i<t.size();i++)
            flag[t[i]-'a']--;
        for(int i=0;i<26;i++){
            if(flag[i])
                return false;
        }
        return true;
    }
};