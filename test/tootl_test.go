package test

import (
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"log"
	"strings"
	"testing"
)

func TestNewUUID(t *testing.T) {
	s, e := easyutils.NewUUID()
	if e != nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestNewUUIDSimplicity(t *testing.T) {
	s, e := easyutils.NewUUIDSimplicity()
	if e != nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestMd5String(t *testing.T) {
	s := easyutils.Md5Encode("123")
	s = strings.ToUpper(s)
	if s != "202CB962AC59075B964B07152D234B70" {
		t.Error("非正常", s)
	}
}

func TestGetCurrentTime(t *testing.T) {
	time := easyutils.TimeGetNowTimeStr()
	t.Log(time)
}

func TestGetTimeToString(t *testing.T) {
	s, e := easyutils.TimeGetTimeToString("1558057058")
	if e != nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestGetTimeStringToTime(t *testing.T) {
	s, e := easyutils.TimeGetStringToTime("2019-05-17")
	if e != nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestGetSession(t *testing.T) {
	session := easyutils.SessionGenerate("dollarkiller", 6*60*60)
	t.Log(session)
	bool := easyutils.SessionCheck(session)
	t.Log(easyutils.SessionMap.Load(session))
	t.Log(bool)

	node, e := easyutils.SessionGetData(session)
	if e != nil {
		t.Fatal(e.Error())
	}
	t.Log(node)

	easyutils.SessionDel(session)
	bool = easyutils.SessionCheck(session)
	t.Log(bool)
}

func TestGenRsaKey(t *testing.T) {
	e, priKey, pubKey := easyutils.GenRsaKey(1024)
	if e == nil {
		t.Log(priKey)
		t.Log(pubKey)
	}

	data := "1231231245asdasd你好"
	s, e := easyutils.RsaEncryptSimple(data, pubKey)
	if e != nil {
		t.Fatal(e.Error())
	}
	simple, e := easyutils.RsaDecryptSimple(s, priKey)
	if e != nil {
		t.Fatal(e.Error())
	}
	if strings.EqualFold(data, simple) {
		t.Log("OK")
	}
	t.Logf(data)
	t.Logf(s)
	t.Logf(simple)
}

func TestRsaSign(t *testing.T) {
	e, priKey, pubKey := easyutils.GenRsaKey(1024)
	if e == nil {
		t.Log(priKey)
		t.Log(pubKey)
	}
	data := "1a2sd1as23d你好"
	s, e := easyutils.RsaSignSimple(data, priKey)
	if e != nil {
		t.Log(e.Error())
	}
	t.Log("签名: ", s)
	e = easyutils.RsaSignVerSimple(data, s, pubKey)
	if e != nil {
		t.Log("验证失败")
	}
	t.Log("验证成功")
}

func TestRsaDecryptSimple(t *testing.T) {
	e, priKey, pubKey := easyutils.GenRsaKey(1024)
	if e == nil {
		t.Log(len(priKey))
		t.Log(len(pubKey))
	}
	t.Logf(priKey)
	t.Logf(pubKey)
}

//@ 旧版本
//func TestNewUtilsToken(t *testing.T) {
//	_, priKey, pubKey := easyutils.GenRsaKey(1024)
//	jwt := easyutils.NewUtilsToken(priKey, pubKey)
//
//	head := &easyutils.JwtHeader{
//		Alg:  "alg",
//		Type: "rsa256",
//	}
//
//	i, _ := strconv.Atoi(easyutils.TimeGetNowTimeStr())
//	exp := strconv.Itoa(i + 60*60*6)
//
//	payload := &easyutils.JwtPayload{
//		Exp: exp,
//		Nbf: easyutils.TimeGetNowTimeStr(),
//	}
//	s, e := easyutils.GeneraJwtToken(head, payload)
//	if e != nil {
//		t.Fatal(e.Error())
//	}
//
//	bool := easyutils.VerificationToken(s)
//	if bool != true {
//		t.Logf("true")
//	}
//
//}

//@ 旧版本
//func TestCheckCaptcha1(t *testing.T) {
//	captchaId, base64Png := easyutils.CaptchaGeneratePng()
//	fmt.Println(captchaId)
//	fmt.Println(base64Png)
//	var key string
//	_, err := fmt.Scanln(&key)
//	if err != nil {
//		log.Println(err.Error())
//		return
//	}
//	i, e := easyutils.CaptchaCheck(captchaId, key)
//	if e != nil {
//		log.Println(e.Error())
//		log.Println(i)
//	} else {
//		log.Println(i)
//	}
//}

func TestCheckCaptcha2(t *testing.T) {
	key := "0EvRS72gtk23T1s4KDZG/n"
	i, e := easyutils.CaptchaCheck(key, "10989")
	if e != nil {
		log.Println(e.Error())
		log.Println(i)
	} else {
		log.Println(i)
	}
}

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		random := easyutils.Random(30, 100)
		t.Log(random)
	}

}

// easy token 测试
func TestEasyToken(t *testing.T) {
	payload := easyutils.EasyJwtPayload{}
	payload.Msg = "1213"
	payload.Iss = "dollarkiller"

	s, e := easyutils.EasyJwtGeneraToken(&payload, 1)
	if e != nil {
		t.Fatal(e.Error())
	}

	t.Log(s)
	// 验证
	e = easyutils.EasyJwtVerification(s)
	if e != nil {
		t.Fatal(e.Error())
	}
	t.Log("OK")
}

// 代理测试
func TestProxy(t *testing.T) {
	st, err := easyutils.InitProxy("127.0.0.1:8001")
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	// 没有问题测试一下
	err = st.CheckProxy()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Ok")

}

// Logs test
func TestLogs(t *testing.T) {
	clog.Println("123")

}


// url test
func TestUrlEncoding(t *testing.T) {
	url := "https://www.baidu.com/search?ok=sada sadad"
	s, e := easyutils.UrlEncoding(url)
	if e!= nil {
		panic(e.Error())
	}
	t.Log(s)
}