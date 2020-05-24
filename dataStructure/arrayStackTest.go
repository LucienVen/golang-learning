package main

import (
	"fmt"
	b "go.learning/dataStructure/baseStruct"
)

func main()  {
	arrayStack := new(b.ArrayStack)

	// 入栈
	arrayStack.Push("cat-1")
	arrayStack.Push("dog-2")
	arrayStack.Push("pig-3")

	// 打印栈
	fmt.Printf("total stack: %v\n", arrayStack)

	// 查看size
	fmt.Printf("size: %d\n", arrayStack.Size())
	// 查看是否为空
	fmt.Printf("is empty? %v\n", arrayStack.IsEmpty())
	// 查看栈顶元素
	fmt.Printf("stack top elem: %s\n", arrayStack.Peek())

	// 出栈
	fmt.Printf("stack pop: %s\n", arrayStack.Pop())

	// 打印栈
	fmt.Printf("total stack: %v\n", arrayStack)
}
