package golanglearning

import (
	"testing"
	"time"
)

func TestNewMessageQueueServer(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: Config{
				BufferSize:  5,
				WorkerCount: 3,
				RetryCount:  3,
				RetryDelay:  time.Second,
			},
			wantErr: false,
		},
		{
			name: "invalid buffer size",
			config: Config{
				BufferSize:  0,
				WorkerCount: 3,
				RetryCount:  3,
				RetryDelay:  time.Second,
			},
			wantErr: true,
		},
		{
			name: "invalid worker count",
			config: Config{
				BufferSize:  5,
				WorkerCount: 0,
				RetryCount:  3,
				RetryDelay:  time.Second,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, err := NewMessageQueueServer(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMessageQueueServer() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && server == nil {
				t.Error("NewMessageQueueServer() returned nil server with no error")
			}
		})
	}
}

func TestMessageQueueServer_StartAndShutdown(t *testing.T) {
	config := DefaultConfig()
	server, err := NewMessageQueueServer(config)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	// Send some messages
	for i := 0; i < 3; i++ {
		err := server.SendMessage("test message")
		if err != nil {
			t.Errorf("Failed to send message: %v", err)
		}
	}

	// Shutdown server
	if err := server.Shutdown(); err != nil {
		t.Errorf("Failed to shutdown server: %v", err)
	}
}

func TestMessageQueueServer_ConcurrentAccess(t *testing.T) {
	config := DefaultConfig()
	server, err := NewMessageQueueServer(config)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	// Test concurrent message sending
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			err := server.SendMessage("concurrent message")
			if err != nil {
				t.Errorf("Failed to send message: %v", err)
			}
			done <- true
		}(i)
	}

	// Wait for all messages to be sent
	for i := 0; i < 10; i++ {
		<-done
	}

	if err := server.Shutdown(); err != nil {
		t.Errorf("Failed to shutdown server: %v", err)
	}
}

func TestMessageQueueServer_RaceConditions(t *testing.T) {
	config := DefaultConfig()
	server, err := NewMessageQueueServer(config)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	if err := server.Start(); err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	// Test concurrent access to dropped messages
	for i := 0; i < 100; i++ {
		go func(i int) {
			server.SendMessage("race test message")
		}(i)
	}

	
	go server.Shutdown()

	time.Sleep(time.Millisecond * 100)
}
