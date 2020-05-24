package main

/* 测试 - 链表数组 */

import (
	"fmt"
	b "go.learning/dataStructure/baseStruct"
)

func main()  {
	// 初始化
	linkQueue := new(b.LinkQueue)

	fmt.Println("initialization: ", linkQueue)

	// 入队
	linkQueue.Push("cat-1")
	linkQueue.Push("dog-2")
	linkQueue.Push("pig-3")

	fmt.Println("peek: ", linkQueue)

	// 查看size
	fmt.Println("size: ", linkQueue.Size())

	// 查看队首元素
	fmt.Println("top Elem: ", linkQueue.First())
	// 队尾
	fmt.Println("last elem: ", linkQueue.Last())

	// 出队
	fmt.Println("pop Elem: ", linkQueue.Pop())
	fmt.Println("pop Elem: ", linkQueue.Pop())

	// size
	fmt.Println("size: ", linkQueue.Size())

	// 队首队尾元素
	fmt.Println("top elem: ", linkQueue.First())
	fmt.Println("last elem: ", linkQueue.Last())

	// END
	fmt.Println("END: ", linkQueue)

}