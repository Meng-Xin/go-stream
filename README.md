# go-stream

A lightweight Go library for functional-style operations on slices and maps, inspired by Java Streams.

## Features

- **Chainable operations**: Filter, map, sort, and more in a fluent style
- **Type-safe**: Uses Go generics for type safety
- **Lazy evaluation**: Operations are applied sequentially
- **Immutable**: Each operation returns a new stream

## Installation

```bash
go get github.com/Meng-Xin/go-stream
```

## Usage

### Slice Operations

#### Creating a Stream from Slice

```go
package main

import go_stream "github.com/Meng-Xin/go-stream"

func main(){
	s := []int{1, 2, 3}
	stream := go_stream.OfSlices(s)
}
```

#### Filter

```go
result := go_stream.OfSlices([]int{1, 2, 3, 4}).
    Filter(func(x int) bool { return x%2 == 0 }).
    Collect() // [2, 4]
```

#### Map

```go
result := go_stream.OfSlices([]int{1, 2, 3}).
    Map(func(x int) int { return x * 2 }).
    Collect() // [2, 4, 6]
```

#### Sort

```go
result := go_stream.OfSlices([]int{3, 1, 2}).
    Sorted(func(a, b int) bool { return a < b }).
    Collect() // [1, 2, 3]
```

#### Chaining Operations

```go
result := go_stream.OfSlices([]int{1, 2, 3, 4}).
    Filter(func(x int) bool { return x%2 == 0 }).
    Map(func(x int) int { return x * 2 }).
    Collect() // [4, 8]
```

### Map Operations

#### Creating a Stream from Map

```go
m := map[string]int{"a": 1, "b": 2}
stream := go_stream.OfMaps(m)
```

#### Filter

```go
result := go_stream.OfMaps(map[string]int{"a": 1, "b": 2, "c": 3}).
    Filter(func(x int) bool { return x%2 != 0 }).
    Collect() // map[string]int{"a": 1, "c": 3}
```

#### Map

```go
result := go_stream.OfMaps(map[string]int{"a": 1, "b": 2}).
    Map(func(x int) int { return x * 2 }).
    Collect() // map[string]int{"a": 2, "b": 4}
```

#### Collect to Sorted Slice

```go
result := go_stream.OfMaps(map[string]int{"a": 3, "b": 1, "c": 2}).
    CollectToSlice(func(a, b int) bool { return a < b }) // [1, 2, 3]
```

## Performance

Operations are optimized for performance:
- Minimal allocations
- Pre-sized collections where possible
- Efficient iteration

## Examples

See the test files for more complete examples:
- `slices_test.go`
- `maps_test.go`

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
