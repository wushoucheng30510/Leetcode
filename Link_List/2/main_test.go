package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				l1: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9, Next: nil}}},
				l2: &ListNode{4, &ListNode{9, nil}},
			},
			want: &ListNode{7, &ListNode{4, &ListNode{0, &ListNode{1, nil}}}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addTwoNumbers(&ListNode{5, &ListNode{8, nil}}, &ListNode{6, &ListNode{6, nil}})
	}
}
