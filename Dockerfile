# Use official Go base image (latest stable)
FROM golang:1.21

# Set working directory inside container
WORKDIR /app

# Copy module files first (to leverage cache)
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy entire project
COPY . .

# Build Go app
RUN go build -o main .

# Expose the port the app listens on
EXPOSE 8000

# Run the app
CMD ["./main"]
