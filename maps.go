package go_stream

type mapsStream[C comparable, T any] struct {
	data map[C]T
}

func OfMaps[C comparable, T any](data map[C]T) Stream[T] {
	return &mapsStream[C, T]{data: data}
}

func (m *mapsStream[C, T]) Filter(predicate func(T) bool) Stream[T] {
	result := make(map[C]T)
	for key, value := range m.data {
		if predicate(value) {
			result[key] = value
		}
	}
	return &mapsStream[C, T]{data: result}
}

func (m *mapsStream[C, T]) Map(mapper func(T) T) Stream[T] {
	result := make(map[C]T)
	for key, value := range m.data {
		result[key] = mapper(value)
	}
	return &mapsStream[C, T]{data: result}
}

func (m *mapsStream[C, T]) Sorted(compare func(T, T) bool) Stream[T] {
	//TODO implement me
	panic("implement me")
}
func (m *mapsStream[C, T]) Collect() interface{} {
	return m.data
}
