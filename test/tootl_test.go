package test

import (
	"fmt"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/dollarkillerx/easyutils/compression"
	"github.com/dollarkillerx/easyutils/gcache"
	"github.com/dollarkillerx/easyutils/gemail"
	"github.com/dollarkillerx/easyutils/httplib"
	"log"
	"strings"
	"testing"
	"time"
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

func TestCache(t *testing.T) {
	//err := gcache.CacheSetTime("ok", "12312", 10)
	//if err != nil {
	//	panic(err)
	//}
	//
	//get, b := gcache.CacheGet("ok")
	//if !b {
	//	clog.Println("木有")
	//}else{
	//	clog.Println(get)
	//}
	//
	//time.Sleep(time.Second * 10)
	//
	//exit := gcache.Exit("ok")
	//if exit {
	//	clog.Println("茨木")
	//}else {
	//	clog.Println("kaka")
	//}

	// 并发安全测试
	go func() {
		for i := 0; i < 10000; i++ {
			go func(i int) {
				err := gcache.CacheSetTime("ok", "12312", 10)
				if err != nil {
					panic(err.Error())
				}
			}(i)
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			go func(i int) {
				get, b := gcache.CacheGet("ok")
				if b {
					clog.Println(get)
					clog.Println("茨木")
				} else {
					clog.Println("kaka")
				}
			}(i)
		}
	}()

	time.Sleep(time.Second * 20)

}

func TestGetSession(t *testing.T) {
	// 增加
	session := easyutils.SessionGetByGoCache()
	node := easyutils.Session{Name: "DollarKiller"}
	s, e := session.Set(&node)
	if e != nil {
		panic(e.Error())
	}
	log.Println(s)

	get, e := session.Get(s)
	if e != nil {
		panic(e)
	} else {
		log.Println(get)
	}

	expired := session.Expired(s)
	if !expired {
		log.Println("not data")
	} else {
		clog.Println("存在啊")
	}

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
	if e != nil {
		panic(e.Error())
	}
	t.Log(s)
}

func TestDirPing(t *testing.T) {
	DirPings("./HELLO/path")
}

func DirPings(path string) {
	split := strings.Split(path, "/")
	log.Println(len(split))
	log.Println(split)
}

// 测试数组切片
func TestSj(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	zsdel(a, 3)

}

func zsdel(data interface{}, index int) []interface{} {

	i, ok := data.([]interface{})
	if ok {
		log.Println(i)
	} else {
		log.Println("not ")
	}
	return nil
}

//type List struct {
//	data interface{}
//}
//
//func (l *List) Add(item interface{})  {
//	t := reflect.ValueOf(item).Kind()
//	l.data = new([]type)
//}

type cc struct {
	Ds string
	He int
}

func TestLogger(t *testing.T) {
	logf()
}

func logf() {
	clog.Println("hello")
	clog.PrintWa("sadsa")

	clog.PrintWa(cc{Ds: "dadas", He: 23})
}

func TestZip(t *testing.T) {
	zip := compression.Zip{}
	err := zip.Zip("./clog", "out.zip")
	if err != nil {
		panic(err)
	}
}

func TestUnZip(t *testing.T) {
	zip := compression.Zip{}
	unzip := zip.UnZip("out.zip", "./ps")
	if unzip != nil {
		panic(unzip)
	}
}

func TestUUID(t *testing.T) {
	s, e := easyutils.NewUUIDSimplicity()
	if e != nil {
		panic(e)
	}

	log.Println(s)
}

func TestDowPy(t *testing.T) {
	_, e := httplib.ProxyDow("https://www.google.com/", "127.0.0.1:8002")
	if e != nil {
		errstr := e.Error()
		index := strings.Index(errstr, "e.Error()")
		if index == -1 {

		}
	}

}

// 测试 打印颜色 log
func TestColor(t *testing.T) {
	//data := fmt.Sprintf("%c[1;31;40m[%v]%c[0m %v", 0x1B, "err", 0x1B,"萨达所大所大")
	//log.Println(data)
	clog.Println("你好")

	clog.PrintWa("wadsa")

	clog.PrintEr("errsa")
}

func TestCol(t *testing.T) {
	fmt.Println("")

	clog.Test()
}

// 测试邮件
func TestEmail(t *testing.T) {
	fromUser := "Hello"
	toUser := "490890221@qq.com"
	subject := "hello,world"
	err := gemail.SendNifoLog([]string{toUser}, fromUser, subject)
	if err != nil {
		log.Println(err.Error())
		log.Println("发送邮件失败")
		return
	}
	log.Println("发送邮件成功")

	//e := email.NewEmail()
	//e.From = "Jordan Wright <notice@dollarkiller.com>"
	//e.To = []string{"adapa@qq.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	//e.Subject = "Awesome Subject"
	//e.Text = []byte("Text Body is, of course, supported!")
	//e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	//send := e.Send("smtp.mail.ru:465", smtp.PlainAuth("", "notice@dollarkiller.com", "%Y4I4qjlqKAy", "smtp.mail.ru"))
	//if send != nil {
	//	panic(send)
	//}

}

func TestPanic(t *testing.T) {
	defer func() {
		if i := recover();i != nil {
			trace := clog.PanicTrace(2048)
			log.Println(string(trace))
		}
	}()
	panic("adasd")
}