package _897递增顺序查找树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 第一种方法

func increasingBST(root *TreeNode) *TreeNode {
	tmp := make([]int, 0)
	head := &TreeNode{}
	next := head
	tmp = append(tmp, traverse(root)...)
	for _, v := range tmp {
		next.Right = &TreeNode{Val: v}
		next = next.Right
	}
	return head.Right
}

// 中序遍历
func traverse(root *TreeNode) (tmp []int) {
	if root == nil {
		return
	}
	tmp = append(tmp, traverse(root.Left)...)
	tmp = append(tmp, root.Val)
	tmp = append(tmp, traverse(root.Right)...)
	return
}

/* 第二种方法
var next *TreeNode  //定义一个全局的指针，保证递归的时候可以正确获取到当前遍历节点的父节点
func increasingBST(root *TreeNode) *TreeNode {
	head := &TreeNode{}
	next = head
	traverse(root)
	return head.Right
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left)
	root.Left = nil
	next.Right = root
	next = next.Right    // 移动全局指针，保证父节点的正确性
	traverse(root.Right)
}
*/
