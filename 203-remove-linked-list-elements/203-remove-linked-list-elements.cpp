class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        ListNode *cur = head, *tail = head;
        while(cur){
            if(cur->val == val){
                if(cur == head)
                    head = head->next;
                else{
                    tail->next = cur->next;
                }
            }
            else
                tail = cur;
            cur = cur->next;
        }
        return head;
    }
};