/**
 * @Author: DollarKiller
 * @Description: 邮件相关
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 14:45 2019-09-12
 */
package gemail

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

const html = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Message}}</title>
</head>
<body>
    <h4>亲爱的{{.ToUserName}},您好！</h4>
    <div>{{.Message}}</div>
    </br>
    <div>
        {{.FromUserName}} </br>
        {{.TimeDate}}
    </div>
</body>
</html>`


// 发送通用 email 通知  (这是一个公用账户)

func SendNifoLog(mailTo []string,subject string, body string) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string {
		"user": "notice@dollarkiller.com",
		"pass": "%Y4I4qjlqKAy",
		"host": "smtp.mail.ru",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From","EasyUtils Notice" + "<" + mailConn["user"] + ">")  //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)  //发送给多个用户
	m.SetHeader("Subject", subject)  //设置邮件主题
	m.SetBody("text/html", body)     //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
//func SendMail(fromUser, toUser, subject string) error {
//	// NewEmail返回一个email结构体的指针
//	e := email.NewEmail()
//	// 发件人
//	e.From = fromUser
//	// 收件人(可以有多个)
//	e.To = []string{toUser}
//	// 邮件主题
//	e.Subject = subject
//	// 解析html模板
//	t,err := template.New("data").Parse(html)
//	if err != nil {
//		return err
//	}
//	// Buffer是一个实现了读写方法的可变大小的字节缓冲
//	body := new(bytes.Buffer)
//	// Execute方法将解析好的模板应用到匿名结构体上，并将输出写入body中
//	t.Execute(body,struct {
//		FromUserName string
//		ToUserName   string
//		TimeDate     string
//		Message      string
//	}{
//		FromUserName: "go语言",
//		ToUserName:   "Sixah",
//		TimeDate:     time.Now().Format("2006/01/02"),
//		Message:      "golang是世界上最好的语言！",
//	})
//
//	e.Text = []byte(" this is test file")
//	// html形式的消息
//	e.HTML = body.Bytes()
//	// 从缓冲中将内容作为附件到邮件中
//	//e.Attach(body, "email-template.html", "text/html")
//	// 以路径将文件作为附件添加到邮件中
//	// 发送邮件(如果使用QQ邮箱发送邮件的话，passwd不是邮箱密码而是授权码)
//	return e.Send("smtp.mail.ru:465", smtp.PlainAuth("", "notice@dollarkiller.com", "%Y4I4qjlqKAy", "smtp.mail.ru"))
//}
//

