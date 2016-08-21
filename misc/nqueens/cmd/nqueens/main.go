package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/sedalu/go/misc/nqueens"
)

func main() {
	var (
		n     uint
		b, p  bool
		start time.Time
		dur   time.Duration
	)

	flag.UintVar(&n, "n", 8, "the number of queens to solve for")
	flag.BoolVar(&p, "p", false, "prints the solution set")
	flag.BoolVar(&b, "b", false, "benchmarks how long it takes to find the solution")
	flag.Parse()

	if b {
		start = time.Now()
	}

	sols := nqueens.Solve(n)

	if b {
		dur = time.Since(start)
	}

	fmt.Printf("There are %v solutions to the %v queens problem.\n", len(sols), n)

	if b {
		fmt.Printf("It took %v to find them.\n", dur)
	}

	if p {
		fmt.Println("\nThey are:")
		for _, s := range sols {
			fmt.Println(nqueens.Sprint(s))
		}
	}
}
