package solutions

func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	var c int
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
