package main

/* 树的遍历 */

import (
	"fmt"
	b "go.learning/dataStructure/baseStruct"
)

func main()  {
	t := &b.TreeNode{Data:"A"}
	t.Left = &b.TreeNode{Data:"B"}
	t.Right = &b.TreeNode{Data:"C"}
	t.Left.Left = &b.TreeNode{Data:"D"}
	t.Left.Right = &b.TreeNode{Data:"E"}
	t.Right.Left = &b.TreeNode{Data:"f"}

	fmt.Println("先序遍历: ")
	b.PreOrder(t)
	fmt.Printf("\n\n")

	fmt.Println("中序遍历: ")
	b.MidOrder(t)
	fmt.Printf("\n\n")

	fmt.Println("后序遍历")
	b.PostOrder(t)
	fmt.Printf("\n\n")

	fmt.Println("--------------------------------------")
	fmt.Println(" --- LayerOrder (queue) --- ")

	b.LayerOrder(t)

}
