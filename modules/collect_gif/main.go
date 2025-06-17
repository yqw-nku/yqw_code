package main

import (
	"collect/config"
	"collect/plugins"
	"fmt"
	"time"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		fmt.Printf("加载配置失败，%s", err.Error())
		time.Sleep(60 * time.Second)
		return
	}

	plugins.DownloadGifFromPaper(config.Conf.PaperUrl)

	time.Sleep(60 * time.Second)
	return
}
