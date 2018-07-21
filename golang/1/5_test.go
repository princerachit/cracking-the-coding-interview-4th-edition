package problems

import (
	"testing"
)

func Test_replaceSpace(t *testing.T) {
	tests := []struct {
		name          string
		args          string
		expectedValue string
	}{
		{
			name:          "no spaces",
			args:          "nospaces",
			expectedValue: "nospaces",
		},
		{
			name:          "only 4 spaces",
			args:          "    ",
			expectedValue: "%20%20%20%20",
		},
		{
			name: "spaces and tabs both. one space and two tabs",
			args: "following1space followingtwotabs		",
			expectedValue: "following1space%20followingtwotabs		",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualValue := replaceSpace(tt.args)
			if actualValue != tt.expectedValue {
				t.Errorf("expected %s got %s", tt.expectedValue, actualValue)
			}
		})
	}
}
