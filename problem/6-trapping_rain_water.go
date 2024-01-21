package problem

/*
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
*/
func trap(height []int) int {
	//双指针未优化 -- 时间复杂度O(n^2)
	//res := 0
	//for i := range height {
	//	if i == 0 || i == len(height)-1 {
	//		continue
	//	}
	//	// res[i]每列的雨水值算法: min(leftHeightMax,rightHeightMax)-height[i]
	//	leftHeightMax, rightHeightMax := height[i], height[i]
	//	for j := i + 1; j < len(height); j++ {
	//		if height[j] > rightHeightMax {
	//			rightHeightMax = height[j]
	//		}
	//	}
	//	for j := i - 1; j >= 0; j-- {
	//		if height[j] > leftHeightMax {
	//			leftHeightMax = height[j]
	//		}
	//	}
	//	res += Min(leftHeightMax, rightHeightMax) - height[i]
	//}
	//return res

	// 双指针优化 -- 时间复杂度O(n)
	//if len(height) == 0 {
	//	return 0
	//}
	//// 优化计算leftHeightMax,rightHeightMax的算法
	//leftHeightMax, rightHeightMax := make([]int, len(height)), make([]int, len(height))
	//leftHeightMax[0] = height[0]
	//for i := 1; i < len(height); i++ {
	//	leftHeightMax[i] = Max(leftHeightMax[i-1], height[i])
	//}
	//rightHeightMax[len(height)-1] = height[len(height)-1]
	//for i := len(height) - 2; i >= 0; i-- {
	//	rightHeightMax[i] = Max(rightHeightMax[i+1], height[i])
	//}
	//res := 0
	//for i := range height {
	//	if i == 0 || i == len(height)-1 {
	//		continue
	//	}
	//	res += Min(leftHeightMax[i], rightHeightMax[i]) - height[i]
	//}
	//return res

	// 单调栈 -- 时间复杂度O(n)
	res := 0
	var ascStack []int //单调递增栈
	ascStack = append(ascStack, 0)
	for i := 1; i < len(height); i++ {
		popIndex := ascStack[len(ascStack)-1]
		if height[i] < height[popIndex] {
			ascStack = append(ascStack, i)
		} else if height[i] == height[popIndex] {
			ascStack = ascStack[:len(ascStack)-1]
			ascStack = append(ascStack, i)
		} else {
			for len(ascStack) > 0 && height[i] > height[ascStack[len(ascStack)-1]] {
				top := ascStack[len(ascStack)-1]
				ascStack = ascStack[:len(ascStack)-1]
				if len(ascStack) > 0 {
					h := Min(height[i], height[ascStack[len(ascStack)-1]]) - height[top]
					w := i - ascStack[len(ascStack)-1] - 1
					res += h * w
				}
			}
			ascStack = append(ascStack, i)
		}
	}
	return res
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
