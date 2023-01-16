FROM golang:latest
COPY .. /app
WORKDIR /app
RUN go build -o fileServer fileServer.go
CWD ["/app/fileServer"]