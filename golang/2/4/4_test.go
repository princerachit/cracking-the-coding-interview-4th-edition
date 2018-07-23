package problems

import (
	"testing"
)

func Test_calcSumFromLinkedList(t *testing.T) {
	tests := []struct {
		name  string
		list1 *node
		list2 *node
		want  *node
	}{
		{
			name: "even sized lists with no carry at the end",
			list1: &node{
				data: 1,
				next: &node{
					data: 2,
					next: nil,
				},
			},
			list2: &node{
				data: 1,
				next: &node{
					data: 2,
					next: nil,
				},
			},
			want: &node{
				data: 2,
				next: &node{
					data: 4,
					next: nil,
				},
			},
		},
		{
			name: "uneven sized lists with no carry at the end",
			list1: &node{
				data: 1,
				next: &node{
					data: 2,
					next: nil,
				},
			},
			list2: &node{
				data: 1,
				next: nil,
			},
			want: &node{
				data: 2,
				next: &node{
					data: 2,
					next: nil,
				},
			},
		},
		{
			name: "uneven sized lists with carry at the end",
			list1: &node{
				data: 1,
				next: &node{
					data: 2,
					next: nil,
				},
			},
			list2: &node{
				data: 1,
				next: &node{
					data: 9,
					next: nil,
				},
			},
			want: &node{
				data: 2,
				next: &node{
					data: 1,
					next: &node{
						data: 1,
						next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sumList := calcSumFromLinkedList(tt.list1, tt.list2)
			for sumList.next != nil {
				if sumList.data != tt.want.data {
					t.Errorf("calcSumFromLinkedList() = %v, want %v", sumList, tt.want)
				}
				sumList = sumList.next
				tt.want = tt.want.next
			}
		})
	}
}
