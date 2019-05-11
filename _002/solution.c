#include <stdio.h>
#include <stdlib.h>

struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *addTwoNumbers(struct ListNode *l1, struct ListNode *l2)
{

    int k = 0;
    struct ListNode *l1Head = l1;
    struct ListNode *l2Head = l2;

    struct ListNode *temp = l1;

    while (l1Head)
    {

        int a1 = l1Head ? l1Head->val : 0;
        int a2 = l2Head ? l2Head->val : 0;
        int a = a1 + a2 + k;

        if (a >= 10)
        {
            k = 1;
            a = a % 10;
        }
        else
        {
            k = 0;
        }

        l1Head->val = a;

        temp = l1Head;
        l1Head = temp->next;

        if (l2Head)
        {
            l2Head = l2Head->next;
        }

        if (l1Head == NULL && l2Head != NULL)
        {
            (*temp).next = l2Head;
            l1Head = temp->next;
            l2Head = NULL;
        }
    }
    //! 溢出，则需要新建一个节点
    if (k > 0)
    {
        struct ListNode *lastHead = (struct ListNode *)malloc(sizeof(struct ListNode));
        lastHead->next = NULL;
        lastHead->val = k;
        temp->next = lastHead;
    }

    return l1;
}

int main(void)
{
    struct ListNode l11 = {3};
    struct ListNode l12 = {4, &l11};
    struct ListNode l13 = {2, &l12};

    struct ListNode l21 = {4};
    struct ListNode l22 = {6, &l11};
    struct ListNode l23 = {5, &l12};

    struct ListNode *result = addTwoNumbers(&l13, &l23);

    printf(
        "number: %d%d%d",
        result->val,
        result->next->val,
        result->next->next->val);
}