package go_stream

// slicesStream 切片泛型Stream
type slicesStream[T any] struct {
	data []T
}

func (s *slicesStream[T]) Sorted(compare func(T, T) bool) Stream[T] {
	//TODO implement me
	panic("implement me")
}

// Filter 过滤器：使用回调函数获取自定义的函数返回值。
func (s *slicesStream[T]) Filter(callFunc func(T) bool) Stream[T] {
	var result []T
	for _, val := range s.data {
		if callFunc(val) {
			result = append(result, val)
		}
	}
	return &slicesStream[T]{data: result}
}

// Map 对于集合中的每个元素进行映射操作的方法
func (s *slicesStream[T]) Map(mapper func(T) T) Stream[T] {
	var result []T
	for _, val := range s.data {
		result = append(result, mapper(val))
	}
	return &slicesStream[T]{data: result}
}

// Collect 返回该集合的一个具体类型
func (s *slicesStream[T]) Collect() interface{} {
	return s.data
}

func OfSlices[T any](data []T) Stream[T] {
	return &slicesStream[T]{data: data}
}
