package Hard

func minWindow(s string, t string) string {

	needMap := make(map[byte]int, 0) // 各个字母需要的个数
	needCnt := len(t)                // 需要的总数
	for i := 0; i < len(t); i++ {
		needMap[t[i]]++
	}
	//
	left := 0
	ans := [2]int{0, len(s) + 1} // ans[0]子串起点,ans[1]长度,默认长度为len(s)+1
	for i := 0; i < len(s); i++ {

		if needMap[s[i]] > 0 {
			// 如果遇到需要的字母将needCnt减一
			needCnt--
		}
		// 遍历的所有字母标记减一，方便去掉多余字母
		needMap[s[i]]--
		// 所需要的数量为0,即left~i子串里面包含了所有的t字母
		if needCnt == 0 {
			for {
				// 去掉多余的字母,如果needMap[s[left]]为0表示找到了一个必要的字母退出循环
				if needMap[s[left]] == 0 {
					break
				}
				// left往前走的同时needMap[s[left]]--
				needMap[s[left]]++
				left++
			}
			// 记录长度和left
			if ans[1]-ans[0] > i-left+1 {
				ans[0] = left
				ans[1] = i + 1
			}
			// 将left往前走,
			needMap[s[left]]++
			needCnt++
			left++
		}

	}

	if ans[1] > len(s) {
		return ""
	}

	return s[ans[0]:ans[1]]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
