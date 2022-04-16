package emailsender

import (
	bytes2 "bytes"
	"email/config"
	"email/structs"
	"fmt"
	"github.com/jordan-wright/email"
	"html/template"
	"log"
	"net/smtp"
	"time"
)

func RenderHtml(ex structs.ActivationEx) ([]byte, error) {
	timeLayout := "2006-01-02 15:04:05"
	t1, err := template.ParseFiles("./template.html")
	if err != nil {
		panic(err)
	}
	var s bytes2.Buffer
	month := ex.UseTimeSpan / (30 * 24 * 3600)
	renderMap := make(map[string]interface{})
	renderMap["AdminName"] = ex.AdminName
	renderMap["UserName"] = ex.UserName
	renderMap["EditionType"] = ex.EditionType
	renderMap["Expired"] = month
	renderMap["SupportLangList"] = ex.SupportLangList
	renderMap["AdminName"] = ex.AdminName
	renderMap["Mark"] = ex.Mark
	renderMap["Sn"] = ex.Sn
	renderMap["CreateDate"] = time.Unix(ex.CreatedAt, 0).Add(time.Hour * 8).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	err = t1.Execute(&s, renderMap)
	if err != nil {
		return nil, err
	}
	return s.Bytes(), nil
}

func SendEmail(ex structs.ActivationEx) error {
	// 简单设置 log 参数
	var err error
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "Matrix翻译后台系统<matrix_trans_sn@126.com>"
	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To, err = config.ParseConfigEmails()
	if err != nil {
		return err
	}
	// 设置主题
	var editionName string
	if ex.EditionType == "Test" {
		editionName = "测试版"
	} else {
		editionName = "正式版"
	}
	em.Subject = fmt.Sprintf("【%s】生成了新的【%s】激活码", ex.AdminName, editionName)
	// 简单设置文件发送的内容，暂时设置成纯文本
	//em.Text = []byte(ex.Sn)
	em.HTML, err = RenderHtml(ex)
	if err != nil {
		return err
	}
	//设置服务器相关的配置
	err = em.Send("smtp.126.com:25", smtp.PlainAuth("", "matrix_trans_sn@126.com", "RNNJYBFOUPYXJEZB", "smtp.126.com"))
	if err != nil {
		return err
	}
	log.Println("send successfully ... ")
	return nil
}
