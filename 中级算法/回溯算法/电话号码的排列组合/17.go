package 电话号码的排列组合

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {

	combinations = []string{}
	if len(digits) == 0 {
		return combinations
	}
	// 可以看做一棵二叉树的root节点,传入 ”“ 与 0
	backTrack(digits, "", 0)
	return combinations

}
// digits:输入的数字组合
// combination:递归传递的字母组合
// index:当前遍历的数字位置
func backTrack(digits, combination string, index int) {

	// 找到所有的字母组合,开始插入到切片中
	if len(digits) == index {
		combinations = append(combinations, combination)
	} else {
		// 获取当前遍历的数字
		digit := string(digits[index])
		// 获取数字对应的字母组合
		letters := phoneMap[digit]
		// 遍历字母组合
		for i := 0; i < len(letters); i++ {
			// 将其与上层传进来的combination拼接,并且index+1
			backTrack(digits, combination+string(letters[i]), index+1)
		}
	}

}
