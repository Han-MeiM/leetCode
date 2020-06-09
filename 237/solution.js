const deleteNode = (node) => {
    node.val = node.next.val;
    node.next = node.next.next;
};

function ListNode(val) {
    this.val = val;
    this.next = null;
}

const node = new ListNode(4);
node.next = new ListNode(5);
node.next.next = new ListNode(1);
node.next.next.next = new ListNode(9);

deleteNode(node.next.next)

console.log(node)