package main

import (
	"fmt"

	"github.com/sedalu/go/misc/nqueens"
)

func main() {
	sols := nqueens.Solve(4)
	fmt.Println("Solutions:", len(sols))
	for _, p := range sols {
		fmt.Println(p)
	}
}
