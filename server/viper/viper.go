package viper

import (
	"bubble/server/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper
// 设置默认值
// 从JSON、TOML、YAML、HCL、envfile和Java properties格式的配置文件读取配置信息
// 实时监控和重新读取配置文件（可选）
// 从环境变量中读取
// 从远程配置系统（etcd或Consul）读取并监控配置变化
// 从命令行参数读取配置
// 从buffer读取配置
// 显式配置值
func Viper() *viper.Viper {

	v := viper.New()
	v.SetConfigFile(ConfigDefaultFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
