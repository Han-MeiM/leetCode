/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        if (!headA || !headB) return NULL;
        ListNode *item1 = headA, *item2 = headB;
        while(item1 != NULL || item2 != NULL) {
            if (item1 == item2) return item1;
            item1 = item1->next;
            item2 = item2->next;
            
            if (item1 == NULL && item2 == NULL) return NULL;
            if (item1 == NULL) item1 = headB;
            if (item2 == NULL) item2 = headA;
        }
        return NULL;
    }
};