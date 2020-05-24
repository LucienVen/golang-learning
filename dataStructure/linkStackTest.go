package main

/* 测试 链表栈 */

import (
	"fmt"
	b "go.learning/dataStructure/baseStruct"
)

func main()  {
	linkStack := new(b.LinkStack)

	// 打印
	fmt.Println(linkStack)

	// 入栈
	linkStack.Push("dog-1")
	linkStack.Push("pig-2")
	linkStack.Push("cat-3")

	fmt.Println(linkStack)


	// 查看size
	fmt.Printf("size: %d\n", linkStack.Size())

	// 是否为空
	fmt.Printf("is empty: %v\n", linkStack.IsEmpty())

	// 查看栈顶元素
	fmt.Printf("top stack elem: %v\n", linkStack.Peek())

	// 出栈
	fmt.Printf("pop stack: %v\n", linkStack.Pop())

	// 查看栈顶元素
	fmt.Printf("top stack elem: %v\n", linkStack.Peek())

	// 查看size
	fmt.Printf("size: %d\n", linkStack.Size())

	// 打印
	fmt.Println(linkStack)
}

