package main

import (
	"flag"
	"fmt"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "3333", "port")

func main()  {
	flag.Parse()
	fmt.Println(host, port)
}