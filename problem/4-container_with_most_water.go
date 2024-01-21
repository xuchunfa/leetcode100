package problem

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		if height[left] < height[right] {
			area := height[left] * (right - left)
			if area > res {
				res = area
			}
			left++
		} else {
			area := height[right] * (right - left)
			if area > res {
				res = area
			}
			right--
		}
	}
	return res
}
