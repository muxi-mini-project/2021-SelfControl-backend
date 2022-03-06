#创建一个基于5.7.30版本的MySql
FROM mysql:5.7.30 

EXPOSE 3306

#设置免密登录
ENV MYSQL_ALLOW_EMPTY_PASSWORD yes

#将所需文件放到容器中
#拷贝安装脚本
COPY /mysql/setup.sh /mysql/setup.sh 
#创建数据库
COPY /mysql/create_db.sql /mysql/create_db.sql
#初始数据
COPY /mysql/initial_data.sql /mysql/initial_data.sql
#设置密码和权限
COPY /mysql/privileges.sql /mysql/privileges.sql 

# #设置容器启动时执行的命令
CMD ["sh", "/mysql/setup.sh"]


FROM golang:1.17-alpine
ENV GOPROXY=https://goproxy.cn
RUN mkdir /app 
ADD . /app
EXPOSE 2333

WORKDIR /app
RUN go mod tidy
RUN go build -o /app main.go
CMD ["/app/main"]
