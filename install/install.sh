#!/bin/bash

unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     machine=Linux;;
    Darwin*)    machine=Mac;;
    CYGWIN*)    machine=Cygwin;;
    MINGW*)     machine=MinGw;;
    *)          machine="UNKNOWN:${unameOut}"
esac

echo  "当前操作系统：" ${machine}

if [ "$machine" == "Mac" ]; then
   env GOOS=darwin GOARCH=amd64 go build -mod=vendor -o http-server-mac main.go
   mv http-server-mac  /usr/local/bin/http-server
fi

if [ "$machine" == "Linux" ]; then
   env GOOS=linux GOARCH=amd64 go build -mod=vendor -o http-server-linux main.go
   mv http-server-linux  /usr/local/bin/http-server
fi

echo "安装成功,运行 http-server 测试服务是否安装成功"