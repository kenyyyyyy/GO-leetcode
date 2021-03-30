package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 4,5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {
	// 维持一个严格递增的切片
	tails := make([]int, len(nums))

	size := 0
	for _, n := range nums {

		i, j := 0, size
		// 二分查找，在tails找到一个最小大于或者等于n的值
		for i != j {
			m := (i + j) >> 1
			if tails[m] < n {
				i = m + 1
			} else {
				j = m
			}
		}


		// 替换
		tails[j] = n
		// 大于当前tails所有的值，可以在循环里提前判断
		if i == size {
			size++
		}
	}

	return size
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
