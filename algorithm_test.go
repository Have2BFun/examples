package examples

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "n=0 返回空切片",
			n:    0,
			want: []int{},
		},
		{
			name: "n为负数返回空切片",
			n:    -1,
			want: []int{},
		},
		{
			name: "n=1",
			n:    1,
			want: []int{1},
		},
		{
			name: "n=2",
			n:    2,
			want: []int{1, 1},
		},
		{
			name: "n=5",
			n:    5,
			want: []int{1, 1, 2, 3, 5},
		},
		{
			name: "n=10",
			n:    10,
			want: []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Fibonacci(tt.n)
			if len(got) != len(tt.want) {
				t.Fatalf("Fibonacci(%d) 长度 = %d, 期望 %d", tt.n, len(got), len(tt.want))
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Fibonacci(%d)[%d] = %d, 期望 %d", tt.n, i, got[i], tt.want[i])
				}
			}
		})
	}
}
