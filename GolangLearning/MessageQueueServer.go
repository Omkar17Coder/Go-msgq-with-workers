package golanglearning

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Config struct {
	BufferSize  int
	WorkerCount int
	RetryCount  int
	RetryDelay  time.Duration
}

func DefaultConfig() Config {
	return Config{
		BufferSize:  5,
		WorkerCount: 3,
		RetryCount:  3,
		RetryDelay:  time.Second,
	}
}

type MessageQueueServer struct {
	userCh         chan string
	quitCh         chan struct{}
	wg             sync.WaitGroup
	droppedMessage []string
	config         Config
	mu             sync.Mutex
	userChClosed   bool
}

func NewMessageQueueServer(config Config) (*MessageQueueServer, error) {
	if config.BufferSize <= 0 {
		return nil, errors.New("buffer size must be positive")
	}
	if config.WorkerCount <= 0 {
		return nil, errors.New("worker count must be positive")
	}
	if config.RetryCount < 0 {
		return nil, errors.New("retry count cannot be negative")
	}

	return &MessageQueueServer{
		userCh:         make(chan string, config.BufferSize),
		quitCh:         make(chan struct{}),
		droppedMessage: make([]string, 0),
		config:         config,
		userChClosed:   false,
	}, nil
}

func (s *MessageQueueServer) StartWorker() error {
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
				time.Sleep(2 * time.Second)
			case <-s.quitCh:
				fmt.Println("Server shutting down, worker exiting.")
				return
			}
		}
	}()
	return nil
}

func (s *MessageQueueServer) SendMessage(message string) error {
	retryCount := 0
	for {
		s.mu.Lock()
		if s.userChClosed {
			s.mu.Unlock()
			return fmt.Errorf("channel is closed, message dropped: %s", message)
		}
		s.mu.Unlock()

		select {
		case s.userCh <- message:

			fmt.Println("Sent message:", message)
			return nil
		case <-s.quitCh:
			s.mu.Lock()
			s.droppedMessage = append(s.droppedMessage, message)
			s.mu.Unlock()
			return fmt.Errorf("server is shutting down, message dropped: %s", message)
		default:
			if retryCount < s.config.RetryCount {
				retryCount++
				fmt.Printf("Retrying to send message (attempt %d/%d)...\n", retryCount, s.config.RetryCount)
				time.Sleep(s.config.RetryDelay)
			} else {
				return fmt.Errorf("message dropped due to retry limit: %s", message)
			}
		}
	}
}

func (s *MessageQueueServer) Shutdown() error {
	close(s.quitCh)
	s.wg.Wait()     
	s.mu.Lock()
	if !s.userChClosed {
		close(s.userCh)
		s.userChClosed = true
	}
	s.mu.Unlock()

	fmt.Println("Server has shut down.")
	return nil
}

func (s *MessageQueueServer) Start() error {
	for i := 0; i < s.config.WorkerCount; i++ {
		if err := s.StartWorker(); err != nil {
			return fmt.Errorf("failed to start worker %d: %w", i, err)
		}
	}
	return nil
}

func (s *MessageQueueServer) PrintUnprocessedMessages() {
	if len(s.userCh) > 0 {
		for message := range s.userCh {
			fmt.Println("Message was queued but not processed:", message)
		}
	}
}

func (s *MessageQueueServer) GetDroppedMessages() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.droppedMessage
}

func (s *MessageQueueServer) GetConfig() Config {
	return s.config
}
