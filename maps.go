package go_stream

import (
	"fmt"
	"sort"
)

type mapsStream[C comparable, T any] struct {
	data map[C]T
}

// OfMaps 转换器：传入一个基础map类型，得到一个Stream对象
func OfMaps[C comparable, T any](data map[C]T) Stream[T] {
	return &mapsStream[C, T]{data: data}
}

// Filter 过滤器：根据回调方法，对map的value进行过滤
func (m *mapsStream[C, T]) Filter(predicate func(T) bool) Stream[T] {
	result := make(map[C]T, len(m.data))
	for key, value := range m.data {
		if predicate(value) {
			result[key] = value
		}
	}
	m.data = result
	return m
}

// Map 映射器：根据毁掉方法，对map的value进行映射修改
func (m *mapsStream[C, T]) Map(mapper func(T) T) Stream[T] {
	result := make(map[C]T, len(m.data))
	for key, value := range m.data {
		result[key] = mapper(value)
	}
	m.data = result
	return m
}

// Sorted 排序器：对map的val进行排序
func (m *mapsStream[C, T]) Sorted(compare func(T, T) bool) Stream[T] {
	keys := make([]C, 0, len(m.data))
	for key, _ := range m.data {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return compare(m.data[keys[i]], m.data[keys[j]])
	})
	sortedStream := &mapsStream[C, T]{
		data: make(map[C]T),
	}
	for _, key := range keys {
		sortedStream.data[key] = m.data[key]
	}
	// 原有data已经舍弃，直接清空。
	clear(m.data)
	return sortedStream
}

// Collect 收集器：将Stream流对象通过interface返回
func (m *mapsStream[C, T]) Collect(accept interface{}) {
	if result, ok := accept.(*map[C]T); ok {
		*result = m.data
	} else {
		fmt.Println("Collect 接收参数需要一个指向map[C]T的地址！")
	}
}
