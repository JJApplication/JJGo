/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

jjmail的body结构体
*/
package entity

// 基本结构体 邮箱和发送信息
type JJMailBody struct {
	Mail_Address string `json:"mail_address"`
	Message      string `json:"message"`
}

// 订阅和取消订阅
type JJMailSub struct {
	Mail_Address string `json:"mail_address"`
	// 需要订阅/取消的服务名
	Service      string `json:"service"`
}

// axios传递的data信息
type JJMailAddress struct {
	Address string `json:"address"`
}

// axios传递的data信息
type JJMailAddressData struct {
	Address string `json:"address"`
	// data是传入的数据json字符串
	Data    string `json:"data"`
}