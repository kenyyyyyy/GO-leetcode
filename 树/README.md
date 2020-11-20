**树的算法框架**

对于树的算法结构的使用，无非就是遍历。

**二叉树的遍历：**
```golang
前序遍历
func traverse(root *TreeNode){
 	if root == nil {
 		return 
 	}
 	fmt.Println(root.Val)
 	traverse(root.Left)
 	traverse(root.Right)
 }
```
```golang
中序遍历
func traverse(root *TreeNode){
 	if root == nil {
 		return 
 	}
 	traverse(root.Left)
 	fmt.Println(root.Val)
 	traverse(root.Right)
 }
```
```golang
后序遍历
func traverse(root *TreeNode){
 	if root == nil {
 		return 
 	}
 	traverse(root.Left)
 	traverse(root.Right)
 	fmt.Println(root.Val)
 }
```
**N叉树的遍历**
```golang
前序遍历
func traverse(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	// 前序遍历，先记录根节点值
	ans = append(ans, root.Val)
	for _, node := range root.Children {
		ans = append(ans, traverse(node)...)
	}
	return ans
 }
```