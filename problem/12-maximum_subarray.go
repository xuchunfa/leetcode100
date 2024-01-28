package problem

/*
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/
func MaxSubArray(nums []int) int {
	var res int
	if len(nums) <= 0 {
		return -1
	}
	res = nums[0]
	dp := make([]int, len(nums))
	for i := range nums {
		dp[i] = nums[i]
	}
	for i := 1; i < len(nums); i++ {
		dp[i] = Max(nums[i], dp[i-1]+nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
