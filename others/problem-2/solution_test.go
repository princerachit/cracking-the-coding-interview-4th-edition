package problem2

import "testing"

func Test_knapSack(t *testing.T) {
	type args struct {
		w   int
		wt  []int
		val []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simpliest case",
			args: args{
				w:   1,
				wt:  []int{1},
				val: []int{1},
			},
			want: 1,
		},
		{
			name: "knapsack is too small",
			args: args{
				w:   1,
				wt:  []int{2, 3, 4, 5, 6, 7, 8, 9},
				val: []int{1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: 0,
		},
		{
			name: "choose one among many",
			args: args{
				w:   1,
				wt:  []int{1, 1, 1, 1, 1, 1, 1, 1},
				val: []int{2, 3, 4, 5, 9, 8, 7, 6},
			},
			want: 9,
		},
		{
			name: "an ordinary case #1",
			args: args{
				w:   40,
				wt:  []int{10, 20, 30},
				val: []int{90, 100, 120},
			},
			want: 210,
		},
		{
			name: "an ordinary case #2",
			args: args{
				w:   40,
				wt:  []int{10, 20, 30, 5},
				val: []int{90, 105, 120, 20},
			},
			want: 215,
		},
		{
			name: "an ordinary case #3",
			args: args{
				w:   17,
				wt:  []int{65, 30, 2, 34, 62, 3, 32, 79, 97, 96},
				val: []int{18, 5, 71, 61, 6, 97, 9, 34, 99, 68},
			},
			want: 168,
		},
		{
			name: "an ordinary case #4",
			args: args{
				w:   70,
				wt:  []int{40, 26, 8, 66, 86, 86, 24, 54, 2, 20},
				val: []int{1, 53, 54, 77, 57, 57, 49, 17, 77, 23},
			},
			want: 233,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapSack(tt.args.w, tt.args.wt, tt.args.val); got != tt.want {
				t.Errorf("knapSack() = %v, want %v", got, tt.want)
			}
		})
	}
}
