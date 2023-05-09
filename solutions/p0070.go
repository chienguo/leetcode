package solutions

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		sum := a + b
		a = b
		b = sum
	}
	return b
}
