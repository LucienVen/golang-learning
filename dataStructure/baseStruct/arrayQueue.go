package baseStruct

import (
	"fmt"
	"sync"
)

/* 队列（数组实现） */

type ArrayQueue struct {
	array 	[]string
	size 	int
	lock 	sync.Mutex
}

/* 入队 */
func (queue *ArrayQueue) Push(v string)  {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 入队, 新元素放在尾部
	queue.array = append(queue.array, v)

	queue.size = queue.size + 1
}

/* 出队 */
func (queue *ArrayQueue) Pop() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("queue empty")
	}

	// 处理空间占用
	v := queue.array[0]

	newQueue := make([]string, queue.size - 1, queue.size - 1)
	for i := 1; i < queue.size; i ++ {
		newQueue[i - 1] = queue.array[i]
	}
	fmt.Println("TEST: ", newQueue)
	queue.array = newQueue

	queue.size = queue.size - 1

	return v
}

/* 查看队首元素 */
func (queue *ArrayQueue) First() string {
	if queue.size == 0 {
		panic("queue empty")
	}

	return queue.array[0]
}

/* 查看队尾元素 */
func (queue *ArrayQueue) Last() string {
	if queue.size == 0 {
		panic("queue empty")
	}

	return queue.array[queue.size - 1]
}

/*  判断是否为空 */
func (queue *ArrayQueue) IsEmpty() bool {
	return queue.size == 0
}

/* 查看队列元素个数 */
func (queue *ArrayQueue) Size() int {
	return queue.size
}

