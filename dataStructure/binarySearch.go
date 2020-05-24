package main

import "fmt"

// 二分查找


func Search(array []int, target, left, right int) int {
	// 出界
	if left > right {
		return -1
	}

	// 中值
	mid := (left + right) / 2
	fmt.Println("mid: ", mid)
	midVal := array[mid]

	if target == midVal {
		return mid
	} else if target > midVal {
		return Search(array, target, mid-1, 1)
	} else {
		return Search(array, target, 1, mid+1)
	}
}

func main()  {
	target := 123
	searchArray := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	fmt.Println("length: ", len(searchArray))
	fmt.Printf("检索值: %d, 查询位置: %d\n", target, Search(searchArray, target, 0, len(searchArray)))
}