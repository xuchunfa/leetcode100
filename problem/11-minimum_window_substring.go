package problem

import "math"

// from: https://leetcode.cn/problems/minimum-window-substring/

/*
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
*/
func minWindow(s string, t string) string {
	var hash [128]int
	counter := 0 //用于收缩窗口的标志
	for _, v := range t {
		hash[int(v-'A')]++
		counter++
	}
	minLen := math.MaxInt32
	left, right, index := 0, 0, 0
	for right < len(s) {
		if hash[int(s[right]-'A')] > 0 {
			counter--
		}
		hash[s[right]-'A']-- //复数标记成可以被淘汰的字符
		if counter == 0 {
			for left < right && hash[s[left]-'A'] < 0 { //开始收缩窗口
				hash[s[left]-'A']++
				left++
			}
			if right-left+1 < minLen {
				index = left
				minLen = right - left + 1
			}
			hash[s[left]-'A']++ //淘汰t中一个字符使得窗口继续像右滑动
			left++
			counter++
		}
		right++
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[index : index+minLen]
}
