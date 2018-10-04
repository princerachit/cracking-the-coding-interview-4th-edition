package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	p, err := Parse(os.Stdin)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}
	for _, q := range p.Q {
		fmt.Println(q.Solve(p))
	}
}

type Problem struct {
	A []int
	B []int
	Q []Query
}

type Query struct {
	Type int
	L    int
	R    int
}

func Parse(r io.Reader) (*Problem, error) {
	p := Problem{}

	scanner := bufio.NewScanner(r)
	step := 0

	for scanner.Scan() {
		intStrs := strings.Split(scanner.Text(), " ")
		ints := []int{}
		for _, s := range intStrs {
			if len(s) == 0 {
				continue
			}
			i, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			ints = append(ints, i)
		}

		switch step {
		case 0:
			p.A = make([]int, ints[0])
			p.B = make([]int, ints[0])
		case 1:
			copy(p.A, ints)
		case 2:
			copy(p.B, ints)
		default:
			p.Q = append(p.Q, Query{ints[0], ints[1], ints[2]})
		}

		step = step + 1
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &p, nil
}

func (q Query) Solve(p *Problem) int {
	// NOTE: Our indexes are 0 based, the problem is 1 based, so we make sure to adjust.
	v := 0
	if q.Type == 1 {
		// 1 L R : Print the value of AL + BL+1 + AL+2 + BL+3 + ... upto Rth term.
		for i := 0; i+q.L <= q.R; i++ {
			if i%2 == 0 {
				v = v + p.A[i+q.L-1]
			} else {
				v = v + p.B[i+q.L-1]
			}
		}
	} else {
		// 2 L R : Print the value of BL + AL+1 + BL+2 + AL+3 + ... upto Rth term.
		for i := 0; i+q.L <= q.R; i++ {
			if i%2 == 0 {
				v = v + p.B[i+q.L-1]
			} else {
				v = v + p.A[i+q.L-1]
			}
		}
	}
	return v
}
