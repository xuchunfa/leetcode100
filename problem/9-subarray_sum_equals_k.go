package problem

/*
输入：nums = [1,1,1], k = 2
输出：2

输入：nums = [-1,1,0,1,-1], k = 0
输出：6

输入：nums = [-1,1,-1,1], k = 0
输出：4
*/

func subarraySum(nums []int, k int) int {
	var counter = map[int]int{0: 1}
	var preSum, res int
	for _, v := range nums {
		preSum += v
		// 每次累加得到的和 sum 都会被记录在 map 中，并且会检查是否存在 sum - k 的键。
		// 如果存在，表示从之前某个位置到当前位置的子数组和为 k。这是因为 sum - (sum - k) = k
		if val, _ := counter[preSum-k]; val > 0 {
			res += val
		}
		// preSum 存在和不存在两种情况
		counter[preSum]++
	}
	return res
}
