package part3

import "fmt"

func ChannalOwnerDemo()  {
	channalOwner := func() <-chan int {
		resultStream := make(chan int, 5)

		go func() {
			defer close(resultStream)
			for i := 0; i < 10; i++ {
				resultStream <- i
				fmt.Printf("Send: %d\n", i)
			}
		}()

		return resultStream
	}

	resultStream := channalOwner()
	for result := range resultStream{
		fmt.Printf("Receviced: %d\n", result)
	}

	fmt.Println("Done.")
}
