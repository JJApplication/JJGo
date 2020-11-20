#!/usr/bin/env bash

function build_jjgo()
{
  echo "检查GOPATH"
  echo $GOPATH
  echo "开始编译JJGo程序"
  export GOOS=linux
  export ARCH=amd64
  # 针对macos的环境配置
  export GOPATH=/Users/landers/code/go
  go build -ldflags "-s -w" ./src/jjgo.go
  if [ ! -f "./jjgo" ];then
    echo "jjgo程序编译失败"
  else
    echo "jjgo编译成功"
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
    echo "目录已存在"
  fi

  chmod +x ./jjgo
  chmod +x ./jjcli
  chmod +x ./jjlog

  echo "修改程序执行权限"
  mv ./jjgo ./pkg_jjgo
  mv ./jjcli ./pkg_jjgo
  mv ./jjlog ./pkg_jjgo

  cp -r ./conf ./pkg_jjgo
  cp -r ./script ./pkg_jjgo
  cp -r ./lib ./pkg_jjgo
  cp -r ./static ./pkg_jjgo
  cp -r ./swagger ./pkg_jjgo
  cp app_define.json ./pkg_jjgo

  mkdir -p ./pkg_jjgo/logs
  touch ./pkg_jjgo/logs/jjgo.log
  touch ./pkg_jjgo/logs/jjgo.pid

  mv jjgo_build.log ./pkg_jjgo
  # 解决可能出现的文件夹权限问题
  echo "移动程序到打包路径"
  echo "文件权限默认为Mac用户 请在Linux上修改用户组"

  echo "开始打包"
  zip -r jjgo.zip ./pkg_jjgo
  echo "打包完毕"
}

function clean_cache()
{
  echo "因为权限问题，需要手动清理缓存文件..."
}

function build_log()
{
  echo "生成build日志"
  touch jjgo_build.log
  echo "build date: $(date)" > jjgo_build.log
}


build_jjgo
build_jjcli
build_jjlog
build_log
generate_pkg
clean_cache

exit