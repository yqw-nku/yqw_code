package config

import (
	"active/common"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

type Config struct {
	Mode        string            `json:"mode"`
	ExtraBottle int               `json:"extra_bottle"`
	NeedDays    int               `json:"need_days"`
	ExpectDays  int               `json:"expect_days"`
	Days        int               `json:"days"`
	GiftArray   []common.GiftType `json:"gift_array"`
}

var Conf Config

var modePrompt = &survey.Select{
	Message: "请选择模式：",
	Options: []string{"day", "gift"},
	Description: func(value string, index int) string {
		if value == "day" {
			return "指定达成天数"
		}

		return "指定礼包搭配"
	},
}

var expectDaysPrompt = &survey.Input{
	Message: "期望达成的天数：",
}

var giftPrompt = &survey.MultiSelect{
	Message: "请选择礼包搭配：",
	Options: []string{common.Gift1.Name(), common.Gift2.Name(), common.Gift3.Name(), common.Gift4.Name(), common.Gift5.Name()},
}

var extraBottlePrompt = &survey.Input{
	Message: "额外的瓶子数量：",
}

var leftDaysPrompt = &survey.Input{
	Message: "还要多少天升级：",
}

func LoadConfig() error {
	var conf Config

	// 选择模式
	if err := survey.AskOne(modePrompt, &conf.Mode, survey.WithValidator(survey.Required)); err != nil {
		return err
	}

	// 根据模式进行处理
	switch conf.Mode {
	case "day":
		// 询问期望达成的天数
		if err := survey.AskOne(expectDaysPrompt, &conf.ExpectDays, survey.WithValidator(survey.Required)); err != nil {
			fmt.Printf("解析失败，%s", err.Error())
			return err
		}
	case "gift":
		// 选择礼包搭配
		if err := survey.AskOne(giftPrompt, &conf.GiftArray, survey.WithValidator(survey.Required)); err != nil {
			fmt.Printf("解析失败，%s", err.Error())
			return err
		}
	default:
		return fmt.Errorf("模式错误")
	}

	// 公共参数
	// 询问是否有额外的瓶子数量
	if err := survey.AskOne(extraBottlePrompt, &conf.ExtraBottle, survey.WithValidator(survey.Required)); err != nil {
		return err
	}

	// 询问当前需要的天数
	if err := survey.AskOne(leftDaysPrompt, &conf.NeedDays, survey.WithValidator(survey.Required)); err != nil {
		return err
	}

	Conf = conf
	return nil
}
