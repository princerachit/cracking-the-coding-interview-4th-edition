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
	fmt.Println("vim-go")
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

func Parse(r io.Reader) Problem {
	p := Problem{}

	scanner := bufio.NewScanner(r)
	step := 0

	for scanner.Scan() {
		intStrs := strings.Split(scanner.Text(), " ")
		ints := []int{}
		for _, s := range intStrs {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err) // TODO
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
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	return p
}
