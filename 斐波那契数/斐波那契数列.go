package 斐波那契数


type matrix [2][2]int

func fib(n int) int {

	if n < 2 {
		return n
	}
	base := matrix{{1, 1}, {1, 0}}
	res := pow(base, n-1)
	return res[0][0]

}

func multiply(a, b matrix) (c matrix) {

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return
}

func pow(a matrix, n int) (res matrix) {
	res = matrix{{1, 0}, {0, 1}}
	for n > 0 {
		if n%2 == 1 {
			res = multiply(res, a)
		}
		a = multiply(a, a)
		n = n / 2
	}
	return
}
