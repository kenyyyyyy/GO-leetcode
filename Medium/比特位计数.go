package main

import "fmt"

func main() {
	fmt.Println(countBits(5))
}

func countBits(num int) []int {
	bit := make([]int, num+1)
	for i := 1; i <= num; i++ {
		if i%2 == 1 {
			bit[i] = bit[i-1] + 1
		} else {
			bit[i] = bit[i/2]
		}
	}
	return bit
}
