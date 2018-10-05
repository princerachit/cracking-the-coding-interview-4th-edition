package problems

import (
	"testing"
)

func TestByteLandianGoldCoinProblem(t *testing.T) {
	tests := []struct {
		name          string
		number        int64
		expectedValue uint64
	}{
		{
			name:          "Execution Sequence for 0",
			number:        0,
			expectedValue: 0,
		},
		{
			name:          "Execution Sequence for 12",
			number:        12,
			expectedValue: 13,
		},
		{
			name:          "Execution Sequence for -1",
			number:        -1,
			expectedValue: 0,
		},
		{
			name:          "Execution Sequence for 20",
			number:        20,
			expectedValue: 21,
		},
		{
			name:          "Execution Sequence for 73",
			number:        73,
			expectedValue: 87,
		},
		{
			name:          "Execution Sequence for 2562",
			number:        2562,
			expectedValue: 3769,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualValue := ByteLandianGoldCoinProblem(tt.number)
			if tt.expectedValue != actualValue {
				t.Errorf("expected %d got %d", tt.expectedValue, actualValue)
			}
		})
	}
}
