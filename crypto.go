package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
)

// 获取md5
func Md5Encode(str string) string {
	data := []byte(str)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// 获取sha1
func Sha1Encode(str string) string {
	data := []byte(str)
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}

// RSA256 公钥密钥对生成
// @params: bits 密钥长度
// @returns: private 密钥
// @returns: public 公钥
func GenRsaKey(bits int) (e error, priKey string, pubKey string) {

	// 生成私钥
	privateKey, e := rsa.GenerateKey(rand.Reader, bits)
	if e != nil {
		return e, "", ""
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:  "RAS PRIVATE KEY",
		Bytes: derStream,
	}
	//fmt.Println("私密钥:",string(pem.EncodeToMemory(priBlock)))
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return e, "", ""
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	//fmt.Printf("=======公钥文件内容=========%v", string(pem.EncodeToMemory(publicBlock)))

	if err != nil {
		return e, "", ""
	}
	return nil, string(pem.EncodeToMemory(priBlock)), string(pem.EncodeToMemory(publicBlock))
}

// Rsa256 加密
// @params: origData 原始数据
// @Params: pubKey 公钥
func RsaEncrypt(origData, pubKey []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// Rsa256 加密简单
// @params: origData 原始数据
// @Params: pubKey 公钥
func RsaEncryptSimple(origData, pubKey string) (string, error) {
	orgData := []byte(origData)
	pubK := []byte(pubKey)
	bytes, e := RsaEncrypt(orgData, pubK)
	if e != nil {
		return "", e
	}
	encode := Base64Encode(bytes)
	return encode, nil
}

// Rsa256 解密
// @params: ciphertext 加密数据
// @Params: prvKey 私钥
func RsaDecrypt(ciphertext, privateKey []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// Rsa256 解密简单
func RsaDecryptSimple(ciphertext, privateKey string) (string, error) {
	decode, i := Base64Decode(ciphertext)
	if i != nil {
		return "", i
	}
	pri := []byte(privateKey)
	bytes, e := RsaDecrypt(decode, pri)
	return string(bytes), e
}

// Base64编码
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64解码
func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
