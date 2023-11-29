package go_stream

type Stream[T any] interface {
	// Filter 中间方法，实现对元素过滤
	Filter(predicate func(T) bool) Stream[T]
	// Map 中间方法，实现对元素替换
	Map(mapper func(T) T) Stream[T]
	// Sorted 中间方法，实现对元素排序
	Sorted(compare func(T, T) bool) Stream[T]
	// Collect 结束方法，收集数据，调用端获取需要断言类型。
	Collect() interface{}
}
