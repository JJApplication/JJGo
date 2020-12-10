#!/usr/bin/env bash
# 控制服务的service脚本
# 可交互的sh脚本 替代之前的jjctl工具，纯shell实现jjservice的管理

# 定义颜色
end="\033[0m"
red="\033[31m"
green="\033[32m"
yellow="\033[33m"
blue="\033[34m"

function help()
{
  echo -e "${green}JJService tools.Type help for more... ${end}"
  echo "Usage: type --help or help."
  echo "1: start jjservice"
  echo "2: stop jjservice"
  echo "3: restart jjservice"
  echo "4: status jjservice"
  echo "5: clear nginx cache"
}

function read_arg()
{
  echo -ne "${green}Input nums to do: ${end}"
  read arg
  case $arg in
  1)
    echo -e "${blue}准备启动jjservice服务${end}"
    if [ ! -f jjservice ];then
      echo -e "${red}找不到启动文件 ${end}"
      return
    fi
    pid=$(pgrep jjservice)
    if [ -n "$pid" ];then
      echo -e "${green}jjservice已经运行, PID: ${pid}${end}"
      return
    fi
    nohup ./jjservice > /dev/null 2>&1 &
    pid=$(pgrep jjservice)
    echo -e "${green}jjservice启动完成, PID: ${pid}${end}"
    exit 0
    ;;
  2)
    echo -e "${blue}准备停止jjservice服务${end}"
    pid=$(pgrep jjservice)
    kill "$pid"
    echo -e "${green}jjservice停止完成${end}"
    ;;
  3)
    echo -e "${blue}准备重启jjservice服务${end}"
    pid=$(pgrep jjservice)
    if [ -z $pid ];then
      echo -e "${red}jjservice没有在运行${end}"
    else
      if [ ! -f jjservice ];then
        echo -e "${red}找不到启动文件${end}"
      else
        kill $pid
        nohup ./jjservice > /dev/null 2>&1 &
        echo -e "${green}jjservice重启完成${end}"
      fi
    fi
    ;;
  4)
    echo -e "${blue}查看jjservice服务状态${end}"
    pid=$(pgrep jjservice)
    if [ ! -f ./conf/production.ini ];then
      conf="无法获取配置文件"
    else
      conf=$(cat ./conf/production.ini)
    fi
    if [ -z $pid ];then
      echo -e "${red}jjservice没有在运行${end}"
    else
      echo -e "${green}jjservice运行正常${end} PID: ${blue}${pid}${end}"
      echo -e "${yellow}配置文件如下:${end}"
      echo -e "${green}${conf}${end}"
    fi
    ;;
  5)
    echo -e "${blue}清空nginx缓存${end}"
    if [ ! -d /home/nginx/cache ];then
      echo -e "${red}找不到nginx缓存路径${end}"
    else
      cd /home/nginx/cache && rm -rf *
      echo -e "${green}清空nginx缓存完成${end}"
      exit 0
    fi
  esac
}

help
read_arg