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

func TestRemoveSingleElement(t *testing.T) {
	h := heap.NewHeap(10)
	h.Insert(42)

	removed := h.Remove()
	if removed != 42 {
		t.Errorf("Remove 반환값이 올바르지 않습니다. 기대값: 42, 실제값: %d", removed)
	}

	if h.Size() != 0 {
		t.Errorf("Remove 후 힙의 크기가 올바르지 않습니다. 기대값: 0, 실제값: %d", h.Size())
	}
}

func TestRemoveFromEmptyHeap(t *testing.T) {
	h := heap.NewHeap(10)

	removed := h.Remove()
	if removed != 0 {
		t.Errorf("빈 힙에서 Remove 시 zero value를 반환해야 합니다. 기대값: 0, 실제값: %d", removed)
	}
}

func TestRemoveMaxHeapProperty(t *testing.T) {
	h := heap.NewHeap(10)

	// [50, 30, 40, 10, 20] 힙 생성
	values := []int{50, 30, 40, 10, 20}
	for _, v := range values {
		h.Insert(v)
	}

	// 첫 번째 Remove: 50이 나와야 함
	removed := h.Remove()
	if removed != 50 {
		t.Errorf("첫 번째 Remove가 최대값을 반환하지 않았습니다. 기대값: 50, 실제값: %d", removed)
	}

	// Remove 후에도 Max Heap 속성 유지되는지 확인
	if h.Peek() != 40 {
		t.Errorf("Remove 후 Max Heap 속성이 유지되지 않았습니다. 기대값: 40, 실제값: %d", h.Peek())
	}
}

func TestRemoveMultipleElements(t *testing.T) {
	h := heap.NewHeap(10)

	// 1부터 10까지 삽입
	for i := 1; i <= 10; i++ {
		h.Insert(i)
	}

	// 내림차순으로 제거되는지 확인
	expected := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i, exp := range expected {
		removed := h.Remove()
		if removed != exp {
			t.Errorf("Remove %d번째: 기대값 %d, 실제값 %d", i+1, exp, removed)
		}
	}

	// 모두 제거 후 힙이 비어있는지 확인
	if h.Size() != 0 {
		t.Errorf("모든 원소 제거 후 힙의 크기가 0이 아닙니다. 실제값: %d", h.Size())
	}
}

func TestRemoveAndInsertMixed(t *testing.T) {
	h := heap.NewHeap(5)

	// 삽입과 삭제를 섞어서 수행
	h.Insert(10)
	h.Insert(20)
	h.Insert(15)

	if h.Remove() != 20 {
		t.Error("첫 번째 Remove가 20을 반환하지 않았습니다")
	}

	h.Insert(25)
	h.Insert(5)

	if h.Peek() != 25 {
		t.Errorf("삽입/삭제 후 최대값이 올바르지 않습니다. 기대값: 25, 실제값: %d", h.Peek())
	}
}

func TestRemoveWithDuplicates(t *testing.T) {
	h := heap.NewHeap(10)

	// 중복 값 삽입
	values := []int{5, 10, 5, 10, 5}
	for _, v := range values {
		h.Insert(v)
	}

	// 첫 두 번은 10이 나와야 함
	if h.Remove() != 10 {
		t.Error("중복 값 제거 - 첫 번째 10이 나오지 않았습니다")
	}
	if h.Remove() != 10 {
		t.Error("중복 값 제거 - 두 번째 10이 나오지 않았습니다")
	}

	// 나머지는 5
	for i := 0; i < 3; i++ {
		if h.Remove() != 5 {
			t.Errorf("중복 값 제거 - %d번째 5가 나오지 않았습니다", i+1)
		}
	}
}

func TestRemoveAllElements(t *testing.T) {
	h := heap.NewHeap(5)

	values := []int{15, 10, 20, 5, 25}
	for _, v := range values {
		h.Insert(v)
	}

	// 모든 원소를 내림차순으로 제거
	prev := h.Remove() // 25
	for h.Size() > 0 {
		current := h.Remove()
		if current > prev {
			t.Errorf("Remove 순서가 올바르지 않습니다. 이전값: %d, 현재값: %d", prev, current)
		}
		prev = current
	}

	// 빈 힙에서 한 번 더 Remove
	if h.Remove() != 0 {
		t.Error("빈 힙에서 Remove 시 zero value를 반환해야 합니다")
	}
}
