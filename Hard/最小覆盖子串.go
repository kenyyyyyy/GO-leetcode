package Hard

func minWindow(s string, t string) string {

	needMap:=make(map[byte]int)	// 各个字母需要的个数
	needCnt:=len(t)  			// 需要的总数
	for i:=0;i<len(t);i++{
		needMap[t[i]]++
	}

	left:=0
	sLen:=len(s)
	start,maxLen:=0,sLen+1 // 默认起始坐标和长度
	for i:=0;i<sLen;i++{

		if needMap[s[i]]>0{
			// 如果遇到需要的字母将needCnt减一
			needCnt--
		}
		// 遍历的所有字母标记减一，方便去掉多余字母
		needMap[s[i]]--

		// 所需要的数量为0,即left~i子串里面包含了所有的t字母
		if needCnt==0{

			// 去掉多余的字母,如果needMap[s[left]]为0表示找到了一个必要的字母退出循环
			for needMap[s[left]] != 0 {
				needMap[s[left]]++
				left++
			}
			// 记录长度和left
			if maxLen>i-left+1{
				start=left
				maxLen=i-left+1
			}
			// 将left往前走
			needMap[s[left]]++
			left++
			needCnt++
		}

	}

	if maxLen>sLen{
		return ""
	}

	return s[start:start+maxLen]
}
