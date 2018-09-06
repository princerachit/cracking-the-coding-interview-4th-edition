package problems

import (
	"strings"
	"testing"
)

func Test_printExecutionSequence(t *testing.T) {
	tests := []struct {
		name          string
		number        int64
		expectedValue string
	}{
		{
			name:          "Execution Sequence for 0",
			number:        0,
			expectedValue: "",
		},
		{
			name:          "Execution Sequence for 1",
			number:        1,
			expectedValue: "",
		},
		{
			name:          "Execution Sequence for -11",
			number:        -11,
			expectedValue: "LLRL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualValue := printExecutionSequence(tt.number)
			if strings.Compare(tt.expectedValue, actualValue) != 0 {
				t.Errorf("expected %s got %s", tt.expectedValue, actualValue)
			}
		})
	}
}
