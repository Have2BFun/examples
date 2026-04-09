package examples

// Fibonacci 返回斐波那契数列的前 n 个数
func Fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	result := make([]int, n)
	if n >= 1 {
		result[0] = 1
	}
	if n >= 2 {
		result[1] = 1
	}

	for i := 2; i < n; i++ {
		result[i] = result[i-1] + result[i-2]
	}

	return result
}
