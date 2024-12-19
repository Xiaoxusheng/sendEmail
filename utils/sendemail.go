package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"time"
)

type Tem struct {
	Nickname string `json:"nickname"`
	Message  string `json:"message"`
	Birthday string `json:"birthday"`
}

func SendEmail(c *Birthday) error {
	//读取html
	// 解析模板文件
	//if err != nil {
	//	log.Println(err.Error())
	//	return err
	//}
	//t := Tem{
	//	Nickname: c.Nickname,
	//	Message:  c.Message,
	//	Birthday: c.Birthday,
	//}
	//
	//// 执行模板，将data中的数据注入到HTML模板中，并写入响应
	//
	//err = tmpl.Execute(open, t)
	//if err != nil {
	//	log.Println(err.Error())
	//	return err
	//}

	e := email.NewEmail()
	//发送者
	e.From = "<" + c.Username + ">"
	//接收者
	e.To = []string{c.Email}
	//主题
	e.Subject = c.Title
	//html
	e.HTML = []byte("<!DOCTYPE html>\n<html lang=\"zh\">\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n    <title>生日祝福</title>\n    <style>\n        body {\n            font-family: 'Arial', sans-serif;\n            margin: 0;\n            padding: 0;\n            background-color: #faf1e6;\n        }\n        .email-container {\n            width: 100%;\n            max-width: 600px;\n            margin: 0 auto;\n            background-color: #ffffff;\n            border-radius: 10px;\n            box-shadow: 0 0 15px rgba(0, 0, 0, 0.1);\n        }\n        .email-header {\n            background-color: #ff7f50;\n            padding: 20px;\n            text-align: center;\n            color: #ffffff;\n            font-size: 26px;\n            font-weight: bold;\n            border-top-left-radius: 10px;\n            border-top-right-radius: 10px;\n        }\n        .email-body {\n            padding: 30px;\n            color: #333;\n            font-size: 16px;\n            line-height: 1.6;\n        }\n        .email-body h2 {\n            color: #ff7f50;\n            font-size: 24px;\n            margin-bottom: 10px;\n        }\n        .email-body p {\n            margin: 15px 0;\n        }\n        .email-footer {\n            background-color: #f8f8f8;\n            padding: 20px;\n            text-align: center;\n            font-size: 14px;\n            color: #777;\n            border-bottom-left-radius: 10px;\n            border-bottom-right-radius: 10px;\n        }\n        .btn {\n            background-color: #ff7f50;\n            color: #ffffff;\n            text-decoration: none;\n            padding: 12px 25px;\n            border-radius: 5px;\n            font-size: 16px;\n            display: inline-block;\n            margin-top: 20px;\n        }\n        .highlight {\n            color: #ff6347;\n            font-weight: bold;\n        }\n    </style>\n</head>\n<body>\n<div class=\"email-container\">\n    <div class=\"email-header\">\n        生日快乐，<span class=\"highlight\">" + c.Nickname + "</span>！\n    </div>\n    <div class=\"email-body\">\n        <h2> " + c.Nickname + "</h2>\n<p>" + c.Message + "</p>\n        <p>再次祝你生日快乐！🎂</p>\n        <a href=\"#\" class=\"btn\">让我们一起庆祝！</a>\n    </div>\n    <div class=\"email-footer\">\n        <p>祝一切顺利，开心每一天。</p>\n        <p><span class=\"highlight\"></span></p>\n    </div>\n</div>\n</body>\n</html>")

	fmt.Println(c)

	err := e.SendWithStartTLS("smtp.qq.com:587", smtp.PlainAuth("", c.Username, c.Password, "smtp.qq.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.gmail.com:465"})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	log.Println("发送成功！")
	return nil
}

func Loop(c *Config) {
	t1 := time.Now()
	t2 := time.Date(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute()+1, 0, 0, t1.Location())
	log.Println("任务启动," + t2.Sub(t1).String() + "后开始执行")

	t3 := time.NewTicker(t2.Sub(t1))
	defer t3.Stop()
	for {
		select {
		case <-t3.C:
			log.Println("开始执行任务")
			go Lk(c)
			//一天以后执行
			t1 = time.Now()
			t2 = time.Date(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute()+1, 0, 0, t1.Location())
			log.Println("任务启动," + t2.Sub(t1).String() + "后开始执行")
		}
	}
}

func Lk(c *Config) {
	fmt.Println(c.Group)
	for _, res := range c.Group {
		t := time.Now()
		t2, err := time.Parse(time.DateOnly, res.Birthday)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(t2.Sub(t))
		if t2.Sub(t).Hours() > (-24) && t2.Sub(t).Hours() < 0 {
			//	今天有人过生日
			err := SendEmail(&Birthday{
				Username: c.Username,
				Password: c.Password,
				Email:    res.Email,
				Title:    res.Title,
				Message:  res.Message,
				Birthday: res.Birthday,
				Nickname: res.Nickname,
			})
			if err != nil {
				log.Println(err)
			}
			return
		}
		log.Println("今天没人过生日！")
	}
}
