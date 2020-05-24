package baseStruct

import "sync"

/* 链表队列 */

type LinkQueue struct {
	root 	*LinkNode  // 链表起点
	size 	int
	lock 	sync.Mutex
}

/* 链表节点 - 复用 */
//type LinkNode struct {
//	Next	*LinkNode		// 下一个节点
//	Value 	string
//}


/* 入队 */
func (queue *LinkQueue) Push(v string)  {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		newNode := new(LinkNode)
		newNode.Value = v
		queue.root = newNode

	} else {

		newNode := new(LinkNode)
		newNode.Value = v

		// 循环，查找到队列最后一个节点，然后把新节点加到当前最后节点的后面
		prevNode := queue.root
		for prevNode.Next != nil  {
			prevNode = prevNode.Next
		}

		prevNode.Next = newNode
	}

	queue.size = queue.size + 1

}

/* 出队 */
func (queue *LinkQueue) Pop() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("queue empty")
	}

	topNode := queue.root
	// 返回值
	v := topNode.Value

	// 出队处理，把当前root的下一个node，设置为root
	// TODO 链表的优势，增删节点，O(1)
	newTopNode := topNode.Next
	queue.root = newTopNode

	queue.size = queue.size - 1
	return v
}

/* 查看队首元素 */
func (queue *LinkQueue) First() string {
	if queue.size == 0 {
		panic("queue empty")
	}

	return queue.root.Value
}

/* 查看队尾元素 */
func (queue *LinkQueue) Last() string {
	if queue.size == 0 {
		panic("queue empty")
	}

	// TODO 链表实现的缺点，查找需要 O(n)
	tempNode := queue.root
	for tempNode.Next != nil {
		tempNode = tempNode.Next
	}

	return tempNode.Value
}

/* 是否为空 */
func (queue *LinkQueue) IsEmpty() bool {
	return queue.size == 0
}

/* 队列长度 */
func (queue *LinkQueue) Size() int {
	return queue.size
}