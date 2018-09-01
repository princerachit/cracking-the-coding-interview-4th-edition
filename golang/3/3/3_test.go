package problems

import (
	"testing"
)

func Test_pop(t *testing.T) {
	tests := map[string]struct {
		ss                  *SetOfStacks
		values              []int
		expectedActiveStack int
	}{
		"1 value": {
			ss:                  NewSetOfStacks(),
			values:              []int{1, 11},
			expectedActiveStack: 0,
		},
		"11 values": {
			ss:                  NewSetOfStacks(),
			values:              []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 11},
			expectedActiveStack: 1,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			for _, i := range testCase.values {
				testCase.ss.push(i)
			}
			if testCase.expectedActiveStack != testCase.ss.activeStack {
				t.Errorf("Expected %d got %d", testCase.expectedActiveStack, testCase.ss.activeStack)
			}
		})
	}
}

func Test_push(t *testing.T) {

}
