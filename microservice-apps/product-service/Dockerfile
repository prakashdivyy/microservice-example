# Start from the latest golang base image
FROM golang:alpine AS build-env

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Download upx
RUN apk add --no-cache gcc musl-dev upx ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -ldflags="-s -w" server.go && upx -9 -q ./server

# Multi stage
FROM scratch

COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /app/server /

# Command to run
ENTRYPOINT ["/server"]
