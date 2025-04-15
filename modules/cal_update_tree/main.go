package main

import (
	"active/config"
	"active/plugins"
	"fmt"
	"time"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		fmt.Printf("加载配置失败，%s", err.Error())
		time.Sleep(60 * time.Second)
		return
	}

	// 根据模式进行处理
	switch config.Conf.Mode {
	case "day":
		plugins.GetGift(config.Conf)
	case "gift":
		plugins.GetDays(config.Conf)
	default:
		fmt.Printf("模式错误")
		time.Sleep(60 * time.Second)
		return
	}

	time.Sleep(60 * time.Second)
	return
}
