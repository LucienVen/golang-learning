package baseStruct

import "sync"

/* 列表实现 - 双端列表/双端队列 - 双向链实现 */

type DoubleList struct {
	head 	*ListNode	// 指向链表头部
	tail 	*ListNode	// 指向链表尾部
	len 	int			// 列表长度
	lock 	sync.Mutex
}

// 列表节点
type ListNode struct {
	pre 	*ListNode 	// 前驱节点
	next 	*ListNode	// 后驱节点
	Value 	string		// 值
}


/* 列表节点 操作 */

// 获取节点值
func (node *ListNode) GetValue() string {
	return node.Value
}

// 获取节点前驱节点
func (node *ListNode) GetPre() *ListNode {
	return node.pre
}

// 获取节点后驱节点
func (node *ListNode) GetNext() *ListNode {
	return node.next
}

// 是否存在前驱节点
func (node *ListNode) HasPre() bool {
	return node.pre != nil
}

// 是否存在后驱节点
func (node *ListNode) HasNext() bool {
	return node.next != nil
}

// 是否为空节点
func (node *ListNode) IsNil() bool {
	return node == nil
}

/* 添加节点到离头部的第N个元素之前， N=0表示新节点成为头节点 */
func (list *DoubleList) AddNodeFromHead(n int, v string)  {

}

























