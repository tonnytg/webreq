# WebReq Examples

This directory contains example programs demonstrating the usage and performance features of the WebReq library.

## Performance Demo

The `performance_demo.go` file demonstrates:

1. **Simple GET Request** - Basic usage of the library
2. **Sequential Requests** - Shows connection reuse across multiple requests
3. **Concurrent Requests** - Demonstrates connection pooling benefits
4. **Custom Context with Timeout** - Using ExecuteWithContext for timeout control
5. **Request Cancellation** - Programmatic cancellation using context

### Running the Performance Demo

```bash
cd examples
go run performance_demo.go
```

**Note**: The demo makes real HTTP requests to httpbin.org, so an internet connection is required.

### Expected Output

```
=== WebReq Performance Demo ===

1. Simple GET Request:
   Request completed in 234ms
   Status code: 200

2. Sequential Requests (connection reuse):
   Request 1 completed (Status: 200)
   Request 2 completed (Status: 200)
   Request 3 completed (Status: 200)
   Request 4 completed (Status: 200)
   Request 5 completed (Status: 200)
   Total time for 5 requests: 876ms
   Average per request: 175ms

3. Concurrent Requests (connection pooling):
   Completed 10 concurrent requests in 245ms
   Average per request: 24ms

4. Custom Context with Timeout:
   Request completed in 2.1s
   Status code: 200

5. Request Cancellation:
   Cancelling request...
   Request cancelled after 501ms
   Error: context canceled
```

### Key Observations

- **Connection Reuse**: Sequential requests are faster than creating new connections each time
- **Connection Pooling**: Concurrent requests are significantly faster (10x+) due to parallel execution with pooled connections
- **Context Control**: ExecuteWithContext provides fine-grained control over request lifecycle
- **Cancellation**: Requests can be cancelled programmatically, saving resources

## Creating Your Own Examples

Feel free to add more examples demonstrating:
- Integration with specific APIs
- Error handling patterns
- Advanced header usage
- Custom timeout strategies
- Retry logic
- Rate limiting

All examples should follow Go best practices and include clear comments.
