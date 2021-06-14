package part3

import "fmt"

func SimpleChannelDemo()  {
	stringStream := make(chan string)
	go func() {
		stringStream <- "hello, channels!"
	}()

	fmt.Println(<-stringStream)
}
