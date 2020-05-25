package baseStruct

import (
	"sync"
)

/* 列表实现 - 双端列表/双端队列 - 双向链实现 */

type DoubleList struct {
	head 	*ListNode	// 指向链表头部
	tail 	*ListNode	// 指向链表尾部
	len 	int			// 列表长度
	lock 	sync.Mutex
}

// 列表节点
type ListNode struct {
	prev 	*ListNode 	// 前驱节点
	next 	*ListNode	// 后驱节点
	Value 	string		// 值
}


/* 列表节点 操作 */

// 获取节点值
func (node *ListNode) GetValue() string {
	return node.Value
}

// 获取节点前驱节点
func (node *ListNode) GetPrev() *ListNode {
	return node.prev
}

// 获取节点后驱节点
func (node *ListNode) GetNext() *ListNode {
	return node.next
}

// 是否存在前驱节点
func (node *ListNode) HasPrev() bool {
	return node.prev != nil
}

// 是否存在后驱节点
func (node *ListNode) HasNext() bool {
	return node.next != nil
}

// 是否为空节点
func (node *ListNode) IsNil() bool {
	return node == nil
}

/* 添加节点到离头部的第N个元素"之前"， N=0表示新节点成为头节点 */
func (list *DoubleList) AddNodeFromHead(n int, v string)  {
	// 加并发锁
	list.lock.Lock()
	defer list.lock.Unlock()

	// 长度超出列表长度，panic
	if n > list.len {
		panic("N too bigger then list.len")
	}

	// 新节点
	newNode := new(ListNode)
	newNode.Value = v

	// 获取头节点
	node := list.head
	// 定位节点
	for i := 1; i <= n; i++ {
		node = node.next
	}

	// 判断定位节点是否为空
	if node.IsNil() {
		// 若为空，表示列表为空, 则新节点为列表头结点
		list.head = newNode
		list.tail = newNode
	} else {
		// 原定位节点的前驱节点为 当前操作节点
		prevNode := node.prev

		if prevNode.IsNil() {
			// 若定位的节点前驱为空，则表明定位到列表开头
			list.head = newNode
			newNode.next = node
			node.prev = newNode

		} else {
			// 定位节点后驱节点为新节点
			prevNode.next = newNode
			// 新节点的前驱节点为 定位节点
			newNode.prev = prevNode
			// 新节点的后驱节点为 原定位节点
			newNode.next = node
			// 原定位节点的前驱节点为 新节点
			node.prev = newNode
		}
	}

	// 列表长度+1
	list.len = list.len + 1

}

/* 从尾部某个位置开始添加节点 */
/* 添加节点到聊表尾部的第N个元素"之后"，N=0表示新节点成为新的尾节点 */
func (list *DoubleList) AddNodeFromTail(n int, v string)  {
	list.lock.Lock()
	defer list.lock.Unlock()

	// N不能超出索引长度
	if n > list.len {
		panic("N too bigger then list.len")
	}

	// 尾节点
	node := list.tail

	// 定位节点
	for i := 1; i <= n; i++ {
		node = node.prev
	}

	// 新节点
	newNode := new(ListNode)
	newNode.Value = v

	if node.IsNil() {
		// 若节点为空，则表明是列表新节点为尾节点
		list.tail = newNode
		//newNode.prev = node
		list.head = newNode

	} else {
		nextNode := node.next

		// 如果定位到的节点为nil，则新节点成为新的尾节点
		if nextNode.IsNil() {
			node.next = newNode
			newNode.prev = node
			list.tail = newNode
		} else {

			// 新节点插入定位节点之后
			node.next = newNode
			// 新节点的前驱是定位节点
			newNode.prev = node
			// 新节点插入定位节点后驱节点之前
			nextNode.prev = newNode
			// 新节点的后驱节点是 定位节点的原后驱节点
			nextNode.next = nextNode

		}
	}

	list.len = list.len + 1

}

/* 从头部开始某个位置获得列表节点 */
/* 从头部开始找，获取N+1个位置的节点，索引从0开始 */
func (list *DoubleList) IndexFromHead(n int) *ListNode {

	// 索引超出
	if n >= list.len  {
		return nil
	}

	// 获取头部节点
	node := list.head

	// 遍历
	for i := 1; i <= n; i++ {
		node = node.next
	}

	return node
}


/* 从尾部往前开始找，获取N+1位置的节点，索引从0开始*/
func (list *DoubleList) IndexFromTail(n int) *ListNode {
	// 索引超出
	if n >= list.len  {
		return nil
	}

	// 获取尾节点
	node := list.tail
	// 遍历
	for i := 1; i <= n; i++ {
		node = node.prev
	}
	return node
}


/* 从头部移除并返回某个位置的节点 */
/* 从头部开始往后找，获取第N+1位置的节点，移除并返回 */
func (list *DoubleList) PopFromHead(n int) *ListNode {
	list.lock.Lock()
	defer list.lock.Unlock()

	// 索引超出
	if n >= list.len  {
		return nil
	}

	// 定位节点
	node := list.head
	for i := 1; i <= n; i++ {
		node = node.next
	}

	// 获取定位节点的上一个节点与下一个节点
	prevNode := node.prev
	nextNode := node.next


	if prevNode.IsNil() && nextNode.IsNil() {
		// 前驱与后驱节点都为空，则当前节点为列表唯一节点
		list.head = nil
		list.tail = nil
	} else if prevNode.IsNil() {
		// 前驱节点为空，则当前节点是列表头节点
		list.head = nextNode
		nextNode.prev = nil
	} else if nextNode.IsNil() {
		// 后驱节点为空，则当前节点是列表尾节点
		list.tail = prevNode
		prevNode.next = nil
	} else {
		// 中间节点
		prevNode.next = nextNode
		nextNode.prev = prevNode
	}

	list.len = list.len - 1
	return node
}


/* 从尾部开始移除并返回某个位置的节点 */
/* 从尾部开始往前找，获取第N+1个位置的节点，并移除返回 */
func (list *DoubleList) PopFromTial(n int) *ListNode {
	list.lock.Lock()
	defer list.lock.Unlock()

	// 索引超出
	if n > list.len {
		return nil
	}

	// 定位节点
	node := list.tail
	for i := 1; i >= n; i++ {
		node = node.prev
	}

	prevNode := node.prev
	nextNode := node.next

	if prevNode.IsNil() && nextNode.IsNil() {
		// 唯一节点
		list.tail = nil
		list.head = nil
	} else if prevNode.IsNil() {
		// 当前节点为头节点
		list.head = nextNode
		nextNode.prev = nil
	} else if nextNode.IsNil() {
		// 当前节点是尾节点
		list.tail = prevNode
		prevNode.next = nil
	} else {
		// 当前节点是中间节点
		prevNode.next = nextNode
		nextNode.prev = prevNode
	}

	// 列表长度-1
	list.len = list.len - 1
	return node
}

/* 返回列表长度 */
func (list *DoubleList) Len() int {
	return list.len
}

/* 获取头节点 */
func (list *DoubleList) First() *ListNode {
	return list.head
}


































