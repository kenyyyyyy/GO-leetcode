package __从上到下打印二叉树_II

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	// 初始化队列
	queue := []*TreeNode{root}
	// 退出条件，队列为空
	for len(queue) > 0 {
		// 临时数组存储每层节点的值
		tmp := make([]int, 0)
		// 当前队列的长度
		n := len(queue)
		// 只需要遍历当前层次的节点就行
		for _, node := range queue[:n] {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 去掉已经遍历过的层次节点
		queue = queue[n:]
		// 保存当前节点的值
		ans = append(ans, tmp)
	}
	return ans
}
