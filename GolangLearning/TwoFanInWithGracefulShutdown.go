package golanglearning

import (
	"fmt"
	"time"
)

//   Presently this code is not scalble and works for less queue
// In this phase 1 iteration graceful shutdown of all channels and  avoid deadlock.

func fanIn(chan1, chan2 chan string, done chan struct{}) chan string {
	mergedChannel := make(chan string)

	go func(chan1, chan2 chan string, done chan struct{}) {
		defer close(mergedChannel)

		for {
			select {
			case <-done:
				fmt.Println("Stopping FanIn Process")
				return
			case msg1 := <-chan1:
				mergedChannel <- msg1
			case msg2 := <-chan2:
				mergedChannel <- msg2
			}
		}

	}(chan1, chan2, done)

	return mergedChannel
}

func simulateFanIn() {
	chan1 := make(chan string, 5)
	chan2 := make(chan string, 5)
	// Now to improve this code we will add done channel to stop the channel.

	done := make(chan struct{})

	mergedChannel := fanIn(chan1, chan2, done)

	// insetinn message in channel1
	go func(chan1 chan string) {
		for {
			select {
			case <-done:
				close(chan1)
				return
			case chan1 <- "Message from chan1":
				time.Sleep(time.Second)
			}

		}
	}(chan1)

	// insetinn message in channel2
	go func(chan2 chan string) {
		for {
			select {
			case <-done:
				close(chan2)
				return
			case chan2 <- "Message from chan2":
				time.Sleep(time.Second * 2)

			}
		}
	}(chan2)

	go func() {
		time.Sleep(time.Second * 6)
		close(done)
	}()

	for msg := range mergedChannel {
		fmt.Println(msg)
	}
	fmt.Println("Fan in Process exited cleanly")

}
