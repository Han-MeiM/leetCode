#include <iostream>
#include "Solution.h"

void print(ListNode *head)        //输出逆置后的单链表
{
	ListNode *p = head;
    std::cout << p->val;
	std::cout << std::endl;
}


int main()
{
    ListNode node1(1);
    ListNode node2(2);
    ListNode node3(3);
    ListNode node4(4);
    ListNode node5(5);
    ListNode node6(6);
    ListNode node7(7);
    ListNode node8(8);

    node1.next = &node2;
    node2.next = &node6;

    node3.next = &node4;
    node4.next = &node5;
    node5.next = &node6;

    node6.next = &node7;
    node7.next = &node8;

    Solution s;
    ListNode *res = s.getIntersectionNode(&node1, &node3);
    print(res);
    return 0;
}
