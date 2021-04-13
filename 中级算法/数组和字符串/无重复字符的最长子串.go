package 数组和字符串

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}
	hashTable := make(map[byte]int)
	low := 0
	maxLen := 0

	for i := 0; i < len(s); i++ {

		// 如果s[i]出现过将更改low
		if index, ok := hashTable[s[i]]; ok {
			// 使用max防止low后退
			low = max(low, index+1)
		}
		hashTable[s[i]] = i
		maxLen = max(maxLen, i-low+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
