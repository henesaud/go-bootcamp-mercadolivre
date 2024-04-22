package cal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "test case 1",
			args: []int{5, 3, 1, 2, 4},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "test case 2",
			args: []int{3, 0, 88, -54, 33},
			want: []int{-54, 0, 3, 33, 88},
		},
	}
	for _, tt := range tests {
		got := Order(tt.args)
		assert.Equal(t, tt.want, got, tt.want)
	}
}
