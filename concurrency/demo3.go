package main

import (
	"fmt"
	"time"
)

// 缓存队列

func main()  {
	bufferedChan := make(chan string, 1)

	go func() {
		bufferedChan<-"first"
		fmt.Println("Send 1st")

		bufferedChan<-"second"
		fmt.Println("Send 2nd")

		bufferedChan<-"third"
		fmt.Println("Send 3rd")
	}()
	//<-time.After(time.Second * 5)

	go func() {
		firstRead := <-bufferedChan
		fmt.Println("Receiving.....")
		fmt.Println(firstRead)

		secondRead := <-bufferedChan
		fmt.Println("Receiving.....")
		fmt.Println(secondRead)

		thirdRead := <-bufferedChan
		fmt.Println("Receiving.....")
		fmt.Println(thirdRead)
	}()
	<-time.After(time.Second * 1)

}