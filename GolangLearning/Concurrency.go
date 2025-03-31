package golanglearning

import (
	"fmt"
	"sync"
	"time"
)

func FetchUserData(userID int, respCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(80 * time.Millisecond)
	respCh <- "UserData"
}

func FetchUserRecommendations(userID int, respCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(120 * time.Millisecond)
	respCh <- "UserRecommendations"
}

func FetchUserLikes(userID int, respCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(50 * time.Millisecond)
	respCh <- "UserLikes"
}

func SimulateConcurrency() {
	// Create a buffered channel
	respCh := make(chan string, 3) 
	wg := &sync.WaitGroup{}

	now := time.Now()
	userID := 10

	wg.Add(3)

	
	go FetchUserData(userID, respCh, wg)
	go FetchUserRecommendations(userID, respCh, wg)
	go FetchUserLikes(userID, respCh, wg)


	wg.Wait()


	close(respCh)

	
	for {
		message, ok := <-respCh
		if !ok {
			fmt.Println("Channel closed, no more messages")
			break
		}
		fmt.Println(message)
	}

	fmt.Printf("Total time taken: %v\n", time.Since(now))
}
