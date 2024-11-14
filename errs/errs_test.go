package errs_test

import (
	"github.com/dingkui/goutil/dlog"
	"github.com/dingkui/goutil/errs"
	"testing"
)

var (
	errUnReached  = errs.Err(19001, "xx") //不应该到达的错误
	errValid      = errs.Err(19002, "xx") //校验错误
	errJsonType   = errs.Err(19003, "xx") //json数据类型错误
	errTarget     = errs.Err(19004, "xx") //目标错误
	errTargetType = errs.Err(19005, "xx") //目标类型错误
)

func TestErr(t *testing.T) {
	defer dlog.Recover()
	//errs.Err(9001, "xx232")    //不应该到达的错误

	err1 := errUnReached.New("errUnReached")
	err2 := errs.ErrSystem.New(err1, "ErrSystem")

	dlog.Info(errUnReached.IsType(err1))
	dlog.Info(errUnReached.IsType(err2))
	dlog.Info(errUnReached.Msg(err1))
	dlog.Info(errUnReached.Msg(err2))
	dlog.Info(errs.ErrSystem.IsType(err1))
	dlog.Info(errs.ErrSystem.IsType(err2))
	dlog.Info(errs.ErrSystem.Msg(err1))
	dlog.Info(errs.ErrSystem.Msg(err2))
	dlog.Info(err1)
	dlog.Info(err2)
}
