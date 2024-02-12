# Dockerfile for Golang backend
FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 5000

CMD ["./main"]
