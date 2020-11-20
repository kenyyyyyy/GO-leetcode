**树的算法框架**

对于数的算法结构的使用，无非三种遍历。

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