package _617合并二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// 当前t1节点为空，就返回t2的节点
	if t1 == nil {
		return t2
	}
	// 当前t2节点为空，就返回t1的节点
	if t2 == nil {
		return t1
	}
	// 以t1树为基点，把t2的节点值或者节点
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
