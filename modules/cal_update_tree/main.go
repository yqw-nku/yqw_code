package main

import (
	"active/config"
	"fmt"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		fmt.Printf("加载配置失败，%s", err.Error())
		return
	}

	// 根据模式进行处理
	switch config.Conf.Mode {
	case "day":
		return
	case "gift":
		return
	default:
		fmt.Printf("模式错误")
		return
	}
}
