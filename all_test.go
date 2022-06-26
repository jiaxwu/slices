package slices

import "testing"

func TestAll(t *testing.T) {
	type args struct {
		slice     []int
		condition func(item int, index int, slice []int) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "number1",
			args: args{
				slice: []int{2, 3, 4, 13},
				condition: func(item int, index int, slice []int) bool {
					return item != 2
				},
			},
			want: false,
		},
		{
			name: "number2",
			args: args{
				slice: []int{2, 3, 4, 13},
				condition: func(item int, index int, slice []int) bool {
					return item != 5
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.args.slice, tt.args.condition); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
