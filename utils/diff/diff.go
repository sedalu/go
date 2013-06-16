package diff

//Diff function
func Diff(left, right []interface{}) (diffs []Delta) {
	if len(left) == 0 || len(right) == 0 {
		diffs = append(diffs, Delta{LEFT: left, RIGHT: right})
		return
	}

	var max, lmax, rmax int
	matrix := make(map[int]map[int]int)

	for li, lv := range left {
		for ri, rv := range right {
			if lv != rv {
				continue
			}

			if _, ok := matrix[li]; !ok {
				matrix[li] = make(map[int]int)
			}

			if n, ok := matrix[li-1][ri-1]; ok {
				matrix[li][ri] = n
			}

			matrix[li][ri]++

			if matrix[li][ri] <= max {
				continue
			}

			max = matrix[li][ri]
			lmax, rmax = li-max+1, ri-max+1
		}
	}

	if max == 0 {
		diffs = append(diffs, Delta{LEFT: left, RIGHT: right})
		return
	}

	for _, diff := range Diff(left[:lmax], right[:rmax]) {
		diffs = append(diffs, diff)
	}

	diffs = append(diffs, Delta{BOTH: right[rmax : rmax+max]})

	for _, diff := range Diff(left[lmax+max:], right[rmax+max:]) {
		diffs = append(diffs, diff)
	}

	return
}