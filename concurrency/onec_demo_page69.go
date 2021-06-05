package main

import (
	"fmt"
	"sync"
)

func main()  {
	var count int

	increments := func() {
		count++
	}

	var once sync.Once
	var incr sync.WaitGroup

	incr.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer incr.Done()
			once.Do(increments)
		}()
	}

	incr.Wait()

	fmt.Println("count: ", count)
}
