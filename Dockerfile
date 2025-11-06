# Build stage
FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Copy source code
COPY . .

# Download dependencies (if any)
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o history-slackbot cmd/bot/main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

# Set timezone (optional, defaults to UTC)
ENV TZ=America/New_York

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/history-slackbot .

# Run the application
CMD ["./history-slackbot"]
