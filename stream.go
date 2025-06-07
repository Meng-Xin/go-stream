package go_stream

// SliceStream 泛型 SliceStream
type SliceStream[T any] interface {
	// Filter 中间方法，实现对元素过滤
	Filter(predicate func(T) bool) SliceStream[T]
	// Map 中间方法，实现对元素替换
	Map(mapper func(T) T) SliceStream[T]
	// Sorted 中间方法，实现对元素排序
	Sorted(compare func(T, T) bool) SliceStream[T]
	// Collect 结束方法，收集数据
	Collect() []T
}

// MapStream 泛型 MapStream
type MapStream[C comparable, T any] interface {
	// Filter 中间方法，实现对元素过滤
	Filter(predicate func(T) bool) MapStream[C, T]
	// Map 中间方法，实现对元素替换
	Map(mapper func(T) T) MapStream[C, T]
	// Collect 结束方法，收集数据
	Collect() map[C]T
	// CollectToSlice 结束方法，收集数据，返回切片
	CollectToSlice(compare func(T, T) bool) []T
}
