package main

import (
	"bubble/models"
	"bubble/routers"
	"bubble/server/global"
	"bubble/server/viper"
	"fmt"
)

func main() {
	//if len(os.Args) < 2 {
	//	fmt.Println("Usage：./bubble config/config.ini")
	//	return
	//}

	global.GVA_VP = viper.Viper() // 初始化Viper,从config读取配置

	// 加载配置文件
	//if err := setting.Init("./config/config.ini"); err != nil {
	//	fmt.Printf("load config from file failed, err:%v\n", err)
	//	return
	//}
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := viper.InitMySQL()
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer viper.Close() // 程序退出关闭数据库连接
	// 模型绑定
	global.GVA_DB.AutoMigrate(&models.User{})
	// 注册路由
	r := routers.SetupRouter()

	r.Run(":8080")
}
