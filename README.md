# 🔍 API Benchmark

This project performs benchmarks on public APIs, measuring response times, HTTP status, and handling multiple requests. Ideal for testing the performance of endpoints.

## 🚀 How to Run with Docker

### 1. Clone the repository

```bash
git clone https://github.com/daviolvr/api-benchmark.git
cd api-benchmark
```

### 2. Build the Docker image
```bash
docker build -t api-benchmark .
```

This command creates a Docker image named api-benchmark.

### 3. Run the container
```bash
docker run -p 8080:8080 api-benchmark
```

The service will be available at http://localhost:8080.

### 4. Available Endpoints
- GET / - Welcome message
- POST /benchmark - Starts the API benchmark
- GET /results - Retrieves benchmark results
- DELETE /results/clear - Clears saved results

### 5. Example curl Request
```bash
curl -X POST "http://localhost:8080/benchmark?url=https://api.github.com&count=3"
```

