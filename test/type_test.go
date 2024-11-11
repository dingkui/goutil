package test

import (
	"fmt"
	"gitee.com/dk83/goutils/apputil"
	"gitee.com/dk83/goutils/dhttp"
	"gitee.com/dk83/goutils/djson"
	"gitee.com/dk83/goutils/dlog"
	"gitee.com/dk83/goutils/encry"
	"gitee.com/dk83/goutils/errs"
	"gitee.com/dk83/goutils/utils/confUtil"
	"gitee.com/dk83/goutils/utils/dateUtil"
	"gitee.com/dk83/goutils/utils/fileUtil"
	"gitee.com/dk83/goutils/utils/idUtil"
	"gitee.com/dk83/goutils/utils/mathUtil"
	"gitee.com/dk83/goutils/utils/runtimeUtil"
	"gitee.com/dk83/goutils/utils/stringUtil"
	"gitee.com/dk83/goutils/utils/valUtil"
	"reflect"
	"testing"
)

type TypeTest1 struct {
	name string
}
type TypeTest1Child struct {
	TypeTest1
}

func TestTypeTest1_1(t *testing.T) {
	child := &TypeTest1Child{}
	child.name = "t2"
	parent := child.TypeTest1 // 正确地进行类型断言
	parent.name = "t1"
	fmt.Printf("parent:%+v t2:%+v", parent, child)
}
func TestTypeTest1_2(t *testing.T) {
	child := &TypeTest1Child{}
	child.name = "t2"
	parent := &child.TypeTest1 // 正确地进行类型断言
	parent.name = "t1"
	fmt.Printf("parent:%+v t2:%+v", parent, child)
}
func checkType(t interface{}) {
	value, ok := t.(TypeTest1Child)
	if ok {
		fmt.Println("t is of type TypeTest1Child:", value)
	} else {
		fmt.Println("t is not of type TypeTest1Child:")
	}
	value2, ok := t.(TypeTest1)
	if ok {
		fmt.Println("t is of type TypeTest1:", value2)
	} else {
		fmt.Println("t is not of type TypeTest1:")
	}
	var d TypeTest1Child
	var b interface{} = d
	fmt.Printf("t type is:%s", reflect.TypeOf(t).String())
	if reflect.TypeOf(b).String() == "main.Base" {
		fmt.Println("b is exactly of type Base")
	} else {
		fmt.Println("b is not of type Base")
	}

	if _, ok := b.(TypeTest1); ok {
		fmt.Println("b is of type TypeTest1")
	} else {
		fmt.Println("b is not of type TypeTest1")
	}
}
func checkType2(t interface{}) {
	switch t.(type) {
	case TypeTest1:
		fmt.Println("checkType2 TypeTest1")
	case TypeTest1Child:
		fmt.Println("checkType2 TypeTest1Child")
	}
}
func checkType3(t interface{}) {
	switch t.(type) {
	case *TypeTest1:
		fmt.Println("checkType3 指针: TypeTest1")
	case TypeTest1:
		fmt.Println("checkType3 类型: TypeTest1")
	case *TypeTest1Child:
		fmt.Println("checkType3 指针: TypeTest1Child")
	case TypeTest1Child:
		fmt.Println("checkType3 类型: TypeTest1Child")
	}
}
func TestCheckType(t *testing.T) {
	child := TypeTest1Child{}
	child.name = "t2"
	checkType(child)
}
func TestCheckType2(t *testing.T) {
	child := &TypeTest1Child{}
	child.name = "t2"
	checkType3(child)
	checkType3(&child.TypeTest1)
}

func TestArray(t *testing.T) {
	s := []int{1}
	ints := s[:0]
	dlog.Info(ints)
}

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

	dlog.Info(errUnReached.Is(err1))
	dlog.Info(errUnReached.Is(err2))
	dlog.Info(errUnReached.Msg(err1))
	dlog.Info(errUnReached.Msg(err2))
	dlog.Info(errs.ErrSystem.Is(err1))
	dlog.Info(errs.ErrSystem.Is(err2))
	dlog.Info(errs.ErrSystem.Msg(err1))
	dlog.Info(errs.ErrSystem.Msg(err2))
	dlog.Info(err1)
	dlog.Info(err2)
}

func TestAll(t *testing.T) {
	dlog.Info(apputil.GetPara("xx", "1"))
	dlog.Info(dhttp.Client{})
	dlog.Info(djson.NewJsonMap())
	dlog.Info(encry.Md5)
	dlog.Info(errs.ErrSystem.New("ErrSystem"))
	dlog.Info(dateUtil.DateTime.FormatNow())
	dlog.Info(fileUtil.Exists("d:/"))
	dlog.Info(idUtil.ID16(32))
	dlog.Info(mathUtil.Round(12.1, 3))
	dlog.Info(runtimeUtil.GetCaller(1))
	dlog.Info(stringUtil.Fmt("%d", 1))
	dlog.Info(valUtil.Str(1))
	data := map[string]interface{}{"a": 2}
	dlog.Info(confUtil.NewConf("d:/xx/xx2.json", &data,true))
	dlog.Info(data)
}
