package heap

import (
	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	array      []*T
	array_size int
	last_index int
}

func NewHeap(size int) *Heap[int] {
	return &Heap[int]{
		array:      make([]*int, size),
		array_size: size,
		last_index: 0,
	}
}

func (h *Heap[T]) Size() int {
	return h.last_index
}

func (h *Heap[T]) Peek() T {
	if h.last_index > 0 {
		return *h.array[0]
	}
	var zero T
	return zero
}

func (h *Heap[T]) Resize() {
	new_arr := make([]*T, h.array_size*2)
	copy(new_arr, h.array)
	h.array = new_arr
	h.array_size = h.array_size * 2
}

func (h *Heap[T]) Insert(value T) {
	if h.last_index == h.array_size {
		h.Resize()
	}

	h.array[h.last_index] = &value
	curr := h.last_index
	h.last_index++ // 다음 삽입 위치로 이동

	parent := (curr - 1) / 2
	for curr > 0 && *h.array[parent] < *h.array[curr] {
		temp := h.array[parent] // 주소
		h.array[parent] = h.array[curr]
		h.array[curr] = temp

		// 위치 업데이트
		curr = parent
		parent = (curr - 1) / 2
	}
}

func (h *Heap[T]) Remove() T {
	var zero T
	if h.last_index == 0 {
		return zero
	}

	result := *h.array[0]
	// root node 랑 last node swap
	h.array[0] = h.array[h.last_index-1]
	h.array[h.last_index-1] = nil
	h.last_index--

	// 자식들 중에 큰걸 찾아서 바꿔줘야 함
	curr := 0
	for {
		left := curr*2 + 1
		right := curr*2 + 2

		if left >= h.last_index {
			break
		}

		larger := left

		if right < h.last_index && *h.array[right] > *h.array[left] {
			larger = right
		}

		if *h.array[curr] >= *h.array[larger] {
			// 현재가 자식보다 크면 복구 할 필요 없음
			break
		}

		tmp := h.array[curr]
		h.array[curr] = h.array[larger]
		h.array[larger] = tmp

		curr = larger
	}

	return result
}
