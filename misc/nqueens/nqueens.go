package main

import "fmt"

func main() {
	sols := NQueens(4)
	fmt.Println("Solutions:", len(sols))
	for _, p := range sols {
		fmt.Println(p)
	}
}

// Solution stores the set of column offsets.
type Solution []int

func (s Solution) String() string {
	var div, str string

	for i := 0; i < len(s); i++ {
		div += "+---"
	}

	div += "+\n"

	for _, v := range s {
		str += div

		for i := 0; i < len(s); i++ {
			switch i {
			case v:
				str += "| Q "
			default:
				str += "|   "
			}
		}

		str += "|\n"
	}

	str += div
	return str
}

// NQueens takes the number of queens to solve for and returns the set of solutions.
func NQueens(n uint) []Solution {
	// list options (0...n)
	opts := make([]int, n, n)
	for i := range opts {
		opts[i] = i
	}

	// find and return possible solutions
	return solve(make(Solution, 0, n), opts)
}

func solve(root Solution, opts []int) (solutions []Solution) {
	// no more options to choose from, return the root solution
	if len(opts) == 0 {
		solutions = append(solutions, root)
	}

	// for each opt, validate the solution and then
	for i, opt := range opts {
		// if adding opt would result in an invalid solution, continue to the next opt
		if !validate(root, len(root), opt) {
			continue
		}

		// copy the root solution
		solution := make(Solution, len(root)+1, cap(root))
		copy(solution, root[:len(root)])

		// add opt to the solution
		solution[len(root)] = opt

		// remove opt from the options used to solve the next position
		o := append(make([]int, 0, cap(opts)-1), opts[:i]...)
		o = append(o, opts[i+1:]...)

		// append the results of solving the next positioin to the solution set
		solutions = append(solutions, solve(solution, o)...)
	}

	// return the solution set
	return
}

// validate recursively checks if adding opt into row i is valid within solution Solution.
func validate(solution Solution, i, opt int) bool {
	if len(solution) < 1 {
		return true
	}

	j := len(solution) - 1

	switch (i - j) / (opt - solution[j]) {
	case -1, 1:
		return false
	default:
		return validate(solution[:j], i, opt)
	}
}
