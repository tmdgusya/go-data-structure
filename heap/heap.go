package heap

import (
	"math"

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
		return *h.array[1]
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
	if h.last_index == h.array_size-1 {
		h.Resize()
	}

	h.last_index++ // 마지막에 삽입하는 거니까
	h.array[h.last_index] = &value

	curr := h.last_index
	parent := int(math.Floor(float64(curr) / 2))
	for parent >= 1 && *h.array[parent] < *h.array[curr] {
		temp := h.array[parent] // 주소
		h.array[parent] = h.array[curr]
		h.array[curr] = temp

		// 위치 업데이트
		curr = parent
		parent = int(math.Floor(float64(curr) / 2))
	}
}
