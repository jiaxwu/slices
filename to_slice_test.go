package slices

import (
	"encoding/json"
	"testing"
)

func TestToSlice(t *testing.T) {
	type args struct {
		dict      map[string]int
		transform func(string, int) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "number1",
			args: args{
				dict: map[string]int{"1": 1, "2": 2, "3": 3, "4": 4},
				transform: func(key string, value int) int {
					return value
				},
			},
			want: []int{
				1, 2, 3, 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToSlice(tt.args.dict, tt.args.transform)
			gotJson, _ := json.Marshal(&got)
			wantJson, _ := json.Marshal(&tt.want)
			if string(gotJson) != string(wantJson) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
