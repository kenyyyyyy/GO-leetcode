package _1022从根到叶的二进制数之和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return findRootVal(root, 0)
}
func findRootVal(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}
	ans := val + root.Val
	// 叶子节点直接返回
	if root.Left == nil && root.Right == nil {
		return ans
	}
	// 继续*2往下层走
	return findRootVal(root.Left, ans*2) + findRootVal(root.Right, ans*2)
}
