package main

import (
	"fmt"
	b "go.learning/dataStructure/baseStruct"
)

//func linkNewTest()  {
//	// 第一个节点
//	node := &r.Ring{Value:-1}
//
//	//node1 := &r.Ring{Value:1}
//	//node2 := &r.Ring{Value:2}
//	//node3 := &r.Ring{Value:3}
//	//node4 := &r.Ring{Value:4}
//
//	// 连接新的节点
//	node.Link1(&r.Ring{Value:1})
//	node.Link1(&r.Ring{Value:2})
//	node.Link1(&r.Ring{Value:3})
//	node.Link1(&r.Ring{Value:4})
//
//
//	choose := node
//
//	for {
//		// 打印节点值
//		fmt.Println(choose.Value)
//		choose = choose.Next()
//
//		if choose == node {
//			return
//		}
//	}
//
//}
//
//func main()  {
//	linkNewTest()
//}

/*  TODO 复制 */
func linkNewTest() {
	// 第一个节点
	r := &b.Ring{Value: -1}

	// 链接新的五个节点
	r.Link1(&b.Ring{Value: 1})
	r.Link1(&b.Ring{Value: 2})
	r.Link1(&b.Ring{Value: 3})
	r.Link1(&b.Ring{Value: 4})

	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)

		// 移到下一个节点
		node = node.Next()

		//  如果节点回到了起点，结束
		if node == r {
			return
		}
	}
}

func main() {
	linkNewTest()
}
