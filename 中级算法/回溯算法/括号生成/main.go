package 括号生成

var parenthesis []string

func generateParenthesis(n int) []string {
	parenthesis = []string{}
	dfs("", n, n)
	return parenthesis
}

func dfs(combination string, left, right int) {

	// 生成完所有的左右括号
	if left == 0 && right == 0 {
		parenthesis = append(parenthesis, combination)
	}

	// 剩余左括号比右括号多，即添加这个右括号会导致这组括号无效。
	if left > right {
		return
	}

	// 有剩余左括号
	if left > 0 {
		dfs(combination+"(", left-1, right)
	}

	// ...
	if right > 0 {
		dfs(combination+")", left, right-1)
	}
}
