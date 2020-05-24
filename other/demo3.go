package main


// <-c   这是什么意思

// sync.WaitGroup

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	//c := make(chan bool)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			//c <- true
			wg.Done()
		}(i)
	}

	wg.Wait()

	//for i := 0; i < 100; i++ {
	//	<-c
	//}
}