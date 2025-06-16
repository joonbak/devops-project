# Stage 1: Build Go Binary
FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main .

# Stage 2: Minimal Image
FROM scratch
COPY --from=builder /app/main /main
ENTRYPOINT ["/main"]
