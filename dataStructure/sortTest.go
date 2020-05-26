package main

/* 测试排序算法 */

import (
	"fmt"
	s "go.learning/dataStructure/sortingAlgorithm"
)

func main()  {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	bubbleList := list
	//list2 := []int{2,1,3,4,5}
	fmt.Printf("%p, %p\n",list, bubbleList)
	/* 冒泡排序 */
	s.BubbleSort(bubbleList)
	fmt.Println("冒泡排序: ", bubbleList)
	//fmt.Println("list ======> ", list)
	fmt.Println("========================================")
	/* 选择排序 */
	selectlist := list
	s.SelectSort(selectlist)
	fmt.Println("选择排序: ", selectlist)


	/* 插入排序 */
	//fmt.Println("list ======> ", list)
	insertList := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	s.InsertSort(insertList)
	fmt.Println("------------------------------------")

	/* 希尔排序 */
	shellList := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	s.ShellSort(shellList)
	fmt.Println("希尔排序: ", shellList)
	fmt.Println("===================================")


	/* 归并排序 */




	//Test()

	//sliceTest()
}

func Test()  {
	i:=0
	for ; i<10; i++ {
		fmt.Println(i)
	}
	fmt.Println("res: ", i)

	j := 9
	for ; j >= 0; j-- {
		fmt.Println(j)
	}
	fmt.Println("result: ", j)
}

func sliceTest() {
	list := []int{2,1,3}
	list1 := list
	fmt.Printf("%p, %p\n", list, list1)
	list1 = append(list1, 10)
	fmt.Printf("%p\n", list1)
	fmt.Println(list)
	fmt.Println(list1)
}