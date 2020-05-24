package main

import "fmt"

// 斐波那契数列

func F(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 1
	}

	return F(n-2) + F(n-1)
}

//func FT()  {
//	
//}

func main()  {
	n := 7
	fmt.Printf("斐波那契数列 - 第%d个数是: %d\n", n, F(n))
}