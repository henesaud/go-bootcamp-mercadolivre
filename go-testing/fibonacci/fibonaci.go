package fibonacci

func fibonacci(n int) []int {
	fibs := make([]int, n+1)
	fibs[0], fibs[1] = 0, 1

	for i := 2; i <= n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}

	return fibs
}
