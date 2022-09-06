
class Solution {
public:
    bool hasCycle(ListNode *head) {
        ListNode *hare = head, *tort = head;
        while(hare && hare->next){
            hare = hare->next->next;
            tort = tort->next;
            if(hare == tort)
                return true;
        }
        return false;
    }
};