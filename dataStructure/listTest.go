package main

/* 列表实现（双向链表） */

import (
	"fmt"
	b "go.learning/dataStructure/baseStruct"
)

func main()  {
	list := new(b.DoubleList)

	// 在列表头部插入新元素
	list.AddNodeFromHead(0, "I")
	list.AddNodeFromHead(0, "love")
	list.AddNodeFromHead(0, "you")
	list.AddNodeFromHead(0, "four")

	// 在列表尾部插入新元素
	list.AddNodeFromTail(0, "may")
	list.AddNodeFromTail(0, "happy")

	fmt.Println("BEGIN: ",list)
	fmt.Println("list len: ", list.Len())
	segmentation()
	//
	//fmt.Println(list.IndexFromHead(0))
	//fmt.Println(list.IndexFromHead(1))
	//fmt.Println(list.IndexFromHead(2))

	//正常遍历
	for i := 0; i < list.Len(); i++ {
		// 从头部开始索引
		node := list.IndexFromHead(i)
		fmt.Println("node: ", node, &node)

		// 节点为空不可能，因为list.Len()使得索引不会越界
		//if !node.IsNil() {
		//	fmt.Println(node.GetValue())
		//}
	}


	segmentation()
}

func segmentation()  {
	fmt.Println("--------------------------------")
}