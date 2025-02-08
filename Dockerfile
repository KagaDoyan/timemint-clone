# Build stage
FROM golang:1.23.5-alpine AS builder
RUN apk add --no-cache git

WORKDIR /go/src/app
COPY *.mod *.go *.yaml *.json ./

RUN go mod download
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./...

# Final stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata nano

# Create necessary directories
RUN mkdir -p /logs

# Copy application binary and configuration file
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/app/config.yaml /config.yaml

# Set timezone
ENV TZ=Asia/Bangkok

# Entrypoint for the container
ENTRYPOINT ["/app"]

# Metadata and port exposure
LABEL Name=my_app
EXPOSE 3000
