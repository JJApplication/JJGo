#!/bin/zsh
# only for my linux docker
# 因为glibc版本原因使用golang1.12.1镜像构建

GOPATH=/Users/landers/code/go
name=go
port=8080
echo "docker编译 挂载路径$GOPATH 标准golang容器内path为/go"
echo "暴露端口$port 仅针对Linux Windows有效"
echo "别名：$name"
echo "正在启动docker容器"

if [ "$(docker attach $name)" != $name ];then
  echo "容器未初始化或者未启动"
  docker stop $name
  if [ "$(docker start $name)" != $name ];then
    echo "未创建容器实例"
    docker run --name=$name -it -v $GOPATH/src:/go/src -p $port:$port golang:1.12.1 /bin/bash
  else
    echo "启动容器中"
    docker attach $name
  fi
fi
