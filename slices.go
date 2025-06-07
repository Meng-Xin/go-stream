package go_stream

import (
	"sort"
)

// slicesStream 切片泛型Stream
type slicesStream[T any] struct {
	data []T
}

func OfSlices[T any](data []T) SliceStream[T] {
	return &slicesStream[T]{data: data}
}

// Sorted 排序器：通过回调函数指定排序顺序。
func (s *slicesStream[T]) Sorted(compare func(T, T) bool) SliceStream[T] {
	sort.Slice(s.data, func(i, j int) bool {
		return compare(s.data[i], s.data[j])
	})
	return s
}

// Filter 过滤器：使用回调函数获取自定义的函数返回值。
func (s *slicesStream[T]) Filter(predicate func(T) bool) SliceStream[T] {
	result := make([]T, 0)
	for _, val := range s.data {
		if predicate(val) {
			result = append(result, val)
		}
	}
	s.data = result
	return s
}

// Map 对于集合中的每个元素进行映射操作的方法
func (s *slicesStream[T]) Map(mapper func(T) T) SliceStream[T] {
	var result []T
	for _, val := range s.data {
		result = append(result, mapper(val))
	}
	s.data = result
	return s
}

// Collect 返回该集合的一个具体类型
func (s *slicesStream[T]) Collect() []T {
	return s.data
}
