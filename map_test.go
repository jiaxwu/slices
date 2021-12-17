package slices

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		slice  []int
		mapper func(int) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "intToString",
			args: args{
				slice: []int{2, 3, 4, 13},
				mapper: func(i int) string {
					return strconv.Itoa(i)
				},
			},
			want: []string{"2", "3", "4", "13"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.slice, tt.args.mapper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
