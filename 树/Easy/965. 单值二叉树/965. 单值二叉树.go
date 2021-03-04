package _965__单值二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {

	return isUnequal(root, root.Val)
}

func isUnequal(root *TreeNode, p int) bool {
	if root == nil {
		return true
	}
	if root.Val != p {
		return false
	}
	return isUnequal(root.Left, p) && isUnequal(root.Right, p)
}
