#!/usr/bin/env bash

# 将main.go 编译成第一个参数名 上传到服务器并ssh连接服务器

ip=39.99.53.8
filename=$1

go build -o bin/${filename} main.go 
scp bin/${filename} root@${ip}:/root/2021-SelfControl-backend
ssh root@${ip}