package 斐波那契数


// 快速乘
func quickMultiply(a, b int) (c int) {
	for b > 0 {
		if b%2 == 1 {
			c += a
		}
		a += a
		b = b / 2
	}
	return c
}
