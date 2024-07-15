package errs

//## 错误类型

var (
	_errs = make(map[int]*ErrType)

	ErrOther = Err(0, "Other") //其他方式构建的错误

	ErrSystem = Err(1000, "System") //系统错误，发生后一般需要修改程序才能恢复
	ErrEnv    = Err(1001, "Env")    //系统错误，一般由于运行条件不满足造成程序无法正常运行

	ErrValidate = Err(2001, "Validate") //校验错误，一般发生在数据验证时
	ErrBusiness = Err(2002, "Business") //校验错误，通常指业务错误
	ErrRuntime  = Err(2003, "Runtime")  //校验错误，通常指业务错误

	ErrRemote      = Err(3000, "Remote")      //校验错误
	ErrCredentials = Err(3001, "Credentials") //校验错误，登录凭据失效
	ErrHttp        = Err(3101, "Http")        //http调用时错误

	ErrDb = Err(4000, "Db") //校验错误
)
