# Point of Sales API

This is a Point of Sales API built with Go.

## Prerequisites

- Go (version 1.18 or higher)
- PostgreSQL (or your configured database)
- (Optional) Docker & Docker Compose

## Running Locally

1. **Clone the repository:**

   ```
   git clone https://github.com/alpredovandy/point-of-sales-api.git
   cd point-of-sales-api
   ```

2. **Install dependencies:**

   ```
   go mod download
   ```

3. **Run the application:**

   ```
   go run main.go
   ```

   The API should now be running at `http://localhost:8080`.

## Using Docker (Optional)

1. **Build and run with Docker Compose:**
   ```
   docker-compose up --build
   ```

## License

MIT
