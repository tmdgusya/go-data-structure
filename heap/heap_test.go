package heap_test

import (
	"testing"

	heap "github.com/tmdgusya/go-data-structure/heap"
)

func TestInsertSingleValue(t *testing.T) {
	h := heap.NewHeap(10)

	h.Insert(5)

	if h.Size() != 1 {
		t.Error("Insert 후 힙의 크기가 1이 아닙니다")
	}
}

func TestInsertMultipleValues(t *testing.T) {
	h := heap.NewHeap(10)

	// 여러 값 삽입
	values := []int{10, 20, 15, 30, 40}
	for _, v := range values {
		h.Insert(v)
	}

	if h.Size() != len(values) {
		t.Errorf("Insert 후 힙의 크기가 올바르지 않습니다. 기대값: %d, 실제값: %d", len(values), h.Size())
	}
}

func TestInsertMaxHeapProperty(t *testing.T) {
	h := heap.NewHeap(10)

	// 값들을 무작위 순서로 삽입
	h.Insert(5)
	h.Insert(15)
	h.Insert(10)
	h.Insert(20)
	h.Insert(8)

	// Max Heap이므로 루트는 가장 큰 값이어야 함
	if h.Peek() != 20 {
		t.Errorf("Max Heap 속성이 올바르지 않습니다. 루트 값이 최대값이 아닙니다. 기대값: 20, 실제값: %d", h.Peek())
	}
}

func TestInsertWithResize(t *testing.T) {
	initial_size := 2
	h := heap.NewHeap(initial_size)

	// 초기 크기보다 많은 값 삽입하여 Resize 트리거
	for i := 1; i <= 5; i++ {
		h.Insert(i * 10)
	}

	if h.Size() != 5 {
		t.Errorf("Resize 후 힙의 크기가 올바르지 않습니다. 기대값: 5, 실제값: %d", h.Size())
	}

	// Resize 후에도 Max Heap 속성 유지되는지 확인
	if h.Peek() != 50 {
		t.Errorf("Resize 후 Max Heap 속성이 유지되지 않았습니다. 기대값: 50, 실제값: %d", h.Peek())
	}
}

func TestInsertAscendingOrder(t *testing.T) {
	h := heap.NewHeap(10)

	// 오름차순으로 삽입
	for i := 1; i <= 10; i++ {
		h.Insert(i)
	}

	// Max Heap이므로 가장 큰 값이 루트에 있어야 함
	if h.Peek() != 10 {
		t.Errorf("오름차순 삽입 후 Max Heap 속성이 올바르지 않습니다. 기대값: 10, 실제값: %d", h.Peek())
	}
}

func TestInsertDescendingOrder(t *testing.T) {
	h := heap.NewHeap(10)

	// 내림차순으로 삽입
	for i := 10; i >= 1; i-- {
		h.Insert(i)
	}

	// Max Heap이므로 가장 큰 값이 루트에 있어야 함
	if h.Peek() != 10 {
		t.Errorf("내림차순 삽입 후 Max Heap 속성이 올바르지 않습니다. 기대값: 10, 실제값: %d", h.Peek())
	}
}

func TestInsertDuplicateValues(t *testing.T) {
	h := heap.NewHeap(10)

	// 중복된 값 삽입
	h.Insert(5)
	h.Insert(10)
	h.Insert(5)
	h.Insert(10)
	h.Insert(5)

	if h.Size() != 5 {
		t.Errorf("중복 값 삽입 후 힙의 크기가 올바르지 않습니다. 기대값: 5, 실제값: %d", h.Size())
	}

	if h.Peek() != 10 {
		t.Errorf("중복 값 삽입 후 Max Heap 속성이 올바르지 않습니다. 기대값: 10, 실제값: %d", h.Peek())
	}
}

func TestInsertLargeData(t *testing.T) {
	h := heap.NewHeap(10)

	// 많은 양의 데이터 삽입
	max_value := 0
	for i := 1; i <= 1000; i++ {
		value := i * 7 % 997 // 무작위처럼 보이는 값
		if value > max_value {
			max_value = value
		}
		h.Insert(value)
	}

	if h.Size() != 1000 {
		t.Errorf("대량 데이터 삽입 후 힙의 크기가 올바르지 않습니다. 기대값: 1000, 실제값: %d", h.Size())
	}

	if h.Peek() != max_value {
		t.Errorf("대량 데이터 삽입 후 Max Heap 속성이 올바르지 않습니다. 기대값: %d, 실제값: %d", max_value, h.Peek())
	}
}
