# goutil
- 可以提高编码效率的工具包
- 没有任何第三方依赖
- 轻量级
- 最低支持1.15版本
- 为效率而生
- 主要解决什么问题？
  - go的类型转换，太头痛，需要自己做类型转换，非常麻烦
  - map赋值，取值，太麻烦，存在判断，一不小心空指针，更不用说多级数据存取处理了
  - 优秀的日志组件很多，但是很难搞清楚是哪个业务代码打出来的啊，难定位

## apputil
> 取得当前app运行参数
- 直接取值并转换
## dlog 日志工具
- 主要特点是可以打印出日志调用的地方，方便定位问题。
- 日志级别控制
- 日志文件输出，支持按日期切割
- 支持堆栈打印
- 日志调用级别自定义，在工具中打印日志也可以知道来源
- 支持自定义日志输出
- 初始化轻松搞定
```go
  dlog.AddAppenderConsole(0) //控制台输出
  dlog.AddAppenderDaily(1, "./logs/test.%s.log") //按天切割，文件日志
  dlog.AddAppenderDaily(2, "./logs/test.err.%s.log") //按天切割，文件日志
  dlog.AddAppenderRemote(2, "https://xxx.com/upload", nil) //上传到远程服务器 支持自定义header
  
  dlog.Info("test", 1, 1, 1) //普通输出
  dlog.Info("test %d %d %d", 1, 1, 1) //格式化输出
  dlog.Error(err,"test %d %d %d", 1, 1, 1)  //错误
  dlog.ErrorStack(err,"test %d %d %d", 1, 1, 1) //错误堆栈
```
## errs
- 自定义错误类型
- 支持错误类型判断
- 支持错误链的打印
- 支持追踪错误发生地
```sh
2024/11/14 18:41:04.327 WARN djson_test.go:144: json_go_gas.go:47 [9003:ErrerrKey] index is out of range:9,from [12.3,666]
2024/11/14 18:41:04.327 TEST djson_test.go:144: OK: jsonB1:failDef
2024/11/14 18:41:04.327 WARN djson_test.go:145: int.go:51 [2004:ErrErrTargetType] string to int err.strconv.Atoi: parsing "xxx": invalid syntax
2024/11/14 18:41:04.327 TEST djson_test.go:145: OK: jsonB1:9
2024/11/14 18:41:04.327 WARN djson_test.go:146: json_go.go:38 [9001:ErrerrJsonType] target is not a Map:[12.3,666]
2024/11/14 18:41:04.327 TEST djson_test.go:146: OK: jsonB1.StrN("@A2.3.x.d"):failDef
```

## djson
- json解析：支持结构体，map, array,json格式字符串
- 多级取值很简单，一个@搞定
- 方便的取值和设置，不用做无数次的存在判断
- 取值同时支持类型转换
```go
  input := `{"A1":{"A":"a1","B":"1"},"A2":{"A":"a2","B":"b2"}}`
  //input 可以是json字符串，也可以是map，struct，array等类型

  jsonB, _ := djson.NewJsonGo(input)
  jsonB.StrN("", "@A1.A")
  //取出字符串 “a1”
  jsonB.IntN(0, "@A1.B")
  //取出 int 1  
  //常规方法20行代码肯定搞不定这个需求  

  jsonB.StrN("")
  //序列化成json字符串 {"A1":{"A":"a1","B":"1"},"A2":{"A":"a2","B":"b2"}}

  //用常规方法做到这一个功能，常规方法要20行代码以上，这里一行搞定
  jsonB.Set("@A2.A", "a3")
  // a1 修改成 a3
  jsonB.StrN("")
  //序列化成json字符串 {"A1":{"A":"a3","B":"1"},"A2":{"A":"a2","B":"b2"}}

  //用常规方法做到这一个功能，常规方法要30行代码以上，这里一行搞定
  jsonB.Set("@C.x", "我是新来的")
  jsonB.StrN("")
  //{"A1":{"A":"a3","B":"1"},"A2":{"A":"a2","B":"b2"},"C":{"x":"我是新来的"}}

  obj := &A{}
  jsonB.As(obj)
  //将jsonB 转换成结构体obj
  
  //此类方法还有其他方法，可以参考源码。
```
## entry
- 加密解密工具，支持aes，des，rsa, base64,sha 加密解密
- 支持md5加密
- 支持基于有效期的加解密
## dhttp
> http请求工具，支持get，post，文件上传，文件下载，支持自定义header，cookie，超时设置，支持自定义http client。
## utils 一些工具，极大简化go代码编写。
- confUtil
####比如：你要写一个配置文件，映射一个结构体，初始化时从文件中读取信息，结构体内容变化需要写到配置文件中。只需要一个方法，就可以完成。
```go
  type Config struct {}
  Conf := confUtil.NewConfig("./conf/conf.json", &Config{},true)
  //保存到文件
  Conf.Save()
  //重新读取文件
  Conf.Reload()
```
- dateUtil
  - 日期类型转换工具，支持时间，字符串，时间戳之间互相转换
  - 预设格式化符串
  - 支持自定义格式化符串
- idUtil
  - 类似UUID算法，可以生成不同长度的随机字符串。
  - 可以是8进制，16进制，10进制，62进制，36进制
  - 可以自定义进制
- valUtil 强大的类型转换工具，可以很方便的实现不同对象或结构体间的类型转换
  - ToBytes
  - ToStr
  - ToInt
  - ToBool
  - ToFloat
  - ToMany...
  - 更多复杂类型，用djson来处理吧
