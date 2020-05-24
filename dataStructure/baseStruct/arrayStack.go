package baseStruct

import (
	"sync"
)

/* 数组栈 */

type ArrayStack struct {
	array	[]string	// 底层切片
	size 	int			// 栈元素数量
	lock 	sync.Mutex	// 并发安全锁
}

/* 入栈 */
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 后进元素放在数组后面
	stack.array = append(stack.array, v)

	// 栈元素+1
	stack.size += 1
}


/* 出栈 */
func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 栈元素已空
	if stack.size == 0 {
		panic("stack empty")
	}

	// 栈顶元素
	v := stack.array[stack.size - 1]

	// 处理空间占用 (创建新的切片)
	newArray := make([]string, stack.size - 1, stack.size - 1)
	// TODO test
	//fmt.Println(newArray)

	for i := 0; i < stack.size - 1; i++ {
		newArray[i] = stack.array[i]
	}

	stack.array = newArray

	// 栈元素-1
	stack.size = stack.size - 1

	// 返回
	return v
}

/* 获取栈顶元素 */
func (stack *ArrayStack) Peek() string {
	// 栈元素已空
	if stack.size == 0 {
		panic("stack empty")
	}

	v := stack.array[stack.size - 1]
	return v
}

/* 获取栈大小和判定是否为空 */
func (stack *ArrayStack) Size() int {
	return stack.size
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}




























