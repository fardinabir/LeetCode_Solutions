class Solution {
public:
    bool validUtf8(vector<int>& data) {
        int toCheck = 0;
        int mask1 = 1<<7, mask2 = 1<<6;
        for(int i=0; i<data.size(); i++){
            if(toCheck == 0){
                int maskt = 1<<7;
                while(maskt&data[i])
                    toCheck++, maskt = maskt>>1;
                if(toCheck == 0)
                    continue;
                if(toCheck>4 || toCheck==1)
                    return false;
            }
            else{
                if(!(data[i]&mask1) || data[i]&mask2)
                    return false;
            }
            toCheck--;
        }
        return toCheck == 0;
    }
};