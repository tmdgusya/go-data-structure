package distance

import (
	"testing"
)

func TestDist(t *testing.T) {
	tests := []struct {
		name     string
		a        float32
		b        float32
		expected float32
	}{
		{
			name:     "positive difference",
			a:        10.0,
			b:        5.0,
			expected: 5.0,
		},
		{
			name:     "negative difference",
			a:        5.0,
			b:        10.0,
			expected: 5.0,
		},
		{
			name:     "zero difference",
			a:        7.5,
			b:        7.5,
			expected: 0.0,
		},
		{
			name:     "decimal values",
			a:        3.7,
			b:        1.2,
			expected: 2.5,
		},
		{
			name:     "negative numbers",
			a:        -5.0,
			b:        -10.0,
			expected: 5.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dist(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("dist(%v, %v) = %v, expected %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestLinearScanClosetNeighbor(t *testing.T) {
	tests := []struct {
		name     string
		arr      []float32
		target   float32
		expected float32
	}{
		{
			name:     "empty array",
			arr:      []float32{},
			target:   5.0,
			expected: -1,
		},
		{
			name:     "single element",
			arr:      []float32{10.0},
			target:   5.0,
			expected: 10.0,
		},
		{
			name:     "target is in array",
			arr:      []float32{1.0, 5.0, 10.0},
			target:   5.0,
			expected: 5.0,
		},
		{
			name:     "target between elements",
			arr:      []float32{1.0, 5.0, 10.0},
			target:   6.0,
			expected: 5.0,
		},
		{
			name:     "target less than all elements",
			arr:      []float32{10.0, 20.0, 30.0},
			target:   5.0,
			expected: 10.0,
		},
		{
			name:     "target greater than all elements",
			arr:      []float32{1.0, 2.0, 3.0},
			target:   10.0,
			expected: 3.0,
		},
		{
			name:     "negative numbers",
			arr:      []float32{-10.0, -5.0, 0.0, 5.0, 10.0},
			target:   -7.0,
			expected: -5.0,
		},
		{
			name:     "decimal values",
			arr:      []float32{1.5, 2.7, 3.9, 5.1},
			target:   3.0,
			expected: 2.7,
		},
		{
			name:     "first element is closest",
			arr:      []float32{5.0, 10.0, 15.0},
			target:   4.0,
			expected: 5.0,
		},
		{
			name:     "last element is closest",
			arr:      []float32{1.0, 5.0, 10.0},
			target:   11.0,
			expected: 10.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LinearScanClosetNeighbor(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("LinearScanClosetNeighbor(%v, %v) = %v, expected %v", tt.arr, tt.target, result, tt.expected)
			}
		})
	}
}
