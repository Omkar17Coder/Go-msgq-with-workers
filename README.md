# Go Message Queue Server

A robust message queue server implementation in Go that provides configurable message processing, graceful shutdown, and concurrent message handling.

## Project Structure

```
.
├── main.go                    # Main application entry point
└── GolangLearning/
    ├── MessageQueueServer.go  # Core message queue server implementation
    ├── MessageQueueServer_test.go  # Test suite for the server
    ├── Concurrency.go        # Basic concurrency examples
    └── TwoFanInWithGracefulShutdown.go  # Fan-in pattern with graceful shutdown
```

## Features

- Configurable message buffer size
- Multiple worker goroutines for concurrent message processing
- Graceful shutdown mechanism
- Message retry mechanism with configurable attempts and delay
- Thread-safe message handling
- Unprocessed message tracking
- Dropped message tracking during shutdown
- Fan-in pattern implementation
- Two-stage message processing pipeline

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

### Concurrency Examples
Basic concurrency patterns demonstrating:
- Channel usage
- Goroutine synchronization
- WaitGroup patterns
- Buffered vs unbuffered channels

### TwoFanInWithGracefulShutdown
Implements a two-stage fan-in pattern with:
- Multiple input channels
- Graceful shutdown handling
- Message aggregation
- Error handling

## Usage Examples

### Basic Message Queue
```go
config := DefaultConfig()
server, err := NewMessageQueueServer(config)
if err != nil {
    log.Fatal(err)
}
server.Start()
server.SendMessage("test message")
server.Shutdown()
```

### Fan-In Pattern
```go
func main() {
    // Create input channels
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    // Create output channel
    out := make(chan string)
    
    // Start fan-in
    go FanIn(ch1, ch2, out)
    
    // Send messages
    ch1 <- "Message 1"
    ch2 <- "Message 2"
    
    // Read results
    fmt.Println(<-out)
    fmt.Println(<-out)
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