# URL Shortener written in Go.

## Features:
- Shortens long URLs into compact, short URLs
- Stores URL mappings in Redis (for retrieval, if needed)
- RESTful API to create short URLs and retrieve long URLs
- Serves a basic frontend

## Tech:
- Go
- Gin for web framework
- Redis as main data store
- HTML for frontend

### Running locally:

- Ensure Go 1.18+ and Redis (either locally or via Docker) is installed.

1. Clone this repository
2. Start Redis (either locally or via Docker)
3. Run the Go server: 'go run main.go'
4. Open your browser @ http://localhost:9808
5. Enter a long URL and a short URL will be created