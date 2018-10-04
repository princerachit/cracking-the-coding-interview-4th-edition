package main

import (
	"reflect"
	"strings"
	"testing"
)

var sample string = "5 5\n1 2 3 4 5\n5 4 3 2 1\n1 1 5\n2 1 5\n1 2 4\n2 2 4\n1 3 5"

func TestParser(t *testing.T) {
	p := Parse(strings.NewReader(sample))
	if len(p.A) != 5 {
		t.Errorf("Length of A wrong. Expected 5, observed %d", len(p.A))
	}
	if len(p.B) != 5 {
		t.Errorf("Length of B wrong. Expected 5, observed %d", len(p.B))
	}

	a := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(p.A, a) {
		t.Errorf("Content of A wrong. Expected %v, observed %v", a, p.A)
	}
	b := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(p.B, b) {
		t.Errorf("Content of B wrong. Expected %v, observed %v", b, p.B)
	}

	if len(p.Q) != 5 {
		t.Errorf("Length of Q wrong. Expected 5, observed %d", len(p.Q))
		return
	}

	expected := []Query{
		Query{1, 1, 5},
		Query{2, 1, 5},
		Query{1, 2, 4},
		Query{2, 2, 4},
		Query{1, 3, 5},
	}

	for i, e := range expected {
		if p.Q[i].Type != e.Type || p.Q[i].L != e.L || p.Q[i].R != e.R {
			t.Errorf("Query at %d doesn't match. Expected %v, observed %v.", i, e, p.Q[i])
		}
	}
}
