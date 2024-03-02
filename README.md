## Splitter Backend

Backend REST APIs service built with Go to support API functionalities for Splitter application.

### Technologies

* [Go](https://go.dev) - Go programming language 
* [Gin](https://gin-gonic.com/) - Go web framework to build HTTP servers
* [Gorm]() - 
* [Postgres]() -
* [Docker]() - 

### Getting Started

NOTE: instructions for running with Docker will be added in the future

1. Install Go
   
   Make sure Go 1.16 or above is installed on your machine. Visit Go [download page](https://go.dev/dl/) to download or install using Homebrew.

   ```bash
   brew install go
   ```

2. Install Dependencies
   
   ```bash
   go mod download
   ```

3. Run the app

    ```bash
    go run .
    ```

### Testing

Run unit tests on the app

```bash
go test 
```

### Database

Running the database on your machine using Docker Compose.

```bash
docker compose up -d postgres
```