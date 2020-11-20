package _589N叉树的前序遍历

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	// 前序遍历，先记录根节点值
	ans = append(ans, root.Val)
	for _, node := range root.Children {
		ans = append(ans, preorder(node)...)
	}
	return ans
}
