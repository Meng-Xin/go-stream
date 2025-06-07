package go_stream

import (
	"sort"
)

// slicesStream 切片泛型Stream
type slicesStream[T any] struct {
	data []T
}

// OfSlices 转换器：传入一个基础切片类型，得到一个Stream对象，注意：这里进行深拷贝，不会影响原始数据。
func OfSlices[T any](data []T) SliceStream[T] {
	stream := slicesStream[T]{data: data}
	copy(data, stream.data)
	return &stream
}

// Sorted 排序器：通过回调函数指定排序顺序。
func (s *slicesStream[T]) Sorted(compare func(T, T) bool) SliceStream[T] {
	if compare == nil {
		return s
	}
	if len(s.data) <= 1 {
		return s
	}
	sort.Slice(s.data, func(i, j int) bool {
		return compare(s.data[i], s.data[j])
	})
	return s
}

// Filter 过滤器：使用回调函数获取自定义的函数返回值。
func (s *slicesStream[T]) Filter(predicate func(T) bool) SliceStream[T] {
	if predicate == nil {
		return s
	}
	result := make([]T, 0, len(s.data)) // 预分配容量
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
	if mapper == nil {
		return s
	}
	result := make([]T, len(s.data)) // 直接分配确定长度
	for i, val := range s.data {
		result[i] = mapper(val)
	}
	s.data = result
	return s
}

// ForEach 遍历集合中的每个元素
func (s *slicesStream[T]) ForEach(consumer func(T)) SliceStream[T] {
	if consumer == nil {
		return s
	}
	for _, val := range s.data {
		consumer(val)
	}
	return s
}

// Collect 返回该集合的一个具体类型
func (s *slicesStream[T]) Collect() []T {
	return s.data
}
