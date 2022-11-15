FROM golang:alpine

WORKDIR /app
COPY . .
RUN go build -o main -ldflags="-s -w" main.go
CMD ["./main"]