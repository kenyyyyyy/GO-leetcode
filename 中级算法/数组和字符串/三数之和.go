package 数组和字符串

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	// 先排序
	sort.Ints(nums)
	ans := make([][]int, 0)

	for k:=0;k<n-2;k++{
		// 以nums[k]为基数,如果大于0，后续不用再找直接返回结果
		if nums[k]>0{
			return ans
		}
		// 判断是否等于前面的值，即目前的nums[k]是否当过基数
		if k>0 && nums[k]==nums[k-1]{
			continue
		}

		// 双指针从k的右边和数组的尾部开始各自靠拢
		l,r:=k+1,n-1
		for l<r{
			sum:=nums[k]+nums[l]+nums[r]
			// 和等于0
			if sum==0{
				tmp:=[]int{nums[k],nums[l],nums[r]}
				ans=append(ans,tmp)
				for l<r && nums[r]==nums[r-1]{  //找到和上次不一样的值
					r--
				}
				for l<r && nums[l]==nums[l+1]{ //找到和上次不一样的值
					l++
				}

				r-- // 等于后一个相等的值
				l++ // 等于前一个不相等的值
			}else{  //如果不等于0
				if sum>0{
					for l<r && nums[r]==nums[r-1]{
						r--
					}
					r--
				}else{
					for l<r && nums[l]==nums[l+1]{
						l++
					}
					l++
				}

			}
		}

	}
	return ans
}