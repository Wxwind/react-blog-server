FROM golang:latest
COPY .. /app
WORKDIR /app
RUN go build -o article article.go
CMD ["/app/article"]