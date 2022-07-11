package main

import (
	"reflect"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums:   []int{1, 2, 4, 8},
				target: 6,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -benchmem -run=none -bench. .
func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSum([]int{1, 2, 5, 8}, 13)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSum([]int{1, 2, 5, 8, 20, 50, 15, 60, 40}, 110)
	}
}
