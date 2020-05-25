package baseStruct

import (
	"fmt"
	"sync"
)

/* 二叉树基础结构（链表实现） */

type TreeNode struct {
	Data 	string		// 节点存放数据
	Left 	*TreeNode 	// 左树节点
	Right	*TreeNode	// 右树节点
}

/* 先序遍历 */
/* 根节点 - 左子树 - 右子树 */
func PreOrder(tree *TreeNode)  {
	if tree == nil {
		return
	}

	// 先打印根节点
	fmt.Print(tree.Data, " ")
	// 再打印左子树
	PreOrder(tree.Left)
	// 最后打印右子树
	PreOrder(tree.Right)
}

/* 中序遍历 */
/* 左子树 - 根节点 - 右子树 */
func MidOrder(tree *TreeNode)  {
	if tree == nil {
		return
	}

	MidOrder(tree.Left)
	fmt.Print(tree.Data, " ")
	MidOrder(tree.Right)
}


/* 后序遍历 */
/* 左子树 - 右子树 - 根节点 */
func PostOrder(tree *TreeNode)  {
	if tree == nil {
		return
	}

	PostOrder(tree.Left)
	PostOrder(tree.Right)
	fmt.Print(tree.Data, " ")
}


/* ———————————————————————————————————————————————————————————— */
/* 层次遍历 - 广度优先 */
/* 需要辅以 链表队列 */
func LayerOrder(tree *TreeNode)  {
	if tree == nil {
		return
	}

	// 新建队列
	queue := new(LinkTreeQueue)
	// 根节点元素先入队
	queue.Add(tree)

	for queue.size > 0 {
		// 不断出队列
		elem := queue.Remove()

		// 打印值
		fmt.Print(elem.Data, " ")

		// 如果左子树不为空，则先入队
		if elem.Left != nil {
			queue.Add(elem.Left)
		}

		// 如果右子树不为空，则接着入队
		if elem.Right != nil {
			queue.Add(elem.Right)
		}
	}

	fmt.Printf("\n --- END --- \n")

}

/* 链表队列(单链表) - 树节点使用 */
/* 队列，先进先出 */
type LinkTreeNode struct {
	Next 	*LinkTreeNode	// 后驱节点
	Value	*TreeNode		// 树节点
}

type LinkTreeQueue struct {
	root	*LinkTreeNode	// 链表起点
	size	int				// 链表长度
	lock 	sync.Mutex		// 锁
}

/* 入队 */
func (queue *LinkTreeQueue) Add(tree *TreeNode)  {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 新节点
	node := new(LinkTreeNode)
	node.Value = tree

	if queue.root == nil {
		// 新节点成为队列头节点
		queue.root = node

	} else {
		tempNode := queue.root
		// 遍历，直到获取到最后的节点
		for tempNode.Next != nil {
			tempNode = tempNode.Next
		}

		// 当前最后节点，添加后驱节点
		tempNode.Next = node

	}

	queue.size = queue.size + 1

}


/* 出队 */
/* 先进先出，出队首元素 */
func (queue *LinkTreeQueue) Remove() *TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 如果是空队列
	if queue.size == 0 {
		return nil
	}

	// 获取将要出队的队首元素
	topNode := queue.root
	v := topNode.Value

	// 队首元素的后驱元素，成为新的队首元素
	nextNode := topNode.Next
	queue.root = nextNode

	queue.size = queue.size - 1
	return v
}

/* 元素数量 */
func (queue *LinkTreeQueue) Size() int {
	return queue.size
}
