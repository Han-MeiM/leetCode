/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
const reverseList = function (head) {
    if (head.next == null) return head;
    const res = reverseList(head.next);
    head.next.next = head;
    head.next = null;
    console.log(res);
    console.log(head);
    return res
};

function ListNode(val) {
    this.val = val;
    this.next = null;
}

const root = new ListNode(1);
root.next = new ListNode(2);
root.next.next = new ListNode(3);
root.next.next.next = new ListNode(4);
root.next.next.next.next = new ListNode(5);

console.log(root);
console.log(reverseList(root));