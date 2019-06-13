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
└── uuid.go  uuid
```

把生活和工作当中常用的东西抽离出来
严格单元测试 保证代码质量
献丑了

### 简单用法
#### crypto加密包
- 把str转为md5
    ``` 
    md5str := Md5Encode(string)
    ```
- sha1
    ``` 
    sha1str := Sha1Encode("hello")
    ```
#### simpleTime 时间包
- 获取当前时间戳,时区默认亚洲上海
    ``` 
    timeString := GetCurrentTime()
    ```
- 时间戳转换为日期
    ``` 
    日期,err := GetTimeToString(时间戳string)
    ```
- 日期转换为时间戳
    ``` 
    时间戳str,err := GetTimeStringToTime(日期)
    ```
#### uuid包
- 获取uuid
    ``` 
    uuidstr,err := NewUUID()
    ```
- 获取当前uuid不带-
    ``` 
    uuidstr,err := NewUUIDSimplicity()
    ```
#### session 包
- 获取session
   ``` 
   session := GetSession("dollarkiller",6*60*60)
   ```
- 验证session
   ``` 
   bool := CheckSession(session)
   ``` 
- 销毁session
    ```
    DelSession(session)
        ```
#### file包
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