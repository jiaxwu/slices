package slices

import (
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	type args struct {
		slice  []int
		reduce func(cur int, item int) int
		init   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "number1",
			args: args{
				slice: []int{3, 4, 5, 3},
				reduce: func(cur int, item int) int {
					return cur + item
				},
				init: 0,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.slice, tt.args.reduce, tt.args.init); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
