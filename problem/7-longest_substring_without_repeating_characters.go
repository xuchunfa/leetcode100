package problem

/*
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
func lengthOfLongestSubstring(s string) int {
	var res = 0
	var queue []byte
	set := make(map[byte]struct{})
	for _, v := range []byte(s) {
		for {
			_, ok := set[v]
			if !ok {
				break
			}
			if len(queue) > 0 {
				delete(set, queue[0])
				queue = queue[1:]
			}
		}
		set[v] = struct{}{}
		queue = append(queue, v)
		res = Max(res, len(queue))
	}
	return res
}
