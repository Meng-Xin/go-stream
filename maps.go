package go_stream

import (
	"sort"
)

type mapsStream[C comparable, T any] struct {
	data map[C]T
}

// OfMaps 转换器：传入一个基础map类型，得到一个Stream对象
func OfMaps[C comparable, T any](data map[C]T) MapStream[C, T] {
	return &mapsStream[C, T]{data: data}
}

// Filter 过滤器：根据回调方法，对map的value进行过滤
func (m *mapsStream[C, T]) Filter(predicate func(T) bool) MapStream[C, T] {
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
func (m *mapsStream[C, T]) Map(mapper func(T) T) MapStream[C, T] {
	result := make(map[C]T, len(m.data))
	for key, value := range m.data {
		result[key] = mapper(value)
	}
	m.data = result
	return m
}

// Collect 收集器：将Stream流对象通过interface返回
func (m *mapsStream[C, T]) Collect() map[C]T {
	return m.data
}

// CollectToSlice 排序器：对map的val进行排序
func (m *mapsStream[C, T]) CollectToSlice(compare func(T, T) bool) []T {
	// 修正：正确初始化keys切片的长度
	keys := make([]C, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return compare(m.data[keys[i]], m.data[keys[j]])
	})
	result := make([]T, 0, len(m.data))
	// 修正：使用append添加元素而不是直接索引赋值
	for _, key := range keys {
		result = append(result, m.data[key])
	}
	return result
}
