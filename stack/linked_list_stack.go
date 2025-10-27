package stack

import (
	linked_list "github.com/tmdgusya/go-data-structure/linked_list"
)

type LinkedListStack struct {
	head *linked_list.LinkedListNode
	next *linked_list.LinkedListNode
}

func NewLinkedListStack(value int) *LinkedListStack {
	return &LinkedListStack{
		head: linked_list.NewLinkedList(value),
		next: nil,
	}
}

func (s *LinkedListStack) Push(value int) *LinkedListStack {
	new_node := linked_list.NewLinkedList(value)
	new_node.Next = s.head
	s.head = new_node

	return s
}

func (s *LinkedListStack) Pop() int {
	if s.head == nil {
		return -1
	}

	value := s.head.Value
	s.head = s.head.Next

	return value
}
