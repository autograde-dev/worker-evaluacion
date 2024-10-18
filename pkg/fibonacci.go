package fibonacci

func Calculate(n int) int {
	if n <= 1 {
		return n
	}
	return Calculate(n-1) + Calculate(n-2)
}
