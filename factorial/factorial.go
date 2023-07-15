package factorial

func factorial(n int) int {
	fact := 1
	for j := 2; j <= n; j++ {
		fact = fact * j
	}
	return fact
}
