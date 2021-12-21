package slices

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "number1",
			args: args{
				slice: []int{1, 2, 3, 2, 1, 2, 3, 5, 6, 3, 4, 2, 1, 2},
			},
			want: []int{1, 2, 3, 5, 6, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicate(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
