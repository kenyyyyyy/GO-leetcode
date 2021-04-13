package 斐波那契数

// 快速幂
func quickPow(x float64, n int) (ans float64) {
	ans = 1.0
	for n > 0 {
		if n%2 == 1 {
			ans *= x
		}
		x *= x
		n = n / 2
	}
	return
}
