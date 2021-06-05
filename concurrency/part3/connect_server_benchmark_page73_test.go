package part3


// TODO some trouble
import (
	"io/ioutil"
	"net"
	"testing"
)

func init()  {
	daemonStart := startNetWorkDaemon()
	daemonStart.Wait()

}

func BenchNetworkRequest(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8081")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}

		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("connot read: %v", err)
		}

		conn.Close()
	}
}
