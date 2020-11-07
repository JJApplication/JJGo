/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package entity

type JJGoPubTest struct {
	Client_IP string `json:"client_ip"`
	Method    string `json:"method"`
	Host      string `json:"host"`
	Content_Type string `json:"content_type"`
}

type JJGoStatus struct {
	PID   int `json:"pid"`
	Port  int `json:"port"`
	Count int `json:"count"`
}

type JJGoDemo struct {
	Rest  map[string]string `json:"rest"`
	File  map[string]string `json:"file"`
	Html  map[string]string `json:"html"`
}
