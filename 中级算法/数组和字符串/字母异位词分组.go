package 数组和字符串

import "sort"

func groupAnagrams(strs []string) [][]string {
	// 利用hash表将相同字母组合到一起
	hashTable:=make(map[string][]string,0)
	for _,str:=range strs{
		s:=[]byte(str)
		// 将其排序,拥有相同字母的str，排序后值一样
		sort.Slice(s, func (i,j int) bool{
			return s[i]<s[j]
		} )
		s1:=string(s)
		hashTable[s1]=append(hashTable[s1],str)
	}
	ans:=[][]string{}
	for _,m:=range hashTable{
		ans=append(ans,m)
	}
	return ans
}
