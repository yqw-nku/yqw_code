package config

import (
	"github.com/AlecAivazis/survey/v2"
)

type Config struct {
	PaperUrl string `json:"paper_url,omitempty"`
}

var Conf Config

var expectDaysPrompt = &survey.Input{
	Message: "文章链接：",
}

func LoadConfig() error {
	var conf Config

	// 询问文章链接
	if err := survey.AskOne(expectDaysPrompt, &conf.PaperUrl, survey.WithValidator(survey.Required)); err != nil {
		return err
	}

	Conf = conf
	return nil
}
