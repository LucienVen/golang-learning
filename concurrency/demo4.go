package main

import (
	"fmt"
	"time"
)

func main()  {
	theMine := []string{"rock", "ore", "ore", "rock", "ore"}
	oreChan := make(chan string)
	mineChan := make(chan string)

	lenMine := len(theMine)

	// finder
	go func(mine []string) {
		for _, item := range mine {
			if item == "ore" {
				oreChan <- item // 发送进入通道
			}
		}
	}(theMine)


	// ore breaker
	go func() {
		for i := 0; i < lenMine; i++ {
			foundOre := <-oreChan
			fmt.Println("From Finder: ", foundOre)
			mineChan<-"minedOre"
		}
	}()

	// smelter
	go func() {
		for i := 0; i < lenMine; i++ {
			minedOre := <-mineChan
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: ore is smelter")
		}
	}()

	<-time.After(time.Second * 5)

}