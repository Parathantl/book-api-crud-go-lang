# Book API – GoLang CRUD App

A RESTful API built in GoLang for managing books. It supports full CRUD operations and keyword-based search, with data persisted in a simple JSON file.

---

## Features

- Create, Read, Update, Delete (CRUD) operations
- Search books by `title` or `description`
- JSON file used for persistence (`books.json`)
- Uses Go's concurrency features in search (goroutines + channels)
- Fully Dockerized

---

## Project Structure

book-api/ 
├── main.go 
├── book.go 
├── storage.go 
├── handlers.go 
├── books.json 
├── go.mod 
├── go.sum 
├── Dockerfile 
└── .dockerignore

---

## Requirements

- Go 1.21+ (for local development)
- Docker installed (for containerized deployment)

---

## Docker Usage

### Build the Docker Image

```bash
docker build -t book-api .
```

## Run with Persistent JSON File

To make sure books.json changes are saved outside the container:

On macOS/Linux:
```bash
docker run -p 8000:8000 -v $(pwd)/books.json:/app/books.json book-api
```

On Windows (PowerShell):
```bash
docker run -p 8000:8000 -v ${PWD}/books.json:/app/books.json book-api
```

### Run the Container (without persistence)
```bash
docker run -p 8000:8000 book-api
```

---

- Data is stored in books.json
- Search is case-insensitive and matches both title and description
- App runs on port 8000 by default