# Go Message Queue Server

A robust message queue server implementation in Go that provides configurable message processing, graceful shutdown, and concurrent message handling.

## Project Structure

```
.
├── main.go                    # Main application entry point
└── GolangLearning/
    ├── MessageQueueServer.go  # Core message queue server implementation
    └── MessageQueueServer_test.go  # Test suite for the server
```

## Features

- Configurable message buffer size
- Multiple worker goroutines for concurrent message processing
- Graceful shutdown mechanism
- Message retry mechanism with configurable attempts and delay
- Thread-safe message handling
- Unprocessed message tracking
- Dropped message tracking during shutdown

## Core Components

### MessageQueueServer

The main server implementation (`GolangLearning/MessageQueueServer.go`) provides:

- **Configuration Management**
  ```go
  type Config struct {
      BufferSize  int
      WorkerCount int
      RetryCount  int
      RetryDelay  time.Duration
  }
  ```

- **Key Methods**
  - `NewMessageQueueServer(config Config)`: Creates a new server instance
  - `Start()`: Initializes worker goroutines
  - `SendMessage(message string)`: Sends a message with retry logic
  - `Shutdown()`: Gracefully shuts down the server
  - `PrintUnprocessedMessages()`: Prints messages that weren't processed
  - `GetDroppedMessages()`: Returns list of dropped messages

### Main Application

The main application (`main.go`) demonstrates:

- Server initialization
- Concurrent message sending
- Graceful shutdown handling
- Message processing monitoring

## Usage Example

```go
func main() {
    // Create server with default configuration
    config := DefaultConfig()
    server, err := NewMessageQueueServer(config)
    if err != nil {
        log.Fatal(err)
    }

    // Start the server
    server.Start()

    // Send messages concurrently
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            server.SendMessage(fmt.Sprintf("user_%d", i))
        }(i)
    }

    // Shutdown after delay
    time.Sleep(time.Second * 2)
    server.Shutdown()

    // Wait for all messages to be processed
    wg.Wait()

    // Print dropped messages
    dropped := server.GetDroppedMessages()
    if len(dropped) > 0 {
        fmt.Println("Dropped messages:", dropped)
    }
}
```

## Configuration

Default configuration values:
```go
{
    BufferSize:  5,        // Message buffer size
    WorkerCount: 3,        // Number of worker goroutines
    RetryCount:  3,        // Number of retry attempts
    RetryDelay:  time.Second, // Delay between retries
}
```

## Error Handling

The server implements comprehensive error handling:
- Configuration validation
- Message sending errors
- Worker initialization errors
- Shutdown errors

## Thread Safety

- Mutex protection for dropped message access
- Safe concurrent message sending
- Proper channel closure handling
- WaitGroup synchronization

## Best Practices

1. **Graceful Shutdown**
   - Proper channel closure
   - Worker cleanup
   - Message tracking

2. **Resource Management**
   - Buffered channels
   - Configurable buffer sizes
   - Worker pool pattern

3. **Error Handling**
   - Validation
   - Error propagation
   - Proper cleanup

## Testing

The project includes comprehensive tests:
- Unit tests
- Concurrent access tests
- Shutdown tests
- Configuration validation tests

## Performance Considerations

- Buffered channels for message queuing
- Worker pool for concurrent processing
- Configurable retry mechanism
- Efficient message tracking

## Dependencies

- Standard library only
- No external dependencies required

## License

This project is open source and available under the MIT License. 