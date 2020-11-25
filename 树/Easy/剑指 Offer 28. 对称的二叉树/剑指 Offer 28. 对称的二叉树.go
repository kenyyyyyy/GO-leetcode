package 剑指_Offer_28__对称的二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Right, root.Left)
}

func check(r, l *TreeNode) bool {
	// 如果都是nil,可以看成相等
	if r == nil && l == nil {
		return true
	}
	// 其中一个为空
	if r == nil || l == nil {
		return false
	}
	// 值不相等
	if r.Val != l.Val {
		return false
	}
	// 左子树的左子树合右子树的右子树相比 && 左子树的右子树和右子树的左子树相比
	return check(l.Left, r.Right) && check(l.Right, r.Left)
}
