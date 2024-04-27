# Use the official Go image as base
FROM golang:latest AS build

# Set the working directory in the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build Go binary, name it main, deposit it in current dir
# GOOS env var influences build to generate a linux compatible binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Serve the application with a lightweight image
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=build /app/main .

# Expose port 8080
EXPOSE 8080

# Command to run the Go application
CMD ["./main"]
