# Define 'builder' stage
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# CGO_ENABLED=0 - Disable C dependencies to generate a static Go binary
# GOOS=linux - Specifies that the binary must be compiled in Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/apple_backend

# Import an alpine container
FROM alpine:3.18

# Create new folder in alpine: alpine/root
WORKDIR /root/

# Copies the binary from 'builder' stage, and moves it into alpine
COPY --from=builder /app/apple_backend .

# Give execute permissions
RUN chmod +x apple_backend

EXPOSE 8080
CMD ["./apple_backend"]