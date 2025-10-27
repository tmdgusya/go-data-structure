package stack

type Stack struct {
	array_size int8
	top        int8
	values     []int8
}

func NewStack() *Stack {
	return &Stack{
		array_size: 10,
		top:        -1,
		values:     make([]int8, 10),
	}
}

func NewStackWith(size int8) *Stack {
	return &Stack{
		array_size: size,
		top:        -1,
		values:     make([]int8, size),
	}
}

func (s *Stack) Length() int8 {
	return s.array_size
}

func (s *Stack) Resize() {
	new_array_size := s.array_size * 2
	new_values := make([]int8, new_array_size)

	copy(new_values, s.values)

	s.values = new_values
	s.array_size = new_array_size
}

func (s *Stack) Push(value int8) {
	if s.top+1 >= s.array_size {
		s.Resize()
	}
	s.top++
	s.values[s.top] = value
}

func (s *Stack) Pop() int8 {
	result := -1
	if s.top > -1 {
		result = int(s.values[s.top])
		s.top--
	}
	return int8(result)
}
