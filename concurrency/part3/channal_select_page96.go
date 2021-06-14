package part3

import (
	"fmt"
	"time"
)

func SelectDefault()  {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
	loop:
		for  {
			select {
			case <-done:
				break loop
			default:

			}

			// 模拟工作行为
			workCounter++
			time.Sleep(1 * time.Second)
		}

	fmt.Printf("Achieved %v cycles. \n", workCounter)
}
