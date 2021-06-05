package main

import "fmt"

/* 反转一个单链表。 */

//输出: 5->4->3->2->1->NULL

type LinkNode struct {
	Next 		*LinkNode
	Value 		int
}



func main()  {
	node := new(LinkNode)
	node.Value = 5

	node1 := new(LinkNode)
	node1.Value = 4
	node.Next = node1

	node2 := new(LinkNode)
	node2.Value = 3
	node1.Next = node2

	node3 := new(LinkNode)
	node3.Value = 2
	node2.Next = node3

	node4 := new(LinkNode)
	node4.Value = 1
	node3.Next = node4

	/* 遍历 */
	fmt.Println("source list: ")
	Trans(node)

	/* 反转 */
	reverseList := ReverseList(node)
	fmt.Println("reverse list: ")

	Trans(reverseList)




}

/* 反转链表 */
func ReverseList(head *LinkNode) *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	var prev *LinkNode

	for cur != nil {
		// 先处理当前节点，然后更换当前节点，再处理
		cur.Next, prev, cur = prev, cur, cur.Next
	}

	return prev

}

/* 遍历链表 */
func Trans(node *LinkNode)  {
	if node == nil {
		fmt.Println("node is nil")
	}

	cur := node

	for cur != nil  {
		fmt.Println(cur.Value)
		cur = cur.Next
	}
}