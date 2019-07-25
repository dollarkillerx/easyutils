package main

import (
	"fmt"
	"github.com/dollarkillerx/easyutils"
	"log"
)

// captcha 验证 需要用户输入 , __test 无法做到这一点 就这样吧

func main() {
	//testnum()
	testmp3()
}

func testnum() {
	captchaId, base64Png := easyutils.CaptchaNum()
	log.Println(captchaId)
	log.Println(base64Png)
	var id string
	fmt.Scan(&id)
	b, e := easyutils.CaptchaCheck(captchaId, id)
	if e != nil {
		panic(e.Error())
	}
	log.Println(b)
}

func testmp3() {
	captchaId, base64MP3 := easyutils.CaptchaMP3()
	log.Println(captchaId)
	log.Println(base64MP3)
}
