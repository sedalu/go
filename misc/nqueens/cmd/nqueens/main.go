package main

import (
	"flag"
	"fmt"

	"github.com/sedalu/go/misc/nqueens"
)

func main() {
	pn := flag.Uint("n", 8, "the number of queens to solve for")
	flag.Parse()

	var n uint = 8

	if pn != nil {
		n = *pn
	}

	sols := nqueens.Solve(n)
	fmt.Println("Solutions:", len(sols))
	//	for _, p := range sols {
	//		fmt.Println(p)
	//	}
}
