package golanglearning

import (
	"fmt"
	"sync"
	"time"
)

func FanInMultiple(done chan struct{}, channels ...chan string) chan string {
	mergedChannel := make(chan string)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(ch chan string) {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				case msg, ok := <-ch:
					if !ok {
						return
					}
					select {
					case <-done:
						return
					case mergedChannel <- msg:
					}
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(mergedChannel)
	}()

	return mergedChannel
}

func SimulateMultipleFanIn() {
	fmt.Println("Hello World! Starting Fan-In simulation...")

	chan1 := make(chan string)
	chan2 := make(chan string)
	chan3 := make(chan string)
	done := make(chan struct{})

	mergedChannel := FanInMultiple(done, chan1, chan2, chan3)

	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-done:
				close(chan1)
				return
			case chan1 <- fmt.Sprintf("Message %d from chan1", i):
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-done:
				close(chan2)
				return
			case chan2 <- fmt.Sprintf("Message %d from chan2", i):
				time.Sleep(time.Millisecond * 700)
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-done:
				close(chan3)
				return
			case chan3 <- fmt.Sprintf("Message %d from chan3", i):
				time.Sleep(time.Millisecond * 1500)
			}
		}
	}()

	// Close done channel after 6 seconds
	go func() {
		time.Sleep(time.Second * 6)
		close(done)
	}()

	for msg := range mergedChannel {
		fmt.Println(msg)
	}
	fmt.Println("Fan-in process completed")
}
