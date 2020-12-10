#!/usr/bin/env bash

# 默认不会打包因为zip程序不一定会存在
# 默认不会清理掉缓存 使用--clean清除

function is_root()
{
  root=$(id -u)
  if [ ! "$root" -eq 0 ]
  then
    echo "请在root用户下运行此脚本构建"
    exit 1
  fi
}

function build_jjgo()
{
  echo "检查GOPATH"
  echo GOPATH="$GOPATH"
  echo "检查cgo编译环境"
  echo CGO_ENABLED="$(go env | grep CGO_ENABLED)"
  echo "编译环境$(go env |grep GOOS) $(go env |grep GOARCH)"
  echo "开始编译JJGo程序"
  export GOOS=linux
  export ARCH=amd64
  go build -ldflags "-s -w" ./src/jjgo.go
  if [ ! -f "./jjgo" ];then
    echo "jjgo程序编译失败"
  else
    echo "jjgo编译成功"
    sha
  fi
}

function build_jjcli()
{
  echo "开始编译JJCLI工具"
  go build -ldflags "-s -w" ./jjtools/jjcli/jjcli.go
  if [ ! -f "./jjcli" ];then
    echo "jjcli程序编译失败"
  else
    echo "jjcli编译成功"
  fi
}

function build_jjlog()
{
  echo "开始编译JJLog工具"
  go build -ldflags "-s -w" ./jjtools/jjlog/jjlog.go
  if [ ! -f "./jjlog" ];then
    echo "jjlog程序编译失败"
  else
    echo "jjlog编译成功"
  fi
}

function generate_pkg()
{
  echo "生成服务包..."
  if [ ! -d "./pkg_jjgo" ];then
    echo "目录不存在即将创建"
    mkdir ./pkg_jjgo
  else
    mkdir ./pkg_jjgo || echo "目录已存在"
  fi

  chmod +x ./jjgo
  chmod +x ./jjcli
  chmod +x ./jjlog

  echo "修改程序执行权限"
  cp ./jjgo ./pkg_jjgo
  cp ./jjcli ./pkg_jjgo
  cp ./jjlog ./pkg_jjgo

  cp -r ./conf ./pkg_jjgo
  cp -r ./script ./pkg_jjgo
  cp -r ./lib ./pkg_jjgo
  cp -r ./static ./pkg_jjgo
  cp -r ./swagger ./pkg_jjgo
  cp app_define.json ./pkg_jjgo

  mkdir -p ./pkg_jjgo/logs
  touch ./pkg_jjgo/logs/jjgo.log
  touch ./pkg_jjgo/logs/jjgo.pid

  cp jjgo_build.log ./pkg_jjgo
  cp jjgo.sha256 ./pkg_jjgo
  # 解决可能出现的文件夹权限问题
  echo "移动程序到打包路径"
}

function zip_pkg()
{
  if [ -n "$1" ]
  then
    echo "开始打包"
    zip -r jjgo.zip ./pkg_jjgo
    echo "打包完毕"
  fi
}

function clean_cache()
{
  if [ -n "$1" ]
  then
    echo "清理缓存文件..."
    if [ -d "./pkg_jjgo" ];then
      rm -rf ./pkg_jjgo
      rm -f ./jjgo
      rm -f ./jjcli
      rm -f ./jjlog
      rm -f ./jjgo.sha256
      rm -f ./jjgo_build.log
      echo "清理完毕"
    fi
  fi
}

function build_log()
{
  echo "生成build日志"
  touch jjgo_build.log
  date=$(date)
  echo "build date: ${date}" > jjgo_build.log
}

function sha()
{
  echo "生成jjgo sha256校验码"
  sha256sum ./jjgo > jjgo.sha256
}

echo "使用--zip开启打包"
echo "打包时使用--clean开启清理缓存"

umask 022
is_root
build_jjgo
build_jjcli
build_jjlog
build_log
generate_pkg
zip_pkg "$1"
clean_cache "$2"

exit