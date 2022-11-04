/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package main

import (
	"bubble/server/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"testing"
)

func TestRedis(t *testing.T) {
	db := redis.NewClient(&redis.Options{
		Addr:     "81.68.182.112:6379",
		Password: "qwertyuiop",
	})
	ctx := context.Background()
	//db.Do(ctx, "set", "k1", "v1")

	res, err := db.Do(ctx, "get", "k1").Result()

	if err != nil {
		if err == redis.Nil {
			fmt.Println("redis 键值对不存在")
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("result:", res)
	}
}

func TestYaml(t *testing.T) {

	v := viper.New()
	v.SetConfigFile("config.yaml")
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
	fmt.Println(global.GVA_CONFIG.JWT.SigningKey)
}

func TestMul(t *testing.T) {
	names := []string{"a", "b", "c", "d"}
	for k, v := range names {
		t.Run(v, func(t *testing.T) {
			fmt.Println(k, ":", v)
		})
	}
}
