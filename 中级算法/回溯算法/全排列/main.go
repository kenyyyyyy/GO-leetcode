package 全排列

// 定义一个全局二维数组
var ans [][]int

func permute(nums []int) [][]int {

	n:=len(nums)

	ans=[][]int{}

	dfs(n,0,nums)
	return ans

}
// n:数组长度
// index:当前遍历层数
// used:传入到下一层的数组,其中index前面的表示使用过的数字
func dfs(n,index int,used []int){

	// 结束条件,当前基数遍历完所有层数,找到所有组合
	if n==index{
		ans=append(ans,used)
		return
	}

	// (1).每次遍历到下一层时都开始遍历从index遍历used,防止重复使用元素,因为index前面是已经使用过的
	for i:=index;i<n;i++{
		// 复制一个新切片,防止下一层的切片变化导致上一层的切片发生变化.(go里面的切片是引用类型,直接递归会导致修改的是同一个底层数组)
		tmp:=make([]int,n)
		copy(tmp,used)
		// 每次for循环都是横向的将切片里的元素当做这一层的各个基点,所以每次都需要和index位置上的值交换再传到下一层,参考(1)
		tmp[i],tmp[index]=tmp[index],tmp[i]
		dfs(n,index+1,tmp)
	}

}