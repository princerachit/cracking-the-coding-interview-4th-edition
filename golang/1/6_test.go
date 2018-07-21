package problems

import (
	"reflect"
	"testing"
)

func Test_rotateMatrix(t *testing.T) {

	tests := []struct {
		name string
		args [][]int64
		want [][]int64
	}{
		{
			name: "4x4",
			args: [][]int64{
				{11, 12, 13, 14},
				{15, 16, 17, 18},
				{19, 20, 21, 22},
				{23, 24, 25, 26},
			},
			want: [][]int64{
				{23, 19, 15, 11},
				{24, 20, 16, 12},
				{25, 21, 17, 13},
				{26, 22, 18, 14},
			},
		},
		{
			name: "3x3",
			args: [][]int64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int64{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "2x2",
			args: [][]int64{
				{1, 2},
				{3, 4},
			},
			want: [][]int64{
				{3, 1},
				{4, 2},
			},
		}, {
			name: "1",
			args: [][]int64{
				{1},
			},
			want: [][]int64{
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateMatrix(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
