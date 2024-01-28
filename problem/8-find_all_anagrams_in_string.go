package problem

/*
示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
*/
func FindAnagrams(s string, p string) []int {
	var res []int
	if len(s) < len(p) {
		return res
	}
	var pCounter, sCounter [26]int
	for i, v := range p {
		pCounter[v-'a']++
		sCounter[s[i]-'a']++
	}
	if pCounter == sCounter {
		res = append(res, 0)
	}
	for i := 0; i < len(s)-len(p); i++ {
		sCounter[s[i]-'a']--
		sCounter[s[i+len(p)]-'a']++
		if pCounter == sCounter {
			res = append(res, i+1)
		}
	}
	return res
}
