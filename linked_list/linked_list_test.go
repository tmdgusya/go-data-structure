package linked_list

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	initial_value := 42
	n := NewLinkedList(initial_value)

	if n.Value != initial_value {
		t.Error("LinkedList 가 제값으로 초기화되지 않았습니다.")
	}
}

func TestAppend(t *testing.T) {
	append_value := 3
	head := NewLinkedList(1)
	head.Append(append_value)

	appended := head.Next

	if appended.Value != append_value {
		print(appended.Value)
		t.Error("제대로 append 되지 않았습니다")
	}
}
