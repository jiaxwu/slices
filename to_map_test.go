package slices

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestToMap(t *testing.T) {
	type args struct {
		slice     []int
		keyFunc   func(int) int
		valueFunc func(int) any
	}
	tests := []struct {
		name string
		args args
		want map[int]any
	}{
		{
			name: "number1",
			args: args{
				slice: []int{2, 3, 4, 13},
				keyFunc: func(i int) int {
					return i
				},
				valueFunc: func(i int) any {
					return i
				},
			},
			want: map[int]any{
				2: 2, 3: 3, 4: 4, 13: 13,
			},
		},
		{
			name: "number2",
			args: args{
				slice: []int{2, 3, 4, 13},
				keyFunc: func(i int) int {
					return i
				},
				valueFunc: func(i int) any {
					return fmt.Sprint(i)
				},
			},
			want: map[int]any{
				2: "2", 3: "3", 4: "4", 13: "13",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToMap(tt.args.slice, tt.args.keyFunc, tt.args.valueFunc)
			gotJson, _ := json.Marshal(&got)
			wantJson, _ := json.Marshal(&tt.want)
			if string(gotJson) != string(wantJson) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
