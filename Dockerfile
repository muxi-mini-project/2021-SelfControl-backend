FROM golang:1.17-alpine
ENV GOPROXY=https://goproxy.cn
RUN mkdir /app 
ADD . /app
EXPOSE 2333

WORKDIR /app
RUN go mod tidy
RUN go build -o /app main.go
CMD ["/app/main"]
