package leetcode

func sortColors(nums []int)  {
	var r,g,b = 0,0,0
	for _,n := range nums {
		if n == 0 {
			r++
		}
		if n == 1 {
			g++
		}
		if n == 2 {
			b++
		}
	}
	for i,_ := range nums {
		if r > 0 {
			nums[i] = 0
			r--
			continue
		}
		if g > 0 {
			nums[i] = 1
			g--
			continue
		}
		if b > 0 {
			nums[i] = 2
			b--
			continue
		}
	}
}

const (
	red   = 0
	white = 1
	blue  = 2
)

func sortColors1(nums []int) {
	swap := func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}

	nextRed := 0
	nextBlue := len(nums) - 1

	for i := 0; i <= nextBlue; i++ {
		switch nums[i] {
		case red:
			swap(nextRed, i)
			nextRed++
			continue
		case blue:
			swap(nextBlue, i)
			nextBlue--
			i--
			continue
		case white:
		}
	}
}