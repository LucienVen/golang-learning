package main

import (
	"fmt"
	"time"
)


// golang 并发编程
// 学习文档：https://learnku.com/go/t/23461/learning-the-concurrency-of-go-through-illustrations

func finder1(resource []string) {
	fmt.Println("Found-1")
	//var oreList []string
	for _, value := range resource {

		if value == "ore" {
			fmt.Println("Found 1 find ore!")
			//oreList = append(oreList, value)
		}
	}
}

func finder2(resource []string) {
	fmt.Println("Found-2")

	//var oreList []string
	for _, value := range resource {

		if value == "ore" {
			fmt.Println("Found 2 find ore!")
			//oreList = append(oreList, value)
		}
	}
}

func main()  {
	theMine := []string{"rock", "ore", "ore", "rock", "ore"}
	go finder1(theMine)
	go finder2(theMine)
	<-time.After(time.Second * 5)
}