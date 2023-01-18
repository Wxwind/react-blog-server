FROM golang:1.19.5
ENV GOPROXY https://goproxy.cn
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build -o ./fileServer/fileServerEXE ./fileServer/main.go
WORKDIR fileServer
CMD ["./fileServerEXE"]