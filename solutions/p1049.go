package solutions

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, stone := range stones {
		for i := target; i >= stone; i-- {
			if dp[i] < dp[i-stone]+stone {
				dp[i] = dp[i-stone] + stone
			}
		}
	}
	return sum - dp[target] - dp[target]
}
