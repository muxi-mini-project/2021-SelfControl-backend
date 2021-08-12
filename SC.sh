#!/usr/bin/env bash

port=2333
filename=main
chmod a+x ./${filename}
nohup ./${filename} &
str=$(lsof -i:${port} | grep ${filename})
i=0
for st in $str
do
    let "i++"
    if [ $i == 2 ] # pid是str里的第二个元素
    then
        pid=$st
        kill $pid
        chmod a+x ./${filename}
        nohup ./${filename} &
        echo 'ok'
        exit 0
    fi
done
