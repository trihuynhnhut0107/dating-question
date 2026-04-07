FROM golang:alpine

WORKDIR /app

# Install air for hot-reloading
RUN go install github.com/air-verse/air@latest

# Copy go.mod and go.sum for dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Expose port and start air
EXPOSE 8080
CMD ["air", "-c", ".air.toml"]
