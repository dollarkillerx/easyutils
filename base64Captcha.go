/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-23
* Time: 下午3:44
* */
package easyutils

// 此处以来开源库 base64Captcha
// 作者: https://github.com/mojocn/base64Captcha

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
)

//config struct for digits
//数字验证码配置
var configD = base64Captcha.ConfigDigit{
	Height:     80,
	Width:      240,
	MaxSkew:    0.7,
	DotCount:   80,
	CaptchaLen: 5,
}

//config struct for audio
//声音验证码配置
var configA = base64Captcha.ConfigAudio{
	CaptchaLen: 6,
	Language:   "zh",
}

//config struct for Character
//字符,公式,验证码配置
var configC = base64Captcha.ConfigCharacter{
	Height: 60,
	Width:  240,
	//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
	Mode:               base64Captcha.CaptchaModeNumber,
	ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
	ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
	IsShowHollowLine:   false,
	IsShowNoiseDot:     false,
	IsShowNoiseText:    false,
	IsShowSlimeLine:    false,
	IsShowSineLine:     false,
	CaptchaLen:         6,
}

//@ 生成数字验证码
func CaptchaNum() (captchaId, base64Png string) {
	//创建数字验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	return idKeyD, base64stringD
}

//@ 生成公式验证码
func CaptchaMath() (captchaId, base64Png string) {
	//创建字符公式验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)
	return idKeyC, base64stringC
}

//@ 生成语音验证码
func CaptchaMP3() (captchaId, base64MP3 string) {
	//创建声音验证码
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	//以base64编码
	base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
	return idKeyA, base64stringA
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

//// 设置CaptchaConfig类
//type CaptchaConfig struct {
//	Id              string
//	CaptchaType     string
//	VerifyValue     string
//	ConfigAudio     base64Captcha.ConfigAudio
//	ConfigCharacter base64Captcha.ConfigCharacter
//	ConfigDigit     base64Captcha.ConfigDigit
//}
//
//// 创建CaptchaConfig类实例
//var (
//	captchaConfig     *CaptchaConfig
//	captchaConfigOnce sync.Once
//)
//
//// 获取base64验证码基本配置
//func GetCaptchaConfig() *CaptchaConfig {
//	captchaConfigOnce.Do(func() {
//		captchaConfig = &CaptchaConfig{
//			Id:          "",
//			CaptchaType: "character",
//			VerifyValue: "",
//			ConfigAudio: base64Captcha.ConfigAudio{},
//			ConfigCharacter: base64Captcha.ConfigCharacter{
//				Height:             60,
//				Width:              240,
//				Mode:               2,
//				IsUseSimpleFont:    false,
//				ComplexOfNoiseText: 0,
//				ComplexOfNoiseDot:  0,
//				IsShowHollowLine:   false,
//				IsShowNoiseDot:     false,
//				IsShowNoiseText:    false,
//				IsShowSlimeLine:    false,
//				IsShowSineLine:     false,
//				CaptchaLen:         0,
//			},
//			ConfigDigit: base64Captcha.ConfigDigit{},
//		}
//	})
//	return captchaConfig
//}
//
////@获取验证码
//func CaptchaGeneratePng() (captchaId, base64Png string) {
//	captchaConfig := GetCaptchaConfig()
//	// 创建base64图像验证码
//	character := captchaConfig.ConfigCharacter
//	// /GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
//	captchaId, digitCap := base64Captcha.GenerateCaptcha("", character)
//
//	// 生成图片
//	base64Png = base64Captcha.CaptchaWriteToBase64Encoding(digitCap)
//	return
//}

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
