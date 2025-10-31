package heap

type Prioritized interface {
	GetPriority() int
}

type Heap[T Prioritized] struct {
	array      []*T
	array_size int
	last_index int
}

func NewHeap[T Prioritized](size int) *Heap[T] {
	return &Heap[T]{
		array:      make([]*T, size),
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

	h.moveUp(curr)
}

func (h *Heap[T]) moveUp(curr int) {
	parent := (curr - 1) / 2
	for curr > 0 && (*h.array[parent]).GetPriority() < (*h.array[curr]).GetPriority() {
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

	// 힙에 원소가 남아있으면 heapify-down
	if h.last_index > 0 {
		h.moveDown(0)
	}

	return result
}

func (h *Heap[T]) Update(idx int, value T) bool {
	if idx >= h.last_index || idx < 0 {
		return false
	}

	old := (*h.array[idx]).GetPriority()
	new := value.GetPriority()

	// 값을 실제로 업데이트
	h.array[idx] = &value

	if old < new {
		// 기존보다 우선순위가 높아졌으므로 위로 이동
		h.moveUp(idx)
	} else if old > new {
		// 기존보다 우선순위가 낮아졌으므로 아래로 이동
		h.moveDown(idx)
	}
	// old == new 이면 아무것도 안 해도 됨

	return true
}

func (h *Heap[T]) moveDown(curr int) {
	for {
		left := curr*2 + 1
		right := curr*2 + 2

		if left >= h.last_index {
			break
		}

		larger := left

		if right < h.last_index && (*h.array[right]).GetPriority() > (*h.array[left]).GetPriority() {
			larger = right
		}

		if (*h.array[curr]).GetPriority() >= (*h.array[larger]).GetPriority() {
			// 현재가 자식보다 크면 복구 할 필요 없음
			break
		}

		tmp := h.array[curr]
		h.array[curr] = h.array[larger]
		h.array[larger] = tmp

		curr = larger
	}
}
