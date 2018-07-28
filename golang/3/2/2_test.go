package problems

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_pop(t *testing.T) {
	s1 := make(stack, 100)
	s2 := make(stack, 100)
	s2[0] = element{data: 1, smallestAfter: 1}
	s2[1] = element{data: 4, smallestAfter: 1}
	tests := []struct {
		name    string
		s       stack
		top     int
		want    *element
		wantTop int
		wantErr error
	}{
		{
			name:    "empty stack",
			s:       s1,
			top:     -1,
			want:    nil,
			wantTop: -1,
			wantErr: fmt.Errorf("underflow"),
		},
		{
			name:    "non empty stack",
			s:       s2,
			top:     1,
			want:    &s2[1],
			wantTop: 0,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, newTop, err := pop(tt.s, tt.top)
			if tt.wantErr != nil {
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("expected error when top = %v", tt.top)
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pop() got = %v, want %v", got, tt.want)
			}
			if newTop != tt.wantTop {
				t.Errorf("pop() got1 = %v, want %v", newTop, tt.wantTop)
			}
		})
	}
}

func Test_push(t *testing.T) {
	s1 := make(stack, 100)
	s2 := make(stack, 100)
	s2[0] = element{data: 1, smallestAfter: 1}
	s2[1] = element{data: 4, smallestAfter: 1}
	tests := []struct {
		s    stack
		e    element
		top  int
		name string
		want int
	}{
		{
			s:    s1,
			e:    element{data: 123},
			top:  -1,
			name: "empty stack",
			want: 0,
		},
		{
			s:    s2,
			e:    element{data: 123},
			top:  1,
			name: "non empty stack",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := push(tt.s, tt.e, tt.top)
			if got != tt.want {
				t.Errorf("push() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.s[got].data, tt.e.data) {
				t.Errorf("push() got = %v, want %v", tt.s[got].data, tt.e.data)
			}
		})
	}
}

func Test_minElement(t *testing.T) {
	s1 := make(stack, 100)
	s2 := make(stack, 100)
	s2[0] = element{data: 1, smallestAfter: 1}
	s2[1] = element{data: 4, smallestAfter: 1}
	_, err := minElement(s1, -1)
	if err == nil {
		t.Errorf("expected error")
	}

	push(s2, element{data: -3}, 1)
	min, _ := minElement(s2, 2)
	if min != -3 {
		t.Errorf("minElement() got = %v, want = %v", min, -3)
	}
	pop(s2, 2)
	min, _ = minElement(s2, 1)
	if min != 1 {
		t.Errorf("minElement() got = %v, want = %v", min, 1)
	}
}
