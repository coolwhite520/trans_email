package config

import (
	"github.com/Unknwon/goconfig"
	"log"
	"strings"
)

func ParseConfigEmails() ([]string,error) {
	cfg, err := goconfig.LoadConfigFile("./emails.ini")
	if err != nil {
		log.Println("加载配置失败:" + err.Error())
		return nil, err
	}
	emailStr, err := cfg.GetValue("emails", "emails")
	if err != nil {
		log.Println("获取emails字段信息失败:" + err.Error())
		return nil, err
	}
	arr := strings.Split(emailStr, ",")
	var emails []string
	for _, item := range arr {
		e := strings.Trim(item, " ")
		if  len(e) > 0{
			emails = append(emails, e)
		}
	}
	return emails, nil
}