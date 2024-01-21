package problem

func groupAnagrams(strs []string) [][]string {
	//hash := make(map[string][]string)
	//res := make([][]string, 0)
	//for _, str := range strs {
	//	bs := []byte(str)
	//	sort.Slice(bs, func(i, j int) bool {
	//		return bs[i] < bs[j]
	//	})
	//	if v, ok := hash[string(bs)]; ok {
	//		v = append(v, str)
	//		hash[string(bs)] = v
	//	} else {
	//		hash[string(bs)] = []string{str}
	//	}
	//}
	//for _, v := range hash {
	//	res = append(res, v)
	//}
	//return res

	//更优时间复杂度,hash表
	hash := make(map[[26]int][]string) //key为数组类型
	for _, str := range strs {
		counter := [26]int{}
		for _, v := range str {
			counter[v-'a']++
		}
		hash[counter] = append(hash[counter], str)
	}
	res := make([][]string, 0, len(hash))
	for _, v := range hash {
		res = append(res, v)
	}
	return res
}
