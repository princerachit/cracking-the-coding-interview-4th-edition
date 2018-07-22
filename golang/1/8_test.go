package problems

import "testing"

func Test_isRotation(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "failure test case anagram with non rotation",
			args: args{
				str1: "golang",
				str2: "oglang",
			},
			want: false,
		},
		{
			name: "positive test case",
			args: args{
				str1: "waterbottle",
				str2: "erbottlewat",
			},
			want: true,
		},
		{
			name: "negative test case non-anagram",
			args: args{
				str1: "waterbottle",
				str2: "waterbottlX",
			},
			want: false,
		},
		{
			name: "negative test case different length strings",
			args: args{
				str1: "waterbottle",
				str2: "water",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRotation(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("isRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
