/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-23
* Time: 下午3:44
* */
package easyutils

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"sync"
)

// 设置CaptchaConfig类
type CaptchaConfig struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

// 创建CaptchaConfig类实例
var (
	captchaConfig     *CaptchaConfig
	captchaConfigOnce sync.Once
)

// 获取base64验证码基本配置
func GetCaptchaConfig() *CaptchaConfig {
	captchaConfigOnce.Do(func() {
		captchaConfig = &CaptchaConfig{
			Id:          "",
			CaptchaType: "character",
			VerifyValue: "",
			ConfigAudio: base64Captcha.ConfigAudio{},
			ConfigCharacter: base64Captcha.ConfigCharacter{
				Height:             60,
				Width:              240,
				Mode:               2,
				IsUseSimpleFont:    false,
				ComplexOfNoiseText: 0,
				ComplexOfNoiseDot:  0,
				IsShowHollowLine:   false,
				IsShowNoiseDot:     false,
				IsShowNoiseText:    false,
				IsShowSlimeLine:    false,
				IsShowSineLine:     false,
				CaptchaLen:         0,
			},
			ConfigDigit: base64Captcha.ConfigDigit{},
		}
	})
	return captchaConfig
}

//@获取验证码
func CaptchaGeneratePng() (captchaId, base64Png string) {
	captchaConfig := GetCaptchaConfig()
	// 创建base64图像验证码
	character := captchaConfig.ConfigCharacter
	// /GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	captchaId, digitCap := base64Captcha.GenerateCaptcha("", character)

	// 生成图片
	base64Png = base64Captcha.CaptchaWriteToBase64Encoding(digitCap)
	return
}

//@验证验证码是否正确
// captchaId verifyValue 用户输入
func CaptchaCheck(captchaId, verifyValue string) (bool, error) {
	verifyResult := base64Captcha.VerifyCaptcha(captchaId, verifyValue)
	if verifyResult {
		return true, nil
	} else {
		return false, fmt.Errorf("captcha is error")
	}
}

//@简单案例
//func GenerateTest(ctx *gin.content) {
//	// get session
//	session := sessions.Default(c)
//	captchaConfig := GetCaptchaConfig()
//	// 创建base64图像验证码
//	character := captchaConfig.ConfigCharacter
//	// /GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
//	captchaId, digitCap := base64Captcha.GenerateCaptcha(captchaConfig.Id, character)
//
//	// 生成图片
//	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)
//	session.Set("captchaId",captchaId)
//}

//@验证 验证码是否输入正确
// captchaId 存在session 中
// verifyValue: 客户端发来的验证码
//func VerfiyCaptcha(captchaId, verifyValue string) (int, error){
//	verifyResult := base64Captcha.VerifyCaptcha(captchaId, verifyValue)
//	if verifyResult {
//		return CAPTCHA_IS_RIGHT,nil
//	} else {
//		return CAPTCHA_IS_ERROR, fmt.Errorf("captcha is error")
//	}
//}

////@获取验证码
//func CaptchaGeneratePng() (captchaId, base64Png string) {
//	//数字验证码配置
//	var configD = base64Captcha.ConfigDigit{
//		Height:     80,
//		Width:      240,
//		MaxSkew:    0.7,
//		DotCount:   80,
//		CaptchaLen: 5,
//	}
//
//	//创建数字验证码.
//	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
//	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
//	//以base64编码
//	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
//	return idKeyD,base64stringD
//}
