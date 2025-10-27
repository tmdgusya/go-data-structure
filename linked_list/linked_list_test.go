package linked_list

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	initial_value := 42
	n := NewLinkedList(initial_value)

	if n.value != initial_value {
		t.Error("LinkedList 가 제값으로 초기화되지 않았습니다.")
	}
}

func TestAppend(t *testing.T) {
	append_value := 3
	head := NewLinkedList(1)
	head.Append(append_value)

	appended := head.next

	if appended.value != append_value {
		print(appended.value)
		t.Error("제대로 append 되지 않았습니다")
	}
}
