package part4

import "fmt"

func ChannelReadLimit()  {
	// 在词法范围内实例化channel
	chanOwner := func() <-chan int {
		result := make(chan int)

		go func() {
			defer close(result)
			for i := 0; i < 10; i++ {
				result <- i
			}
		}()

		return result
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Recevied: %v. \n", result)
		}

		fmt.Printf("Done receving. \n")
	}

	results := chanOwner()
	consumer(results)
}
