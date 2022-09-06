
class Solution {
public:
    bool hasCycle(ListNode *head) {
        ListNode *hare, *tort;
        hare = tort = head;
        while(hare && hare->next && tort){
            hare = hare->next->next;
            tort = tort->next;
            if(hare == tort)
                return true;
        }
        return false;
    }
};