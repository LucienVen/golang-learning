package part3

import (
	"io/ioutil"
	"net"
	"testing"
)

func init()  {
	daemonStart2 := startNetworkDaemon()
	daemonStart2.Wait()
}


func BenchStartNetworkDaemon(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8082")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}

		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("connot read: %v", err)
		}

		conn.Close()
	}
}
