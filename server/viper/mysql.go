package viper

import (
	"bubble/server/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitMySQL() (err error) {
	cfg := global.GVA_CONFIG.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)

	global.GVA_DB, err = gorm.Open("mysql", dsn)
	global.GVA_DB.LogMode(true) //启动调试模式，控制台会输出对应的SQL语句
	if err != nil {
		return
	}

	return global.GVA_DB.DB().Ping()
}

func Close() {
	global.GVA_DB.Close()
}
