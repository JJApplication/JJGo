# JJGO

<a href="https://goreportcard.com/report/github.com/landers1037/dirichlet"><img src="./copyright/goreport.svg" /></a>
<a href="http://service.renj.io"><img src="./copyright/renj.io.svg"/></a>
<a href="https://github.com/JJApplication"><img src="./copyright/copyright-JJService.svg"/></a>



## what

JJGo是一个高性能的Rest接口框架

支持cluster模式运行，开启多实例
支持多数据库同步

## 更新日志

详见changelog.json

- 最新版 v2.0

- v2.1 修复数组越界 修正python脚本

    当前问题：jjmail的接口请求取消了jjauth校验，但是存在refer和host校验，取消订阅是不能正常返回的

