package main

/* 在数组中取出一个或多个不相邻数，使其和最大，即找到max(不相邻元素组成的子数组)。 */

import "fmt"

func main()  {
	//toDoList := []int{1, 2, 4, 1, 7, 8, 3}
	toDoList := []int{3,1,6,10}


	// 递归
	fmt.Println("递归实现：", RecOpt(toDoList, len(toDoList)-1))

	//DTOpt(toDoList)
	fmt.Println("动态规划：", DTOpt(toDoList))

}

/* 递归 */
func RecOpt(list []int, i int) int  {
	if i == 0 {
		return list[0]
	} else if i == 1 {
		return intMax(list[0], list[1])
	} else {
		res1 := RecOpt(list, i-2) + list[i]
		res2 := RecOpt(list, i-1)

		return intMax(res1, res2)
	}
}

/* 动态规划 */
func DTOpt(arr []int) int {
	opt := make([]int, len(arr))
	opt[0] = arr[0]
	opt[1] = intMax(arr[0], arr[1])

	for i := 2; i < len(arr); i++  {
		fmt.Println(i)
		res1 := opt[i-2] + arr[i]
		res2 := opt[i-1]

		opt[i] = intMax(res1, res2)
	}
	//fmt.Println(opt)
	return opt[len(arr)-1]
}


func intMax(x,y int) int {
	if x > y {
		return x
	}

	return y
}