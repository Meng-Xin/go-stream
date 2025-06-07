package go_stream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOfSlices(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		s := make([]int, 0)
		stream := OfSlices(s)
		assert.Equal(t, s, stream.Collect())
	})

	t.Run("non-empty slice", func(t *testing.T) {
		s := []int{1, 2, 3}
		stream := OfSlices(s)
		assert.Equal(t, s, stream.Collect())
	})
}

func TestSliceFilter(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		expected  []int
	}{
		{
			name:      "filter even numbers",
			input:     []int{1, 2, 3, 4},
			predicate: func(x int) bool { return x%2 == 0 },
			expected:  []int{2, 4},
		},
		{
			name:      "filter none",
			input:     []int{1, 2, 3},
			predicate: func(x int) bool { return true },
			expected:  []int{1, 2, 3},
		},
		{
			name:      "filter all",
			input:     []int{1, 2, 3},
			predicate: func(x int) bool { return false },
			expected:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stream := OfSlices(tt.input).Filter(tt.predicate)
			assert.Equal(t, tt.expected, stream.Collect())
		})
	}
}

func TestSliceMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		mapper   func(int) int
		expected []int
	}{
		{
			name:     "double values",
			input:    []int{1, 2, 3},
			mapper:   func(x int) int { return x * 2 },
			expected: []int{2, 4, 6},
		},
		{
			name:     "identity",
			input:    []int{1, 2, 3},
			mapper:   func(x int) int { return x },
			expected: []int{1, 2, 3},
		},
		{
			name:     "zero all",
			input:    []int{1, 2, 3},
			mapper:   func(x int) int { return 0 },
			expected: []int{0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stream := OfSlices(tt.input).Map(tt.mapper)
			assert.Equal(t, tt.expected, stream.Collect())
		})
	}
}

func TestSliceSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		compare  func(int, int) bool
		expected []int
	}{
		{
			name:     "ascending order",
			input:    []int{3, 1, 2},
			compare:  func(a, b int) bool { return a < b },
			expected: []int{1, 2, 3},
		},
		{
			name:     "descending order",
			input:    []int{3, 1, 2},
			compare:  func(a, b int) bool { return a > b },
			expected: []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stream := OfSlices(tt.input)
			result := stream.Sorted(tt.compare).Collect()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSliceChaining(t *testing.T) {
	t.Run("filter then map", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		expected := []int{4, 8}

		result := OfSlices(input).
			Filter(func(x int) bool { return x%2 == 0 }).
			Map(func(x int) int { return x * 2 }).
			Collect()

		assert.Equal(t, expected, result)
	})

	t.Run("map then filter", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		expected := []int{}

		result := OfSlices(input).
			Map(func(x int) int { return x * 2 }).
			Filter(func(x int) bool { return x%2 != 0 }).
			Collect()

		assert.Equal(t, expected, result)
	})
}

func TestEmptySliceOperations(t *testing.T) {
	t.Run("filter empty slice", func(t *testing.T) {
		result := OfSlices([]int{}).
			Filter(func(x int) bool { return true }).
			Collect()
		assert.Empty(t, result)
	})

	t.Run("map empty slice", func(t *testing.T) {
		result := OfSlices([]int{}).
			Map(func(x int) int { return x * 2 }).
			Collect()
		assert.Empty(t, result)
	})

	t.Run("sort empty slice", func(t *testing.T) {
		result := OfSlices([]int{}).
			Sorted(func(a, b int) bool { return true }).
			Collect()
		assert.Empty(t, result)
	})
}

func TestSliceDistinct(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "with duplicates",
			input:    []int{1, 2, 2, 3, 3, 3},
			expected: []int{1, 2, 2, 3, 3, 3},
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := OfSlices(tt.input).Collect()
			assert.Equal(t, tt.expected, result)
		})
	}
}
