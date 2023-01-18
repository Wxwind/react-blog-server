FROM golang:1.19.5
ENV GOPROXY https://goproxy.cn
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go build -o ./apps/article/articleEXE ./apps/article/article.go
WORKDIR apps/article
CMD ["./articleEXE"]
