package main

import (
	"fmt"
	"sync"
)

func main()  {
	var numCalcsCreated int

	calcPool := &sync.Pool{New: func() interface{}{
		numCalcsCreated += 1
		men := make([]byte, 1024)
		return &men
	}}

	// 用4KB初始化pool
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWokers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWokers)

	for i := numWokers; i > 0; i-- {
		go func() {
			defer wg.Done()
			men := calcPool.Get().(*[]byte) // 断言
			defer calcPool.Put(men)
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculations were created", numCalcsCreated)
}
