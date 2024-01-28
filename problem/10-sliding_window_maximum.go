package problem

/*
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3

	1 [3  -1  -3] 5  3  6  7       3
	1  3 [-1  -3  5] 3  6  7       5
	1  3  -1 [-3  5  3] 6  7       5
	1  3  -1  -3 [5  3  6] 7       6
	1  3  -1  -3  5 [3  6  7]      7
*/
func maxSlidingWindow(nums []int, k int) []int {
	// 单调递减栈 -> 栈头到栈尾元素单调递减
	var res []int
	if len(nums) < k {
		return res
	}
	var descStack []int
	for i := 0; i < k; i++ {
		for len(descStack) > 0 && nums[i] > descStack[len(descStack)-1] {
			descStack = descStack[:len(descStack)-1]
		}
		descStack = append(descStack, nums[i])
	}
	res = append(res, descStack[0])
	for i := k; i < len(nums); i++ {
		// 移除元素
		removeNum := nums[i-k]
		if removeNum == descStack[0] {
			descStack = descStack[1:]
		}
		for len(descStack) > 0 && nums[i] > descStack[len(descStack)-1] {
			descStack = descStack[:len(descStack)-1]
		}
		descStack = append(descStack, nums[i])
		res = append(res, descStack[0])
	}
	return res
}
