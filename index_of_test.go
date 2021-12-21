package slices

import "testing"

func TestIndexOf(t *testing.T) {
	type args struct {
		slice []int
		item  int
		start int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "number1",
			args: args{
				slice: []int{3, 4, 5, 3, 2},
				item:  4,
				start: 0,
			},
			want: 1,
		},
		{
			name: "number2",
			args: args{
				slice: []int{3, 4, 5, 3, 2},
				item:  4,
				start: -7,
			},
			want: 1,
		},
		{
			name: "number3",
			args: args{
				slice: []int{3, 4, 5, 3, 2},
				item:  4,
				start: -2,
			},
			want: -1,
		},
		{
			name: "number3",
			args: args{
				slice: []int{3, 4, 5, 3, 2},
				item:  3,
				start: -2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOf(tt.args.slice, tt.args.item, tt.args.start); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
