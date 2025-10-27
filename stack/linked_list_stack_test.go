package stack_test

import (
	"testing"

	stack "github.com/tmdgusya/go-data-structure/stack"
)

func TestLinkedListStackPush(t *testing.T) {
	s := stack.NewLinkedListStack(1)

	// Push 여러 값들 - 링크드 리스트는 동적이므로 크기 제한 없음
	for i := 2; i <= 100; i++ {
		s.Push(i)
	}

	// 링크드 리스트는 동적이므로 많은 양의 데이터도 처리 가능
	// Pop으로 가장 마지막 값 확인
	if s.Pop() != 100 {
		t.Error("Push가 올바르게 작동하지 않았습니다")
	}
}

func TestLinkedListStackPop(t *testing.T) {
	s := stack.NewLinkedListStack(1)

	// Push 2개의 값 더 추가
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

func TestLinkedListStackDynamicGrowth(t *testing.T) {
	// 링크드 리스트 스택은 크기 제한이 없으므로 계속 추가 가능
	s := stack.NewLinkedListStack(1)

	// 많은 양의 데이터 추가
	for i := 2; i <= 1000; i++ {
		s.Push(i)
	}

	// LIFO 순서로 확인
	if s.Pop() != 1000 {
		t.Error("동적 확장 후 Pop이 올바르게 작동하지 않았습니다")
	}

	if s.Pop() != 999 {
		t.Error("동적 확장 후 Pop이 올바르게 작동하지 않았습니다")
	}
}

func TestLinkedListStackEmptyPop(t *testing.T) {
	s := stack.NewLinkedListStack(5)

	// 모든 값을 Pop
	s.Pop()

	// 빈 스택에서 여러 번 Pop
	if s.Pop() != -1 {
		t.Error("빈 스택에서 Pop이 -1을 반환하지 않았습니다")
	}

	if s.Pop() != -1 {
		t.Error("빈 스택에서 연속 Pop이 -1을 반환하지 않았습니다")
	}
}

func TestLinkedListStackPushFromEmpty(t *testing.T) {
	// 빈 스택 생성 (head가 nil인 상태)
	s := &stack.LinkedListStack{}

	// 빈 스택에서 Push
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// LIFO 순서 확인
	if s.Pop() != 3 {
		t.Error("빈 스택에서 Push 후 Pop이 올바르게 작동하지 않았습니다")
	}

	if s.Pop() != 2 {
		t.Error("Pop이 올바른 순서로 값을 반환하지 않았습니다")
	}

	if s.Pop() != 1 {
		t.Error("Pop이 올바른 순서로 값을 반환하지 않았습니다")
	}

	if s.Pop() != -1 {
		t.Error("모든 값을 Pop한 후 -1을 반환하지 않았습니다")
	}
}
