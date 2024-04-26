package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "Fibonacci of 5",
			n:    5,
			want: []int{0, 1, 1, 2, 3, 5},
		},
		{
			name: "Fibonacci of 7",
			n:    7,
			want: []int{0, 1, 1, 2, 3, 5, 8, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
