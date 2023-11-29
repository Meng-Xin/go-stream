package go_stream

import "testing"

func TestMaps(t *testing.T) {
	data := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	t.Logf("data original address %p\n", data)
	stream := OfMaps(data)

	// 过滤出值大于2的键值对
	filteredStream := stream.Filter(func(value int) bool {
		return value > 2
	})

	// 将值加倍
	doubledStream := filteredStream.Map(func(value int) int {
		return value * 2
	})

	// 收集结果
	doubledStream.Collect(&data)
	t.Logf("data changed address %p\n", data)
	t.Log("End Val", data) // 输出: map[three:6 four:8]
}

func TestMapsStream_Sorted(t *testing.T) {
	data := map[string]int{"a": 1, "d": 4, "c": 3, "b": 2}
	stream := OfMaps(data)
	t.Logf("data original address %p\n", data)
	stream.Sorted(func(x int, y int) bool {
		return (x - y) > 0
	})
	stream.Collect(&data)
	t.Logf("data changed address %p\n", data)
	t.Log("End Val", data)
}
