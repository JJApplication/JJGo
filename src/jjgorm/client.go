/*
App: jjgo
Author: Landers
Copyright: Landers1037 renj.io
Github: https://github.com/landers1037

和sqlite3数据库交互的引擎
通过调用.Error()获取数据库的错误信息
*/

package jjgorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"jjgo/src/logger"
)

type JJGorm struct {
	// 初始化内部注册一个DB结构体
	dbc *gorm.DB
}

func JJGormClient() JJGorm {
	return JJGorm{}
}

// 连接数据库
//
// db: 需要连接的数据库路径
func (d *JJGorm) Connect(db string) *gorm.DB {
	dbc, err := gorm.Open("sqlite3", db)
	if err != nil {
		logger.JJGoLogger.Error("数据库连接失败", err)
		return dbc
	}
	d.dbc = dbc
	// 每次数据库使用完毕应该关闭
	return dbc
}

func (d *JJGorm) Close() *gorm.DB {
	_ = d.dbc.Close()
	return d.dbc
}
// 自动迁移
func (d *JJGorm) AutoMigrate(table interface{}) *gorm.DB {
	d.dbc.AutoMigrate(table)
	return d.dbc
}

// 新建表
func (d *JJGorm) CreateTable(table interface{}) *gorm.DB {
	// 首先判断表是否存在
	if !d.dbc.HasTable(table) {
		return d.dbc.CreateTable(table)
	}
	return d.dbc
}

// 删除表
func (d *JJGorm) DeleteTable(table interface{}) *gorm.DB {
	return d.dbc.DropTableIfExists(table)
}

// 插入值
func (d *JJGorm) Insert(table interface{}) *gorm.DB {
	return d.dbc.Create(table)
}

// 删除值, 根据条件查询
func (d *JJGorm) Delete(table interface{}, con ...interface{}) *gorm.DB {
	return d.dbc.Delete(table, con)
}

// 计数
func (d *JJGorm) Count(table interface{}) *gorm.DB {
	return d.dbc.Count(table)
}

// 查询表，根据结构体查询
func (d *JJGorm) FindFirst(table interface{}, where ...interface{}) *gorm.DB {
	return d.dbc.First(table, where...)

}

// 查询表，根据结构体查询
func (d *JJGorm) FindLast(table interface{}, where ...interface{}) *gorm.DB {
	return d.dbc.Last(table, where...)
}

// 查询全部，根据条件
func (d *JJGorm) FindAll(table interface{}, where ...interface{}) *gorm.DB {
	return d.dbc.Find(table, where...)
}

// 条件查询
func (d *JJGorm) FindBy(table interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return d.dbc.Where(query, args...).Find(table)
}

func (d *JJGorm) FindFirstBy(table interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return d.dbc.Where(query, args...).First(table)
}

// 模糊查询 仅仅是传入的condition不同 内部逻辑一致
func (d *JJGorm) FindLike(table interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return d.dbc.Where(query, args...).Find(table)
}

func (d *JJGorm) FindFirstLike(table interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return d.dbc.Where(query, args...).First(table)
}

func (d *JJGorm) FindIn(table interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return d.dbc.Where(query, args...).Find(table)
}

func (d *JJGorm) FindFirstIn(table interface{}, query interface{}, args ...interface{}) *gorm.DB {
	return d.dbc.Where(query, args...).First(table)
}

// 通过结构体查询
func (d *JJGorm) FindByStruct(condition interface{}, table interface{}) *gorm.DB {
	return d.dbc.Where(condition).Find(table)
}

// 通过结构体查询
func (d *JJGorm) FindFirstByStruct(condition interface{}, table interface{}) *gorm.DB {
	return d.dbc.Where(condition).First(table)
}

// 排除
func (d *JJGorm) Not(condition interface{}, table interface{}) *gorm.DB {
	return d.dbc.Not(condition).Find(table)
}

// 根据struct查询or
func (d *JJGorm) Or(con1 interface{}, con2 interface{}, table interface{}) *gorm.DB {
	return d.dbc.Where(con1).Or(con2).Find(table)
}

// 保存数据
func (d *JJGorm) Save(table interface{}) *gorm.DB {
	return d.dbc.Save(table)
}

// 通过结构体更新属性值
func (d *JJGorm) Update(table interface{}, new interface{}) *gorm.DB {
	return d.dbc.Model(table).Updates(new)
}

// 使用原生的gorm映射
func (d *JJGorm) Raw() *gorm.DB {
	return d.dbc
}

func (d *JJGorm) Commit() *gorm.DB {
	return d.dbc.Commit()
}

func (d *JJGorm) RollBack() *gorm.DB {
	return d.dbc.Rollback()
}