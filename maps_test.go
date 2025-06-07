package go_stream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOfMaps(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		m := make(map[string]int)
		stream := OfMaps(m)
		assert.Equal(t, m, stream.Collect())
	})

	t.Run("non-empty map", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		stream := OfMaps(m)
		assert.Equal(t, m, stream.Collect())
	})
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		input     map[string]int
		predicate func(int) bool
		expected  map[string]int
	}{
		{
			name:      "filter even numbers",
			input:     map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			predicate: func(x int) bool { return x%2 == 0 },
			expected:  map[string]int{"b": 2, "d": 4},
		},
		{
			name:      "filter none",
			input:     map[string]int{"a": 1, "b": 2},
			predicate: func(x int) bool { return true },
			expected:  map[string]int{"a": 1, "b": 2},
		},
		{
			name:      "filter all",
			input:     map[string]int{"a": 1, "b": 2},
			predicate: func(x int) bool { return false },
			expected:  map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stream := OfMaps(tt.input).Filter(tt.predicate)
			assert.Equal(t, tt.expected, stream.Collect())
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		mapper   func(int) int
		expected map[string]int
	}{
		{
			name:     "double values",
			input:    map[string]int{"a": 1, "b": 2},
			mapper:   func(x int) int { return x * 2 },
			expected: map[string]int{"a": 2, "b": 4},
		},
		{
			name:     "identity",
			input:    map[string]int{"a": 1, "b": 2},
			mapper:   func(x int) int { return x },
			expected: map[string]int{"a": 1, "b": 2},
		},
		{
			name:     "zero all",
			input:    map[string]int{"a": 1, "b": 2},
			mapper:   func(x int) int { return 0 },
			expected: map[string]int{"a": 0, "b": 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stream := OfMaps(tt.input).Map(tt.mapper)
			assert.Equal(t, tt.expected, stream.Collect())
		})
	}
}

func TestCollectToSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		compare  func(int, int) bool
		expected []int
	}{
		{
			name:     "ascending order",
			input:    map[string]int{"a": 3, "b": 1, "c": 2},
			compare:  func(a, b int) bool { return a < b },
			expected: []int{1, 2, 3},
		},
		{
			name:     "descending order",
			input:    map[string]int{"a": 3, "b": 1, "c": 2},
			compare:  func(a, b int) bool { return a > b },
			expected: []int{3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := OfMaps(tt.input).CollectToSlice(tt.compare)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestChaining(t *testing.T) {
	t.Run("filter then map", func(t *testing.T) {
		input := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		expected := map[string]int{"b": 4, "d": 8}

		result := OfMaps(input).
			Filter(func(x int) bool { return x%2 == 0 }).
			Map(func(x int) int { return x * 2 }).
			Collect()

		assert.Equal(t, expected, result)
	})

	t.Run("map then filter", func(t *testing.T) {
		input := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		expected := map[string]int{"a": 2, "c": 6}

		result := OfMaps(input).
			Map(func(x int) int { return x * 2 }).
			Filter(func(x int) bool { return x%2 != 0 }).
			Collect()

		assert.Equal(t, expected, result)
	})
}

func TestEmptyMapOperations(t *testing.T) {
	t.Run("filter empty map", func(t *testing.T) {
		result := OfMaps(map[string]int{}).
			Filter(func(x int) bool { return true }).
			Collect()
		assert.Empty(t, result)
	})

	t.Run("map empty map", func(t *testing.T) {
		result := OfMaps(map[string]int{}).
			Map(func(x int) int { return x * 2 }).
			Collect()
		assert.Empty(t, result)
	})

	t.Run("collect to slice empty map", func(t *testing.T) {
		result := OfMaps(map[string]int{}).
			CollectToSlice(func(a, b int) bool { return true })
		assert.Empty(t, result)
	})
}
