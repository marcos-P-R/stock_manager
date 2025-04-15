FROM golang:1.24-alpine AS builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o service_product ./cmd/server/main.go


FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/service_product .

EXPOSE 8080

CMD ["./service_product"]
