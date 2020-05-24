package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// 打印 2-100 之间的全部质数

var goal int

func primestask(c chan int)  {
	p := <-c

	if p > goal {
		os.Exit(0)
	}

	fmt.Println(p)

	nc := make(chan int)

	go primestask(nc)

	for {
		i := <-c

		if i%p != 0 {
			nc <- i
		}
	}
}

func main()  {
	flag.Parse()

	args := flag.Args()

	if args != nil && len(args) > 0 {
		var err error
		goal, err = strconv.Atoi(args[0])
		if err != nil {
			goal = 10
		}

	} else {
		goal = 10
	}

	fmt.Println("goal = ", goal)

	c := make(chan int)

	go primestask(c)

	for i := 2;; i++ {
		c <- i
	}
}