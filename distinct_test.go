package slices

import (
	"reflect"
	"testing"
)

func TestDistinct(t *testing.T) {
	type args struct {
		slice     []int
		condition func(item1, item2 int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "number1",
			args: args{
				slice: []int{1, 3, 2, 1, 2, 3, 2, 1, 4},
				condition: func(item1, item2 int) bool {
					return item1 == item2
				},
			},
			want: []int{1, 3, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distinct(tt.args.slice, tt.args.condition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
