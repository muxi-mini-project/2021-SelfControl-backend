#!/usr/bin/env bash

ip=39.99.53.8
filename=$1

go build -o ${filename} main.go 
scp ./${filename} root@${ip}:/root/2021-SelfControl-backend
ssh root@${ip}