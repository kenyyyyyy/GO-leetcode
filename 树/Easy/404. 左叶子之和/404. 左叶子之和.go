package _404__左叶子之和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {

	if root == nil {
		return 0
	}
	var sumOfLeft func(root *TreeNode)
	var ans int
	sumOfLeft = func(root *TreeNode) {
		if root.Left != nil {
			if isLeafNode(root.Left) {
				ans += root.Left.Val
			} else {
				sumOfLeft(root.Left)
			}
		}
		if root.Right != nil && !isLeafNode(root.Right) {
			sumOfLeft(root.Right)
		}

	}
	sumOfLeft(root)
	return ans
}

func isLeafNode(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	return false
}
