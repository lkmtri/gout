package gout

func MinInt(nums ...int) int {
	if len(nums) == 0 {
		panic("No parameters are passed to MinInt")
	}

	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		if ans < nums[i] {
			ans = nums[i]
		}
	}

	return ans
}

func AbsInt(num int) int {
	if num >= 0 {
		return num
	}

	return -num
}
