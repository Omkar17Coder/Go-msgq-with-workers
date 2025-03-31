package golanglearning

import (
	"fmt"
	"sync"
	"time"
)

func producer(userch chan<- string, start, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i < start+count; i++ {
		time.Sleep(time.Second * 2)
		userch <- fmt.Sprintf("user_%d", i)
	}
	close(userch)
}

func consumer(userch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	free:
	for {
		select {
		case message, ok := <-userch:
			if !ok {
				break free
			}
			fmt.Printf("Recived on the channel %s", message)
		default:
			time.Sleep(time.Second*1)
		}
		fmt.Println("All are sleeping")
	}
}

func SimulateChannels() {
	userch := make(chan string, 3)
	var wg sync.WaitGroup

	wg.Add(1)
	go producer(userch, 0, 3, &wg)

	wg.Add(1)
	go consumer(userch, &wg)

	wg.Wait()
}
