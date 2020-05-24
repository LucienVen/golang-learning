package main

import "fmt"

// 递归 - N 的阶乘

func Recursive(n int) int {
	if n == 0 {
		return 1
	}

	return Recursive(n - 1) * n
}

func main()  {
	n := 6
	result := Recursive(6)
	fmt.Printf("%d 阶乘结果: %d\n", n, result)
}