# easyutils
easyutils Golang 常用工具库
``` 
.
├── LICENSE
├── README.md
├── base64Captcha.go 验证码
├── crypto.go 加密解密相关
├── file.go 文件相关
├── go.mod
├── go.sum
├── miscellaneous.go 杂项
├── new_session.go session相关 新(开发中)
├── reptile.go 爬虫相关
├── session.go session相关 旧
├── simpleTime.go 时间相关
├── test 测试
│   └── tootl_test.go
├── token.go token相关
└── uuid.go uuid相关
```

把生活和工作当中常用的东西抽离出来
严格单元测试 保证代码质量

## 简单用法

### 获取
``` 
go get github.com/dollarkillerx/easyutils
```

### crypto加密包
- 把str转为md5
    ``` 
    md5str := Md5Encode(string)
    ```
- sha1
    ``` 
    sha1str := Sha1Encode("hello")
    ```
- RSA256 公钥密钥对生成
    ``` 
    e, priKey, pubKey := GenRsaKey(1024) // 1024 密钥长度
    ```
- Rsa256 加密
    ``` 
    RsaEncrypt(origData,pubKey []byte) ([]byte,error)
    ```
- Rsa256 加密简单
    ``` 
    RsaEncryptSimple(origData,pubKey string) (string,error)
    ```
- Rsa256 解密
    ```
    RsaDecrypt(ciphertext,privateKey []byte) ([]byte, error)
    ```
- Rsa256 解密简单
    ```
    RsaDecryptSimple(ciphertext,privateKey string) (string, error)
    ```
- Rsa256 签名
    ``` 
    RsaSignVer(data,signature,publicKey []byte) error
    ```
- Rsa256 简单签名
    ``` 
    RsaSignSimple(data, prvKey string) (string, error)
    ```
- Rsa256 验签
    ``` 
    RsaSignVer(data,signature,publicKey []byte) error
    ```
- Rsa256 简单验签
    ``` 
    RsaSignVerSimple(data,signature,publicKey string) error
    ```
- Base64编码
    ``` 
    Base64Encode(data []byte) string
    ```
- Base64解码
    ``` 
    Base64Decode(s string) ([]byte,error)
    ```
    
    
### simpleTime 时间包
- 获取当前时间戳,时区默认亚洲上海
    ``` 
    timeString := TimeGetNowTimeStr()
    ```
- 时间戳转换为日期
    ``` 
    日期,err := TimeGetTimeToString(时间戳string)
    ```
- 日期转换为时间戳
    ``` 
    时间戳str,err := TimeGetStringToTime(日期)
    ```
    
    
### uuid包
- 获取uuid
    ``` 
    uuidstr,err := NewUUID()
    ```
- 获取当前uuid不带-
    ``` 
    uuidstr,err := NewUUIDSimplicity()
    ```
    
    
### session 包
- 获取session
   ``` 
   session := SessionGenerate("dollarkiller",6*60*60)
   ```
- 获得session数据
    ``` 
    node, e := SessionGetData(session)
    ```
- 验证session
   ``` 
   bool := SessionCheck(session)
   ``` 
- 销毁session
    ```
    SessionDel(session)
    ```
    
    
### file包
- 判断文件夹是否存在
    ``` 
    ok,err := PathExists("./file")
    ```
- 如果文件夹不存在就会创建文件夹
    ``` 
    err := DirPing("./file")
    ```
- 获取文件后缀
    ``` 
    str,err := FileGetPostfix("123.text")
    ```
- 获得随机文件名 传入postfilx后缀
    ``` 
    filename := FileGetRandomName(postfilx string)
    ```
- 获取文件sha1
    ``` 
    str := FileGetSha1(file *os.File)
    ```
- 获取文件MD5
    ``` 
    str := FileGetMD5(file *os.File)
    ```
- 保存文件 并从命名 Simple 版本
    ``` 
    FileSaveRenameSimple(name string, data []byte, path string) (string, error)
    ```
- 保存文件 Simple版本
    ``` 
    FileSaveSimple(name string, data []byte, path string) error 
    ```
### Token
- 初始化
    ``` 
    _, priKey, pubKey := GenRsaKey(1024)
        jwt := NewUtilsToken(priKey, pubKey)
    ```
- 生成JWT
    ```
    s, e := jwt.GeneraJwtToken(head, payload)
    ```
- 生成JWT 缓存版本
    ``` 
    s, e := jwt.GeneraJwtTokenToData(head, payload)
    ```
- 验证JWT
    ``` 
    bool := jwt.VerificationToken(s)
    ```
- 验证JWT 缓存版本
    ``` 
    bool := jwt.VerificationTokenByData(s)
    ```
    
### 验证码  是对base64Captcha包装 
- 获取数字验证码
    ``` 
    CaptchaNum() (captchaId, base64Png string)
    ```
- 获取算数验证码
    ```
    CaptchaMath() (captchaId, base64Png string)
    ```
- 获取音频验证码
    ```
    CaptchaMP3() (captchaId, base64MP3 string)
    ```
- 验证
    ```
    CaptchaCheck(captchaId, verifyValue string) (int, error)
    ```
    
### 杂项
- 获取区间随机数
    ``` 
    Random(min,max int) int
    ```