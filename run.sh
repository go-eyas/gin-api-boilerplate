#!/bin/sh


set -e
#set -x

SCRIPTPATH=$(cd `dirname $0`; pwd)
cd $SCRIPTPATH

export GOPROXY=https://goproxy.cn

APPNAME="server"
APPPORT="9000"

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
    if [ $1 == "upx" ];then
      echo "下载安装: https://github.com/upx/upx/releases"
    fi
    exit 1
  fi
}

build_docker() {
  # git pull origin master
  set +e
  docker rm -f $APPNAME
  docker rmi $APPNAME:latest
  set -e
  docker build -t $APPNAME:latest .
  docker run --name $APPNAME -tid \
    -p $APPPORT:$APPPORT \
    -v `pwd`/runtime:/opt/app/runtime \
    -v `pwd`/api.local.toml:/opt/app/api.local.toml \
    $APPNAME:latest
}

build_bin() {
  green_text "[1/3]buildding assets..."
  go generate ./assets
  green_text "[1/3]build assets successfull!"

  green_text "[2/3]go build..."
  GitCommit=$(git rev-list -1 HEAD)
  BuildTime=$(date "+%F　%T")
  goVersion=$(go version)
  go build -o $APPNAME -ldflags "-s -w -X main.gitCommit=\"$GitCommit\" -X main.buildTime=\"$BuildTime\" " main.go
  green_text "[2/3]build successfull!"
  
  # 如果安装了 upx 就压缩一下
  if command -v "upx" > /dev/null;then 
    green_text "[3/3]you are installed upx, compress..."
    upx_tmp="__upx_tmp"
    rm -rf $upx_tmp
    upx --best -o $upx_tmp $APPNAME
    mv -f $upx_tmp $APPNAME
    green_text "[3/3]compress successfull!"
  else 
    green_text "[3/3]upx not install, skip compress!"
  fi
  green_text "BUILD OK!!!"
}

help_demo() {
    echo -e "
USAGE:
  $0 <command>

COMMANDS: 
  build                     编译项目
  docker                    使用 docker 运行项目
  docs                      根据 apidoc 注释生成文档
  help                      打印帮助信息
"
}

case "$1" in
  "docker")
    go generate ./assets
    build_docker $@
    ;;
  "build")
    build_bin
    ;;
  "docs")
    mkdir -p assets/docs
    cd assets/docs
    apidoc --single -i ../../ -o .
    green_text "文档已生成, 请至 assets/docs 目录查看 "  
    go generate ./assets
    ;;
  *)
    help_demo
    ;;

esac;
