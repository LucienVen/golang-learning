package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "9992", "port")

var count int = 1

func main()  {
	flag.Parse()

	conn, err := net.Dial("tcp", *host + ":" + *port)
	//conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("Connecting to " + *host + ":" + *port)

	done := make(chan string)

	go handleWirte(conn, done)
	go handleRead(conn, done)

	fmt.Println(<-done)
	fmt.Println(<-done)

}

func handleWirte(conn net.Conn, done chan string)  {
	for i := 10; i > 0; i-- {
		// 在循环中 持续写入数据
		_, e := conn.Write([]byte("Hello " + strconv.Itoa(i) + "\r\n"))

		if e != nil {
			fmt.Println("Error to Send message because of ", e.Error())
			break
		}
	}
	Count("Wirte")
	done <- "Sent"
}

func handleRead(conn net.Conn, done chan string)  {
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Error to read message because of ", err)
		return
	}

	fmt.Println(string(buf[:reqLen-1]))
	Count("Read")
	done <- "Read"

}

func Count(target string)  {
	fmt.Printf("type: %s\t count:%d\n", target, count)
	count++

}