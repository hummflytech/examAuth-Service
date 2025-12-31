FROM golang:1.25 AS builder

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

RUN go build -o bin cmd/main.go

# FROM alpine

# WORKDIR /app

# COPY --from=builder /app/bin /app/bin

EXPOSE 8080

EXPOSE 50051

CMD ["/app/bin"]


