package 排序和搜索

func sortColors(nums []int) {

	// q0表示数组中0应该在的位置
	// q2表示数组中2应该在的位置
	// 初始时，0应该在首部，2应该在尾部
	q0, q2 := 0, len(nums)-1

	// 当遍历到q2时，表示后面全是2了不用在遍历了。
	for i := 0; i <= q2; i++ {
		// 当遍历到0时将其与q0指针交换
		if nums[i] == 0 {
			nums[q0], nums[i] = nums[i], nums[q0]
			q0++
		}

		// 遍历到2时与q2指针交换
		if nums[i] == 2 {
			nums[q2], nums[i] = nums[i], nums[q2]
			q2--
			// 这里要特殊处理，因为q2是在当前遍历的前面所以不知道换过来的是什么
			// 当换过来的值不是1时，需要原地再处理一遍
			if nums[i] != 1 {
				i--
			}
		}
	}
	return
}
