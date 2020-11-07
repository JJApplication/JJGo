/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037
*/
package database

// API 访问次数
// 当前统计全部接口
type JJGoAPICount struct {
	Id    int   	`gorm:"column:id;AUTO_INCREMENT"`
	API   string    `gorm:"column:api" json:"api"`
	Count int		`gorm:"column:count" json:"count"`
}

func (JJGoAPICount) TableName() string {
	return "jjgo_api_count"
}