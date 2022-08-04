package mergesort

// Merge func realisation
func Merge(a []int, b []int) []int {
	res := make([]int, len(a)+len(b), len(a)+len(b))
	f := 0
	s := 0
	for i := 0; i < len(a)+len(b); i++ {
		if f >= len(a) {
			res[i] = b[s]
			s++
		} else if s >= len(b) {
			res[i] = a[f]
			f++
		} else {
			if a[f] < b[s] {
				res[i] = a[f]
				f++
			} else {
				res[i] = b[s]
				s++
			}
		}
	}
	return res
}

// MergeSort is used to sort an array of integer
func MergeSort(input []int) []int {
	length := len(input)

	if length == 0 {
		return input
	}

	if length == 1 {
		return input
	}

	middle := length / 2
	left := input[:middle]
	right := input[middle:]

	return Merge(MergeSort(left), MergeSort(right))
}
