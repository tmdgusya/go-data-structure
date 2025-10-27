package linked_list

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

func NewLinkedList(value int) *LinkedListNode {
	return &LinkedListNode{
		Value: value,
		Next:  nil,
	}
}

func (n *LinkedListNode) Append(value int) *LinkedListNode {
	curr := n
	for curr.Next != nil {
		curr = curr.Next
	}

	newNode := &LinkedListNode{Value: value, Next: nil}
	curr.Next = newNode
	return n
}
