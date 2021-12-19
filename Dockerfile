FROM golang:1.17-alpine AS builder

RUN mkdir /app
ADD ./src /app

WORKDIR /app

RUN go mod download
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -o server .
CMD ["./server"]
