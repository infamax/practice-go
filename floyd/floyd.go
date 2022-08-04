package floyd

// Triangle makes a Floyd's triangle matrix with rows count.
func Triangle(rows int) [][]int {
	res := make([][]int, 0, rows)
	n := ((1 + rows) * rows) / 2
	tmp := make([]int, 0, 1)
	length := 1
	last := 1
	plus := 2
	for i := 1; i <= n; i++ {
		tmp = append(tmp, i)
		if i == last {
			res = append(res, tmp)
			length++
			last += plus
			plus++
			tmp = make([]int, 0, length)
		}
	}
	return res
}
