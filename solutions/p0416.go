package solutions

func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = 0
	}
	for _, num := range nums {
		for j := target; j >= num; j-- {
			if dp[j] < dp[j-num]+num {
				dp[j] = dp[j-num] + num
			}

		}
	}
	return dp[target] == target
}
