package errs

import (
	"fmt"
	"github.com/dingkui/goutil/utils/runtimeUtil"
	"github.com/dingkui/goutil/utils/stringUtil"
)

//## 错误类型
var (
	_errs = make(map[int]*ErrType)

	ErrOther = Err(0, "Other") //其他方式构建的错误

	ErrSystem = Err(1000, "System") //系统错误，发生后一般需要修改程序才能恢复
	ErrEnv    = Err(1001, "Env")    //系统错误，一般由于运行条件不满足造成程序无法正常运行

	ErrValidate   = Err(2001, "Validate")      //校验错误，一般发生在数据验证时
	ErrBusiness   = Err(2002, "Business")      //校验错误，通常指业务错误
	ErrRuntime    = Err(2003, "Runtime")       //校验错误，通常指业务错误
	ErrTargetType = Err(2004, "ErrTargetType") //目标类型错误

	ErrRemote      = Err(3000, "Remote")      //校验错误
	ErrCredentials = Err(3401, "Credentials") //登录凭据失效
	ErrForbidden   = Err(3403, "Forbidden")   //无访问权限
	ErrHttp        = Err(3101, "Http")        //http调用时错误

	ErrDb = Err(4000, "Db") //校验错误
)

func Err(code int, msg string) *ErrType {
	err, exists := _errs[code]
	if exists {
		panic(fmt.Sprintf(
			"Make new ErrType [%d:%s] has error,the ErrType [%d:%s] is maked in %s",
			code, msg, code, err.msg, err.addr))
	}
	e := &ErrType{
		code: code,
		msg:  msg,
		addr: runtimeUtil.GetCaller(2),
	}
	_errs[code] = e
	return e
}

type ErrType struct {
	code int
	msg  string
	addr string
}

func (e *ErrType) New(msg interface{}, a ...interface{}) error {
	return e.new(nil, msg, a...)
}
func (e *ErrType) NewWithData(d interface{}, msg interface{}, a ...interface{}) error {
	return e.new(d, msg, a...)
}
func (e *ErrType) new(d interface{}, msg interface{}, a ...interface{}) error {
	err := &errInfo{
		t: e,
		a: runtimeUtil.GetCaller(3),
		d: d,
	}

	_e, ok := msg.(*Error)
	if ok {
		if len(a) > 0 {
			err.m = stringUtil.Fmt(a[0], a[1:]...)
		} else {
			err.m = e.msg
		}
		return &Error{trace: append([]*errInfo{err}, _e.trace...)}
	} else {
		err.m = stringUtil.Fmt(msg, a...)
		return &Error{trace: []*errInfo{err}}
	}
}


func (e *ErrType) IsType(err error) bool {
	_, ok := e.Is(err)
	return ok
}
func (e *ErrType) Is(err error) (*errInfo, bool) {
	_e, ok := err.(*Error)
	if !ok {
		return nil, false
	}
	return _e.Is(e)
}
func (e *ErrType) Msg(err error) string {
	_e, ok := err.(*Error)
	if !ok {
		return ""
	}
	for _, i2 := range _e.trace {
		if i2.t == e {
			return fmt.Sprintf("[%d:%s]:%s", e.code, e.msg, i2.m)
		}
	}
	return ""
}
func (e *ErrType) Data(err error) interface{} {
	_e, ok := err.(*Error)
	if !ok {
		return nil
	}
	for _, i2 := range _e.trace {
		if i2.t == e {
			return i2.d
		}
	}
	return nil
}
func (e *ErrType) Error() string {
	return fmt.Sprintf("%d:%s", e.code, e.msg)
}
