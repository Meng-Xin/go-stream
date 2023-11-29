package go_stream

import (
	"fmt"
	"testing"
)

func TestMaps(t *testing.T) {
	data := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	fmt.Printf("data original address %p\n", data)
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
	result := doubledStream.Collect()
	if _, ok := result.(map[string]int); ok {
		fmt.Println("类型转换成功")
	}
	fmt.Printf("data changed address %p\n", data)
	fmt.Println("End Val", result) // 输出: map[three:6 four:8]
}
