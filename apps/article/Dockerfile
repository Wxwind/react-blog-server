FROM golang:latest
COPY .. /app
WORKDIR /app
RUN go build -o article article.go
CWD ["/app/article"]