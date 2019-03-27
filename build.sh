#!/bin/bash

set -e
# set -x

SCRIPTPATH=$(cd `dirname $0`; pwd)
cd $SCRIPTPATH

APPNAME="api"

red_text() {
  echo -e "\e[31m$*\e[0m"
}

green_text() {
  echo -e "\e[32m$*\e[0m"
}

# 检查命令
check_cmd_or_exit() {
  if ! command -v $1 > /dev/null;then
    red_text "未安装 $1"
    if [ $1 == "go" ];then
      echo "下载安装: https://studygolang.com/dl"
    fi

    if [ $1 == "docker" ];then
      echo "windows 下载安装: https://store.docker.com/editions/community/docker-ce-desktop-windows"
      echo "linux 命令安装: curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun"
    fi

    if [ $1 == "gqlgen" ];then
      echo "命令安装: GO111MODULE=off go get -u -v github.com/99designs/gqlgen"
    fi

    
    exit 1
  fi
}


help_demo() {
    echo -e "
USAGE:
  ./run.sh <command>

COMMANDS:
  docker [version]            构建项目成 docker 镜像
  build                       构建适用于当前操作系统的可执行文件
  dev                         启动开发环境，自动重启，自动生成文档
  help                        打印帮助信息
"
}

build_docker() {
  version=$1
  version=${version:-"1.0"}
  echo $version
  docker run --rm \
    -e GOPROXY=http://goproxy.io \
    -v `pwd`:/usr/src/app \
    -w /usr/src/app \
    golang:1.12-alpine \
    go build -v -o build/$APPNAME -ldflags "-s" main.go
  
  docker build -t $APPNAME:$version build/
}

check_cmd_or_exit go

case "$1" in
  "docker")
    build_docker
  ;;
  "dev")
    cd build
    go run ../main.go
    ;;
  "build")
    go build -v -o build/$APPNAME -ldflags "-s" main.go
    green_text "build/$APPNAME: build successfully!"
    ;;
  *)
    help_demo
    ;;
esac