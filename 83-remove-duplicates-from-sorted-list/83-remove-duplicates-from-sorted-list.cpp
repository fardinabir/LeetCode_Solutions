
class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        if(head == NULL || head->next == NULL)
            return head;
        head->next = deleteDuplicates(head->next);
        return head->val == head->next->val? head->next : head;
    }
};