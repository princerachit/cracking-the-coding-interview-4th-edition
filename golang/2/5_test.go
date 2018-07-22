package problems

import "testing"

func Test_findLoop(t *testing.T) {

	l1 := &linkedList{
		data: 'A',
		next: nil,
	}
	l2 := &linkedList{
		data: 'B',
		next: nil,
	}
	l3 := &linkedList{
		data: 'C',
		next: nil,
	}
	l4 := &linkedList{
		data: 'D',
		next: nil,
	}

	l1.next = l2
	l2.next = l3
	l3.next = l4
	l4.next = l4

	m1 := &linkedList{
		data: 'A',
		next: nil,
	}
	m2 := &linkedList{
		data: 'B',
		next: nil,
	}
	m3 := &linkedList{
		data: 'C',
		next: nil,
	}
	m4 := &linkedList{
		data: 'D',
		next: nil,
	}

	m1.next = m2
	m2.next = m3
	m3.next = m4
	m4.next = m2

	tests := []struct {
		name string
		list *linkedList
		want rune
	}{
		{
			name: "element loops to itself",
			list: l1,
			want: 'D',
		},
		{
			name: "element loops to other element",
			list: m1,
			want: 'B',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLoop(tt.list); got != tt.want {
				t.Errorf("findLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
