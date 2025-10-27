package stack_test

import (
	"testing"

	stack "github.com/tmdgusya/go-data-structure/stack"
)

func TestResizing(t *testing.T) {
	s := stack.NewStack()
	initial_size := 10

	if s.Length() != int8(initial_size) {
		t.Error("스택이 초기값보다 더 크게 초기화 되었습니다.")
	}

	s.Resize()

	if s.Length() != int8(initial_size)*2 {
		t.Error(("스택이 초기값에서 두배로 증가하지 않았습니다"))
	}
}

func TestPush(t *testing.T) {
	size := 2
	s := stack.NewStackWith(int8(size))

	for i := 1; i <= size; i++ {
		s.Push(int8(i))
	}

	if s.Length() != int8(size) {
		t.Error("배열이 리사이즈 되었습니다")
	}

	s.Push(1)

	if s.Length() != int8(size)*2 {
		t.Error("배열이 리사이즈가 되지 않았습니다")
	}
}

func TestPop(t *testing.T) {
	s := stack.NewStack()

	// Push 3개의 값
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Pop으로 LIFO 순서 확인
	if s.Pop() != 3 {
		t.Error("Pop이 마지막에 Push한 값을 반환하지 않았습니다")
	}

	if s.Pop() != 2 {
		t.Error("Pop이 올바른 순서로 값을 반환하지 않았습니다")
	}

	if s.Pop() != 1 {
		t.Error("Pop이 올바른 순서로 값을 반환하지 않았습니다")
	}

	// 빈 스택에서 Pop하면 -1 반환
	if s.Pop() != -1 {
		t.Error("빈 스택에서 Pop이 -1을 반환하지 않았습니다")
	}
}
