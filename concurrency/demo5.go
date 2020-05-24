package main

import (
	"fmt"
	"time"
)

func main()  {
	go func() {
		fmt.Println("ooooooooooooojbk")
	}()

	doneChan := make(chan string)

	go func() {
		doneChan<-"done"
	}()

	<-doneChan
	<-time.After(time.Second * 5) //5秒钟之后从通道里接收数据
}