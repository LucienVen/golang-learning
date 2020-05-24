package main

import (
	"fmt"
	"time"
)

func main()  {
	theMine := []string{"rock", "ore", "ore", "rock", "ore"}
	//无缓存通道
	oreChan := make(chan string)

	go func(mine []string) {
		for _, item := range mine{
			oreChan<-item
		}
	}(theMine)

	go func() {
		foundOre := <-oreChan
		for i := 0; i < 3; i++ {
			fmt.Println("Miner: Received " + foundOre + " from finder")
		}
	}()
	<-time.After(time.Second * 3)
}