/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-3
* Time: 下午8:47
* */
package utils

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"
)

/**
Token 包
思索后,考虑到未来的扩展型 只是做了 生成和验证    (只提供一个基于内存的存储)
*/

type JwtHeader struct {
	Alg  string `json:"alg"`  // 算法名称
	Type string `json:"type"` // 类型
}

type JwtPayload struct {
	Iss string `json:"iss"` // 签发人
	Exp string `json:"exp"` // 过期时间
	Sub string `json:"sub"` // 主题
	Nbf string `json:"nbf"` // 生效时间
	Iat string `json:"iat"` // 签发时间
	Jti string `json:"jti"` // 编号
}

// token
type JwtUtils struct {
	priKey string   // 私钥
	pubKey string   // 公钥
	data   sync.Map // 存储
}

// 初始化token
func NewUtilsToken(prikey string, pubkey string) *JwtUtils {
	return &JwtUtils{
		priKey: prikey,
		pubKey: pubkey,
		data:   sync.Map{},
	}
}

// 生成JWT
func (t *JwtUtils) GeneraJwtToken(header *JwtHeader, payload *JwtPayload) (string, error) {
	headerJson, e := json.Marshal(header)
	if e != nil {
		return "", e
	}
	payloadJson, e := json.Marshal(payload)
	if e != nil {
		return "", e
	}

	headerEnco := Base64Encode(headerJson)
	payloadEnco := Base64Encode(payloadJson)

	head := headerEnco + "." + payloadEnco        // 头 + 载荷
	signature, e := RsaSignSimple(head, t.priKey) // 签名
	if e != nil {
		return "", e
	}

	jwt := head + "." + signature

	return jwt, nil
}

// 生成token并存入 内存中
func (t *JwtUtils) GeneraJwtTokenToData(header *JwtHeader, payload *JwtPayload) (string, error) {
	s, e := t.GeneraJwtToken(header, payload)
	if e != nil {
		return "", e
	}
	t.data.Store(s, payload.Exp)
	return s, nil
}

// 验证token
func (t *JwtUtils) VerificationToken(jwt string) bool {
	data := strings.Split(jwt, ".")
	head := data[0] + "." + data[1]
	signature := data[2]

	simple := RsaSignVerSimple(head, signature, t.pubKey)
	if simple != nil {
		return false
	}

	payload := &JwtPayload{}
	err := json.Unmarshal([]byte(data[1]), payload)
	if err != nil {
		return false
	}

	end_time, _ := strconv.Atoi(payload.Exp)
	star_time, _ := strconv.Atoi(payload.Nbf)
	now_time, _ := strconv.Atoi(TimeGetNowTimeStr())

	if now_time < end_time && now_time > star_time {
		return true
	}

	return false
}

// 验证token从缓存中
func (t *JwtUtils) VerificationTokenByData(jwt string) bool {
	if value, ok := t.data.Load(jwt); ok != true {
		return false
	} else {
		s := value.(string)
		end_time, _ := strconv.Atoi(s)
		now_time, _ := strconv.Atoi(TimeGetNowTimeStr())

		if now_time > end_time {
			t.data.Delete(jwt)
			return false
		} else {
			return true
		}
	}
}
