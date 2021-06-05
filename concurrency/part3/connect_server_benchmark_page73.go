package part3

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

// 模拟连接服务器
func connectTosService() interface{} {
	time.Sleep(1 * time.Second)
	return struct {

	}{}
}

func startNetWorkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		server, err := net.Listen("tcp", "localhost:8081")
		if err != nil {
			log.Fatalf("cannot listien: %v", err)
		}
		
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Fatalf("cannot accept connect: %v", err)
				continue
			}

			connectTosService()
			fmt.Fprintf(conn, "")
			conn.Close()
		}
	}()

	return &wg
}