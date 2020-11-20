package _24__二叉树中的最大路径和

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	// 全局变量最大值MAX
	var maxSum = math.MinInt32
	// 递归函数
	var oneSideMax func(root *TreeNode) int
	//
	oneSideMax = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 与0比较，如果这个子树贡献的最大路径和小于0，就视为这个子树不存在，不走这个子树。
		left := max(0, oneSideMax(root.Left))
		right := max(0, oneSideMax(root.Right))
		// 求当前节点的最大路径和，并且与最大值MAX比较
		maxSum = max(maxSum, left+right+root.Val)
		// 贡献给父节点的最大路径和，因为祖父节点的路径只能从两个子孙节点二选其一
		return root.Val + max(left, right)
	}
	// 开始递归
	oneSideMax(root)
	return maxSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
