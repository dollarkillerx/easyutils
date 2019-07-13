# easyutils
easyutils Golang 常用工具库
``` 
.
├── crypto.go  加密解密相关
├── go.mod
├── LICENSE
├── README.md
├── session.go 简单的session
├── simpleTime.go  时间相关
├── tootl_test.go 测试
├── reptile.go 爬虫 
├── miscellaneous.go 杂项
└── uuid.go  uuid
```

把生活和工作当中常用的东西抽离出来
严格单元测试 保证代码质量
献丑了

## 简单用法
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
- 简单保存文件 返回随机名称(几乎不会重复,当目录文件夹存在会自行创建)
    ``` 
    FileSaveRenameSimple(name string, data []byte, path string) (string, error)
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
    
### 简单爬虫包   (未来会出一个高级包,参考scrapy)
- 获取 随机 普通浏览器 UserAgent
    ``` 
    userAgent := ReptileGetUserAgent()
    ```
- 伪装 随机 搜索引擎 UserAgent
    ``` 
    userAgent := ReptileGetSpiderAgent()
    ```
- 模拟搜索引擎发送请求
    ``` 
    ReptileRequestFrom(targerUrl string,body io.Reader,cookies []*http.Cookie) (*http.Response,error)
    ```
- 模拟浏览器获得下载文件 Simple版本
    ```
    ReptileDownloadSimple(targerUrl string,cookies []*http.Cookie) ([]byte,error)
    ```
        
        
### 杂项 
- 过滤HTML元素  
    ``` 
    TrimHtml(src string) string
    ```
- 随机中文名
    ``` 
    GetFullChinaName() string
    ```
