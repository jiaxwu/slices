package slices

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		slice  []int
		filter func(int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "number1",
			args: args{
				slice: []int{18, 19, 6, 3, 43, 1, 32},
				filter: func(i int) bool {
					return i > 18
				},
			},
			want: []int{19, 43, 32},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.slice, tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterWithIndex(t *testing.T) {
	type args struct {
		slice  []int
		filter func(int, int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "number1",
			args: args{
				slice: []int{18, 19, 6, 3, 43, 1, 32},
				filter: func(i int, index int) bool {
					return index > 2
				},
			},
			want: []int{3, 43, 1, 32},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterWithIndex(tt.args.slice, tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
