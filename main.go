package main

import (
	"fmt"
	"sync"
	"time"
)

type Server struct {
	userCh         chan string
	quitCh         chan struct{}
	wg             sync.WaitGroup
	droppedMessage []string
}

func NewServer() *Server {
	return &Server{
		userCh:         make(chan string, 5), // Buffered channel with size 5
		quitCh:         make(chan struct{}),
		droppedMessage: []string{},
	}
}

func (s *Server) StartWorker() {

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		for {
			select {
			case message, ok := <-s.userCh:
				if !ok {
					fmt.Println("Worker exiting: Channel closed.")
					return
				}
				fmt.Println("Processing:", message)
				time.Sleep(2*time.Second)
			case <-s.quitCh:
				fmt.Println("Server shutting down, worker exiting.")
				return
			}
		}
	}()
}

func (s *Server) SendMessage(message string) {

	select {
	case s.userCh <- message:
		fmt.Println("Sent message:", message)
	case <-s.quitCh:
		s.droppedMessage = append(s.droppedMessage, message)

		fmt.Println("Server is shutting down. Message dropped:", message)
	}
}

func (s *Server) Shutdown() {

	close(s.quitCh)

	close(s.userCh)

	s.wg.Wait()
	
	fmt.Println("Server has shut down.")
}

func (s *Server) Start(workers int) {

	for i := 0; i < workers; i++ {
		s.StartWorker()
	}
}

func (s *Server) PrintUnprocessedMessages() {

	if len(s.userCh) > 0 {

		for message := range s.userCh {
			fmt.Println(" message was queued but  not processed", message)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	server := NewServer()
	server.Start(3)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			server.SendMessage(fmt.Sprintf("user_%d", i))
		}(i)
	}

	go func() {
		time.Sleep(time.Second * 2)
		server.Shutdown()
	}()

	wg.Wait()

	server.wg.Wait()
	if len(server.droppedMessage) > 0 {
		fmt.Println("The following messages were not processed due to shutdown:")
		for _, msg := range server.droppedMessage{
			fmt.Println(msg)
		}
	}

	server.PrintUnprocessedMessages()

	fmt.Println("Main function has finished.")
}
