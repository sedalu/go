// Package nqueens provides a data structure, Solution, and a function, Solve, for finding the set of valid solutions to the N Queens problem. Please refer to https://en.m.wikipedia.org/wiki/Eight_queens_puzzle.
package nqueens

// Sprint returns a string formated as a grid with the location of each queen marked by "Q".
func Sprint(s []int) string {
	if len(s) == 0 {
		return "++\n++\n"
	}

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

// Solve returns only the valid solutions for placing n queens on an n x n board, such that no queen can travel to the location of any other queen in a single move, given stardard chess movement rules.
// Each permutation is calculated recursively. At each stage, the solution is recursively validated and invalid solutions are immediately thrown out.
func Solve(n uint) [][]int {
	opts := make([]int, n, n)
	for i := range opts {
		opts[i] = i
	}

	return solve(make([]int, 0, n), opts)
}

// solve recursively calculates each permutation and validates it before adding to the set of solutions
func solve(root []int, opts []int) (solutions [][]int) {
	// no more options to choose from
	if len(opts) == 0 {
		if len(root) == 0 {
			return
		}

		return append(solutions, root)
	}

	for i, opt := range opts {
		// if adding opt would result in an invalid solution, continue to the next opt
		if !validate(opt, len(root), root) {
			continue
		}

		// copy the root solution
		s := make([]int, len(root)+1, cap(root))
		copy(s, root)

		s[len(root)] = opt

		// remove opt from the options used to solve the next position
		o := append(make([]int, 0, cap(opts)-1), opts[:i]...)
		o = append(o, opts[i+1:]...)

		// append the results of solving the next positioin to the solution set
		solutions = append(solutions, solve(s, o)...)
	}

	return
}

// validate recursively checks if slope between the points (x, y) and (s[len(s)-1], len(s)-1) is +/-1. If the slope is +/-1, then false is returned. Otherwise, the length of s is shortened by 1 until len(s) is 0.
func validate(x, y int, s []int) bool {
	if len(s) == 0 {
		return true
	}

	i := len(s) - 1

	// check for horizontal or verticle placements
	if s[i] == x || i == y {
		return false
	}

	if isDiagonal(s[i], i, x, y) {
		return false
	}

	return validate(x, y, s[:i])
}

func isDiagonal(x1, y1, x2, y2 int) bool {
	dx, dy := x2-x1, y2-y1

	// infinite slopes
	if dx == 0 {
		return false
	}

	// slopes < 1
	if dy%dx != 0 {
		return false
	}

	// slopes >= 1
	switch dy / dx {
	case -1, 1:
		return true
	default:
		return false
	}
}
