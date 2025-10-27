package queue

import (
	linked_list "github.com/tmdgusya/go-data-structure/linked_list"
)

type Queue struct {
	front *linked_list.LinkedListNode
	back  *linked_list.LinkedListNode
}

func NewQueue(value int) *Queue {
	node := linked_list.NewLinkedList(value)
	return &Queue{
		front: node,
		back:  node,
	}
}

func (q *Queue) Enqueue(value int) *Queue {
	node := linked_list.NewLinkedList(value)

	// 빈 큐인 경우
	if q.back == nil {
		q.front = node
		q.back = node
		return q
	}

	q.back.Next = node
	q.back = node
	return q
}

func (q *Queue) Dequeue() int {
	if q.front == nil {
		return -1
	}

	value := q.front.Value
	q.front = q.front.Next

	if q.front == nil {
		q.back = nil
	}
	return value
}
