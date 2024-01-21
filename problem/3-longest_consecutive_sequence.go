package problem

/*
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
*/
func longestConsecutive(nums []int) int {
	hash := make(map[int]bool)
	for _, v := range nums {
		hash[v] = true
	}
	res := 0
	for _, num := range nums {
		if !hash[num-1] { // 剪枝遍历
			counter := 1
			v := num
			for hash[v+1] {
				v++
				counter++
			}
			if counter > res {
				res = counter
			}
		}
	}
	return res
}
