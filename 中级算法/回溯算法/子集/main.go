package 子集

var ans [][]int

func subsets(nums []int) [][]int {
	ans = [][]int{}
	used := make([]int, 0)
	ans = append(ans, used)
	dfs(used, nums)
	return ans
}

// used:表示当前遍历已经使用过的数组
// res:表示当前层剩余遍历的数组
func dfs(used, res []int) {
	// 当没有剩余需要遍历的数时返回
	if len(res) == 0 {
		return
	}

	for i := 0; i < len(res); i++ {
		tmp := make([]int, len(used))
		copy(tmp, used)
		tmp = append(tmp, res[i])
		ans = append(ans, tmp)
		dfs(tmp, res[i+1:])
	}
}
