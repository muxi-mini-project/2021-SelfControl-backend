#!/usr/bin/env bash

# 杀死旧的服务 将名为第一个参数的文件部署在本地

port=2333
filename=$1

cd /root/2021-SelfControl-backend
chmod a+x ./${filename}
nohup ./${filename} &
str=$(lsof -i:${port})
i=0
for st in $str
do
    # echo $st
    let "i++"
    if [ $i == 11 ] # pid是str里的第11个元素
    then
        pid=$st
        kill $pid
        chmod a+x ./${filename}
        nohup ./${filename} &
        echo 'ok'
        exit 0
    fi
done
