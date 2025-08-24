func kthFactor(n int, k int) int {
    small := []int{}
	big := []int{}

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			small = append(small, i)
			if i != n/i {
				big = append([]int{n / i}, big...)
			}
		}
	}

	factors := append(small, big...)

	if k <= len(factors) {
		return factors[k-1]
	}
	return -1
}
