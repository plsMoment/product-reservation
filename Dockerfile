FROM golang:1.22-alpine

WORKDIR /usr/local/src

# dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# build
COPY ./ ./
RUN go build -o backend-service ./cmd/main.go