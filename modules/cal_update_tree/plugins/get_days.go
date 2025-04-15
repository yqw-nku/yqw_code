package plugins

import (
	"active/common"
	"active/config"
	"fmt"
)

// GetDays 用户指定礼包搭配，计算出完成的天数以及对应的消耗
func GetDays(conf config.Config) {
	// 把升级所需的天数转为瓶子数
	var needBottleSum = (conf.NeedDays * 24 * 60) / common.BottleEqualTime

	// 额外的瓶子
	needBottleSum -= conf.ExtraBottle

	// 每天免费获取的瓶子
	var freeBottleSum = common.GetEveryDayFreeBottle()

	// 获取当前礼包搭配，每天可以获取的瓶子数、消耗
	var everyDayBottle = 0
	var everyDayCost = 0
	for _, gift := range conf.GiftArray {
		everyDayBottle += gift.Value()
		everyDayCost += gift.Cost()
	}
	fmt.Printf("需要的瓶子：%d，每天免费可以获取瓶子数：%d，礼包每天可以获取瓶子数：%d\n", needBottleSum, freeBottleSum, everyDayBottle)

	// 死循环，直到瓶子数达到目标
	var days = 0
	for {
		if needBottleSum-freeBottleSum-everyDayBottle <= 0 {
			days += 1
			break
		}
		days += 1
		needBottleSum = needBottleSum - freeBottleSum - everyDayBottle
	}

	// 打印消耗和天数
	fmt.Printf("实际升级天数：%d，消耗：%d\n", days, everyDayCost*days)
}
