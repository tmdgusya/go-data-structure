package linked_list

type LinkedListNode struct {
	value int
	next  *LinkedListNode
}

func NewLinkedList(value int) *LinkedListNode {
	return &LinkedListNode{
		value: value,
		next:  nil,
	}
}

func (n *LinkedListNode) Append(value int) *LinkedListNode {
	curr := n
	for curr.next != nil {
		curr = curr.next
	}

	newNode := &LinkedListNode{value: value, next: nil}
	curr.next = newNode
	return n
}
