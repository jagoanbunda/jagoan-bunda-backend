# Build stage
FROM golang:1.25.5-alpine AS builder

# Install git dan ca-certificates (diperlukan untuk fetch dependencies)
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod dan sum files
COPY go.mod go.sum ./

# Download dependencies (cached jika go.mod/go.sum tidak berubah)
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build aplikasi dengan cache mount untuk Go build cache
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/main .

# Final stage - menggunakan distroless untuk image yang lebih kecil dan aman
FROM alpine:latest

# Install ca-certificates dan timezone data
RUN apk --no-cache add ca-certificates tzdata

# Set timezone
ENV TZ=Asia/Jakarta

WORKDIR /app

# Copy binary dari builder stage
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Run aplikasi
CMD ["./main"]
