package _112路径总和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	// 判断当前节点是否为叶子节点
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	// 减去当前父节点的值
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
