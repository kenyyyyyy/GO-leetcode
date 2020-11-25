package _606根据二叉树创建字符串

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	// 叶子节点，直接返回值
	if t.Left == nil && t.Right == nil {
		return strconv.Itoa(t.Val)
	}
	// 只有左子树
	if t.Right == nil {
		return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")"
	}
	// 这种情况，可能包含只有右子树的情况，例如：1(2()(4))(3)
	return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")" + "(" + tree2str(t.Right) + ")"
}
