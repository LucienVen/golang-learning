package baseStruct

import (
	"sync"
)

/* 链表栈 - 下压栈（后进先出） */

type LinkStack struct {
	root	*LinkNode		// 链表根节点
	size 	int				// 链表元素数量
	lock 	sync.Mutex		// 并发安全锁
}

/* 链表节点 */
type LinkNode struct {
	Next	*LinkNode		// 下一个节点
	Value 	string
}


/* 入栈 */
func (stack *LinkStack) Push(v string)  {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 如果栈为空
	if stack.size == 0 {
		newNode := new(LinkNode)
		stack.root = newNode
		stack.root.Value = v
		//fmt.Println("First: ", stack)
	} else {
		preNode := stack.root

		// 创建新节点
		newNode := new(LinkNode)
		newNode.Value = v

		// 原来的链表放在新元素后面
		newNode.Next = preNode

		// 新节点放在开头
		stack.root = newNode
	}

	// 栈元素加1
	stack.size = stack.size + 1
}

/* 出栈 */
func (stack *LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 若栈为空
	if stack.size == 0 {
		panic("stack empty")
	}

	// 取值
	topNode := stack.root
	v := topNode.Value

	// 后续节点更新为根节点
	stack.root = topNode.Next

	stack.size = stack.size - 1

	return v
}

/* 获取栈顶元素 （不出栈） */
func (stack *LinkStack) Peek() string {
	if stack.size == 0 {
		panic("stack empty")
	}

	v := stack.root.Value
	return v
}

/* 获取栈大小和判定是否为空 */
func (stack *LinkStack) Size() int {
	return stack.size
}

func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}









