package problem

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	hash := make(map[int]int)
	for i, v := range nums {
		if index, ok := hash[target-v]; ok {
			return []int{index, i}
		} else {
			hash[v] = i
		}
	}
	return []int{}
}
