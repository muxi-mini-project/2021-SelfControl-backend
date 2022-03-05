
FROM golang:1.17-alpine
ENV GOPROXY=https://goproxy.cn
WORKDIR /build
COPY . .
EXPOSE 2333
RUN mkdir /app
RUN  go mod tidy
RUN go build -o /app main.go
WORKDIR /app
CMD ["/app/main"]
