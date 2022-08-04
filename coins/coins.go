package coins

func Piles(n int) int {
	partitions := make([]int, n+1, n+1)
	partitions[0] = 1
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			partitions[j] = partitions[j] + partitions[j-i]
		}
	}
	return partitions[n]
}
