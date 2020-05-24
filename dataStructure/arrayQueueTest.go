package main

/* 测试 - 数组队列 */

import (
	"fmt"
	b "go.learning/dataStructure/baseStruct"
)

func main()  {
	// 初始化
	queue := new(b.ArrayQueue)

	fmt.Println("new queue: ", queue)
	// 是否为空
	fmt.Printf("queue empty: %v\n", queue.IsEmpty())
	// 入队
	queue.Push("pig-1")
	fmt.Println("size: ", queue.Size())

	queue.Push("cat-2")
	queue.Push("dog-3")

	fmt.Println("after queue: ", queue)


	// 查看队首元素
	fmt.Printf("top elem: %v\n", queue.First())

	// 查看队尾元素
	fmt.Printf("last elem: %v\n", queue.Last())
	// 出队
	fmt.Printf("pop elem: %v\n", queue.Pop())
	//fmt.Printf("pop elem: %v\n", queue.Pop())

	// 查看size
	fmt.Println("size: ", queue.Size())
	// 是否为空
	fmt.Printf("queue empty: %v\n", queue.IsEmpty())

	fmt.Println("queue end: ", queue)

}