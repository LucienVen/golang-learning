package part3

import (
	"bytes"
	"fmt"
	"os"
)

func ChannalBufDemo()  {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 2)

	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")

		for i := 0; i < 10; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for item := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received: %v.\n", item)
	}
}
