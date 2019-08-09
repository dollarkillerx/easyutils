package easyutils

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
)

// 有token 包 为什么还要写一个easyToken ？ token包 用户可以更灵活的使用。 easyToken包 分装的更为彻底

// easyToken 包的宗旨是让用户更简单的使用

type EasyJwtPayload struct {
	Iss  string      `json:"iss"` // 签发人
	Exp  int      `json:"exp"` // 过期时间
	Msg  string      `json:"msg"` // 用户自义定数据
	Data interface{} `json:"data"`
}

// token 池
type easyJwtUtils struct {
	priKey string   // 私钥
	pubKey string   // 公钥
	data   sync.Map // 存储
}

var (
	easyJwtUtil *easyJwtUtils
	db_com *sync.Map
)

func init() {
	// 初始化秘钥
	e, priKey, pubKey := GenRsaKey(1024)
	if e != nil {
		log.Println("++++++++++秘钥 生成失败+++++++++++")
		panic(e.Error())
	}
	data := easyJwtUtils{
		priKey: priKey,
		pubKey: pubKey,
	}
	easyJwtUtil = &data

	// 初始化 通用存储
	db_com = &sync.Map{}
}

// 生成tooken
// 载荷
func EasyJwtGeneraToken(data *EasyJwtPayload,hour int) (string,error) {
	addtime := hour * 60 * 60 + TimeGetNowTime()
	data.Exp = addtime

	payloadJson, e := json.Marshal(data)
	if e != nil {
		log.Println(e.Error())
		return "",e
	}
	payloadEnco := Base64Encode(payloadJson)
	// 签名
	signature, e := RsaSignSimple(payloadEnco, easyJwtUtil.priKey) // 签名
	if e != nil {
		log.Println(e.Error())
		return "", e
	}

	jwt := payloadEnco + "." + signature
	// 存入内存中
	easyJwtUtil.data.Store(jwt,data)
	// 将用户id 和 jwt 存入到通用存储中
	db_com.Store(data.Iss,jwt)

	return jwt,nil
}

// 验证签名
// jwt
func EasyJwtVerification(jwt string) error {
	value, ok := easyJwtUtil.data.Load(jwt)
	if !ok {
		return errors.New("not data")
	}
	payload := value.(*EasyJwtPayload)
	// 如果存在 则验证时间
	time := TimeGetNowTime()
	// 如果过期
	if payload.Exp < time {
		// 在内存中删除 并返回错误
		easyJwtUtil.data.Delete(jwt)
		// 删除 通用存储中的用户信息
		db_com.Delete(payload.Iss)
		return errors.New("time out")
	}
	return nil
}

// 验证用户是否已经登录
func EasyJwtCheckUserLogin(email string) error {
	value, ok := db_com.Load(email)
	if !ok {
		return errors.New("not data")
	}
	token := value.(string)
	// 验证token 是否有效
	verification := EasyJwtVerification(token)
	if verification != nil {
		return errors.New("not login")
	}
	return nil
}