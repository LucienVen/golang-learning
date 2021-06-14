package main

import (
	"fmt"
	"go.learning/concurrency/part3"
)

func main()  {
	fmt.Println("=== Base Running Func ===")


	//part3.SimpleChannelDemo()
	//part3.ChannalBufDemo()

	part3.ChannalOwnerDemo()
}
