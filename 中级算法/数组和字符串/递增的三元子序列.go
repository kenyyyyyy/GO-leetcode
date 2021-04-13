package 数组和字符串

func increasingTriplet(nums []int) bool {

	n := len(nums)
	ans := make([]int, n)
	ans[0] = nums[0]
	maxLen := 1
	for i := 1; i < n; i++ {
		if nums[i] > ans[maxLen-1] {
			ans[maxLen] = nums[i]
			maxLen++
		} else {
			left, right := 0, maxLen-1
			for left < right {
				mid := (left + right) >> 1
				if nums[i] > ans[mid] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			ans[left] = nums[i]
		}
	}
	if maxLen >= 3 {
		return true
	}
	return false
}
