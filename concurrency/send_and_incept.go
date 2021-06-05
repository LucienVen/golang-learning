package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	c := sync.NewCond(&sync.Mutex{})

	queue := make([]interface{}, 0, 10)
	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()

		// 模拟头元素出队
		queue = queue[1:]

		fmt.Println("Remove from queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()  // 暂停 main goroutine，直到等待一个信号的条件以及发送
		}

		fmt.Println("Adding to queue")
		queue = append(queue, struct {}{})

		go removeFromQueue(1*time.Second)
		c.L.Unlock()
	}
}