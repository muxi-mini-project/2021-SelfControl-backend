#!/usr/bin/env bash

# 将main.go 编译成main 上传到服务器并ssh连接服务器

ip=39.99.53.8
filename=main

go build -o ${filename} main.go 
scp ${filename} root@${ip}:/root/2021-SelfControl-backend
rm ${filename}
ssh root@${ip}