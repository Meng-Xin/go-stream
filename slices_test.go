package go_stream

import (
	"reflect"
	"testing"
)

func Test_slicesStream_Filter(t *testing.T) {
	type args[T any] struct {
		predicate func(T) bool
	}
	type testCase[T any] struct {
		name string
		s    Stream[T]
		args args[T]

		want []T
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "[PASS] int集合 (Filter)偶数场景",
			s:    OfSlices([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}),
			args: args[int]{func(i int) bool {
				return (i % 2) == 0
			}},

			want: []int{0, 2, 4, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := []int{}
			if tt.s.Filter(tt.args.predicate).Collect(&got); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slicesStream_Sorted(t *testing.T) {
	type args[T any] struct {
		compare func(T, T) bool
	}
	type testCase[T any] struct {
		name string
		s    Stream[T]
		args args[T]

		want []T
	}
	// Int 类型测试
	testsInt := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "[PASS] int数组 升序排序",
			s:    OfSlices([]int{1, 5, 6, 3, 9, 7}),
			args: args[int]{compare: func(x int, y int) bool {
				return x-y < 0
			}},
			want: []int{1, 3, 5, 6, 7, 9},
		},
		{
			name: "[PASS] int数组 降序排序",
			s:    OfSlices([]int{1, 5, 6, 3, 9, 7}),
			args: args[int]{compare: func(x int, y int) bool {
				return x-y > 0
			}},
			want: []int{9, 7, 6, 5, 3, 1},
		},
	}
	for _, tt := range testsInt {
		t.Run(tt.name, func(t *testing.T) {
			got := []int{}
			if tt.s.Sorted(tt.args.compare).Collect(&got); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sorted() = %v, want %v", got, tt.want)
			}
		})
	}

	// String 类型测试
	testsString := []testCase[string]{
		// TODO: Add test cases.
		{
			name: "[PASS] String数组 升序排序",
			s:    OfSlices([]string{"a", "c", "d", "b"}),
			args: args[string]{compare: func(x string, y string) bool {
				return x < y
			}},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "[PASS] String数组 降序排序",
			s:    OfSlices([]string{"a", "c", "d", "b"}),
			args: args[string]{compare: func(x string, y string) bool {
				return x > y
			}},
			want: []string{"d", "c", "b", "a"},
		},
	}
	for _, tt := range testsString {
		t.Run(tt.name, func(t *testing.T) {
			got := []string{}
			if tt.s.Sorted(tt.args.compare).Collect(&got); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slicesStream_Collect(t *testing.T) {
	type args[T any] struct {
		predicate func(T) bool
	}
	type testCase[T any] struct {
		name string
		s    Stream[T]
		args args[T]

		want []T
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "[PASS] int集合 (Collect)正确的赋值场景！",
			s:    OfSlices([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}),
			args: args[int]{func(i int) bool {
				return i > 5
			}},

			// 正确赋值可以收集到数据
			want: []int{6, 7, 8},
		},
		{
			name: "[FAIL] int集合 (Collect)错误的赋值场景！",
			s:    OfSlices([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}),
			args: args[int]{func(i int) bool {
				return i > 5
			}},

			// 错误赋值导致最终无法收集数据
			want: []int{},
		},
	}

	// PASS 正确接受参数
	passResult := []int{}
	t.Run(tests[0].name, func(t *testing.T) {
		if tests[0].s.Filter(tests[0].args.predicate).Collect(&passResult); !reflect.DeepEqual(passResult, tests[0].want) {
			t.Errorf("Filter() = %v, want %v", passResult, tests[0].want)
		}
	})

	// FAIL 错误接收参数
	failResult := []int{}
	t.Run(tests[1].name, func(t *testing.T) {
		if tests[1].s.Filter(tests[1].args.predicate).Collect(failResult); !reflect.DeepEqual(failResult, tests[1].want) {
			t.Errorf("Filter() = %v, want %v", failResult, tests[1].want)
		}
	})
}
