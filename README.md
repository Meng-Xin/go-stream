# go-stream
> "Go-stream 是一个受 Java Collections 框架中 Stream 接口操作启发的学习项目，支持泛型
> (Go-stream is a learning project inspired by the Stream interface operations in the Java Collections framework, with support for generics!)"

## Prerequisites
> GoVersion >= 1.21.0
## QuickStart

**SlicesDemo**
```go
package main

import (
	"fmt"
	go_stream "github.com/Meng-Xin/go-stream"
)

func main() {
	sourceData := []int{5, 2, 0, 1, 3, 1, 4, 1024}
	// 过滤偶数 filter even numbers
	go_stream.OfSlices(sourceData).Filter(func(i int) bool {
		return i%2 == 0
	}).Map(func(i int) int {
		// 基数*2 Base*2
		return 2 * i
	}).Sorted(func(i int, i2 int) bool {
		// 升序排序 Sort in ascending order
		return i-i2 < 0
	}).Collect(&sourceData)

	// EndVal: [0,2,4,8,2048]
	fmt.Println("EndVal:", sourceData)
}
```
**MapsDemo**
```go
package main

import (
	"fmt"
	go_stream "github.com/Meng-Xin/go-stream"
)

func main() {
	sourceData := map[int]string{
		1: "Tom",
		2: "Jerry",
		3: "Dog",
	}

	go_stream.OfMaps(sourceData).Filter(func(s string) bool {
		// 过滤掉Tom和Jerry
		return s != "Tom" && s != "Jerry"
	}).Map(func(s string) string {
		// 修改Dog->Speike
		if s == "Dog" {
			return "Speike"
		} else {
			return s
		}
	}).Collect(&sourceData)

	// map[3:Speike]
	fmt.Println(sourceData)
}
```