# Product Search API

A Go-based REST API for searching pharmaceutical products using Elasticsearch. This service provides fuzzy search capabilities across product names, generic drugs, and companies with relevance scoring.

## Features

- **Elasticsearch Integration**: Leverages Elasticsearch for fast, scalable searchvariables

## Prerequisites

- Go 1.24.4 or later
- Elasticsearch instance (Cloud)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/baihakhi/product-search.git
cd product-search
```

2. Install dependencies:

```bash
go mod download
```

3. Set up environment variables by creating a `.env` file:

```env
ES_CLOUD_ID=your_elasticsearch_cloud_id
ES_API_KEY=your_elasticsearch_api_key
ES_INDEX=products
```

## Configuration

The application uses the following environment variables:

- `ES_CLOUD_ID`: Elasticsearch Cloud deployment ID
- `ES_API_KEY`: Elasticsearch API key for authentication
- `ES_INDEX`: Name of the Elasticsearch index (default: products)

## Building

### Local Build

```bash
go build -o app ./cmd
```

### Docker Build

```bash
docker build -t product-search .
```

## Running

### Dev Run

```bash
go run ./cmd
```

### Local Run

```bash
./app
```

The server will start on `http://localhost:8080`

### Docker Run

```bash
docker run -p 8080:8080 --env-file .env product-search
```

## API Documentation

### Search Products

Search for products using a query string.

**Endpoint:** `GET /search`

**Query Parameters:**

- `q` (required): Search query string
- `size` (optional): Number of results to return (default: 20)

**Example Request:**

```bash
curl "http://localhost:8080/search?q=paracetamol&size=10"
```

**Postman Collection**
`https://pf56v4a4-3763763.postman.co/workspace/Postpartum's-Workspace~fa37146d-0df7-4649-84d1-5157e2f98b1d/collection/51495705-a7a17800-72c4-4a94-94de-63a19289ec96?action=share&creator=51495705`

## Project Structure

```
.
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── elastic_search.go   # Elasticsearch configuration
│   ├── handler/
│   │   └── product_handler.go  # HTTP handlers
│   ├── model/
│   │   └── product_model.go    # Data models
│   ├── repository/
│   │   └── product_repo.go     # Data access layer
│   ├── response/
│   │   └── json_wrapper.go     # Response utilities
│   ├── service/
│   │   └── product_service.go  # Business logic
│   └── util/
│       └── string.go           # Utility functions
├── tes/
│   └── test.go                 # Test utilities
├── Dockerfile                  # Docker configuration
├── go.mod                      # Go module file
└── README.md                   # This file
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
