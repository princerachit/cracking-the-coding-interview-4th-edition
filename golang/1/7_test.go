package problems

import (
	"reflect"
	"testing"
)

func Test_destroyRowsCols(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name           string
		args           args
		expectedMatrix args
	}{
		{
			name: "matrix with no zero elts",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			expectedMatrix: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
		},
		{
			name: "matrix with one zero elts",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 0},
					{7, 8, 9},
				},
			},
			expectedMatrix: args{
				matrix: [][]int{
					{1, 2, 0},
					{0, 0, 0},
					{7, 8, 0},
				},
			},
		},
		{
			name: "matrix with two zero elts in one row",
			args: args{
				matrix: [][]int{
					{1, 0, 0},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			expectedMatrix: args{
				matrix: [][]int{
					{0, 0, 0},
					{4, 0, 0},
					{7, 0, 0},
				},
			},
		},
		{
			name: "matrix with two zero elts in one col",
			args: args{
				matrix: [][]int{
					{1, 2, 0},
					{4, 5, 0},
					{7, 8, 9},
				},
			},
			expectedMatrix: args{
				matrix: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{7, 8, 0},
				},
			},
		},
		{
			name: "matrix with many elts",
			args: args{
				matrix: [][]int{
					{1, 0, 0},
					{4, 5, 0},
					{0, 8, 9},
				},
			},
			expectedMatrix: args{
				matrix: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
		},
		{
			name: "non square matrix",
			args: args{
				matrix: [][]int{
					{1, 0, 3},
					{4, 5, 6},
				},
			},
			expectedMatrix: args{
				matrix: [][]int{
					{0, 0, 0},
					{4, 0, 6},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			destroyRowsCols(tt.args.matrix)
			if !reflect.DeepEqual(tt.args.matrix, tt.expectedMatrix.matrix) {
				t.Errorf("expected %v got %v", tt.expectedMatrix.matrix, tt.args.matrix)
			}
		})
	}
}
