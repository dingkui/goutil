package jsonutil_test

import (
	"gitee.com/dk83/goutils/jsonutil"
	"gitee.com/dk83/goutils/zlog"
	"testing"
)

type TypeA struct {
	A string
	B string
}
type TypeB struct {
	A1 *TypeA
	A2 *TypeA
}
type TypeA2 struct {
	A string
	B string
	b string
	C string
	d string
}
type TypeC struct {
	A1 *TypeA2
}

const unCheck = "unCheck"
const failDef = "failDef"

func check(t *testing.T, fmt string, got interface{}, want string) {
	t.Helper()
	_got, e := jsonutil.Str(failDef, got)
	if e != nil {
		t.Errorf("Str err: got %#q; want %s", got, want)
		return
	}
	if want != unCheck && _got != want {
		t.Errorf("check fail: got %s; want %s", _got, want)
		return
	}
	zlog.TEST.Log(fmt, _got)
}

//构造函数测试
func TestJsonGo1(t *testing.T) {
	a1 := &TypeA{A: "a1", B: "b1"}
	a1Str := `{"A":"a1","B":"b1"}`
	a2 := &TypeA{A: "a2", B: "b2"}
	a2Str := `{"A":"a2","B":"b2"}`
	b1 := &TypeB{A1: a1, A2: a2}
	b1Str := `{"A1":{"A":"a1","B":"b1"},"A2":{"A":"a2","B":"b2"}}`

	jsonB1, _ := jsonutil.NewJsonGo(b1)

	check(t, "a1:%s", a1, a1Str)
	check(t, "a2:%s", a2, a2Str)
	check(t, "b1:%s", b1, b1Str)
	check(t, "jsonB1:%s", jsonB1, b1Str)

	jsonStr1, _ := jsonutil.NewJsonGo("aaa")
	jsonFloat1, _ := jsonutil.NewJsonGo(1.1)
	jsonInt1, _ := jsonutil.NewJsonGo(123.1)
	check(t, "jsonStr1.Str:%s", jsonStr1, "aaa")
	check(t, "jsonStr2.Str:%s", jsonFloat1, "1.1")
	check(t, "jsonInt1.Str:%s", jsonInt1, "123")
	check(t, "jsonFloat1.IntN:%s", jsonFloat1.IntN(2), "2")
	check(t, "jsonInt1.FloatN:%s", jsonInt1.FloatN(1), "123")
}

// TestJsonGo2 As 测试
func TestJsonGo2(t *testing.T) {
	b1Str := `{"A1":{"A":"a1","B":"b1","b":1},"A2":{"A":"a2","B":"b2"}}`
	jsonB1, _ := jsonutil.NewJsonGo(b1Str)
	c2 := &TypeC{}
	c2Str := `{"A1":{"A":"a1","B":"b1","C":""}}`
	jsonB1.As(c2)
	check(t, "TypeC:%s", c2, c2Str)
	zlog.Info("------------------------------------------------------")
	a21 := &TypeA{}
	jsonB1.As(a21, "A1")
	a21Str := `{"A":"a1","B":"b1"}`
	check(t, "TypeA:%s", jsonutil.StrN(failDef, a21), a21Str)
	zlog.Info("------------------------------------------------------")
	a3 := &TypeA2{}
	a31 := &TypeA2{}
	a3Str := `{"A":"a1","B":"b1","C":""}`
	jsonB1.As(a3, "A1")
	jsonB1.As(&a31, "A1")
	check(t, "TypeA2:%s", jsonutil.StrN(failDef, a3), a3Str)
	check(t, "&TypeA2:%s", jsonutil.StrN(failDef, a31), a3Str)
	zlog.Info("------------------------map------------------------------")
	a4 := make(map[string]interface{})
	a4Str := `{"A":"a1","B":"b1","b":1}`
	jsonB1.As(&a4, "A1")
	check(t, "map a4:%s", jsonutil.StrN(failDef, a4), a4Str)
}

//Get Set 测试
func TestJsonGo3(t *testing.T) {
	c1Str := `{"A1":{"A":"a1","B":"b1"}}`
	jsonB1, _ := jsonutil.NewJsonGo(c1Str)
	check(t, "jsonB1.Str:%s", jsonB1, c1Str)
	check(t, "jsonB1.Get(\"A1.A\"):%s", jsonB1.StrN(failDef, "A1.A"), failDef)
	check(t, "jsonB1.Get(\"A1.A\"):%s", jsonB1.StrN(failDef, "A1", "A"), "a1")
	check(t, "jsonB1.Get(\"@A1.A\"):%s", jsonB1.StrN(failDef, "@A1.A"), "a1")
	zlog.Info("-------------------------Set-----------------------------")
	jsonB1.Set("bb", "A1.A")
	c1Str1 := `{"A1":{"A":"a1","B":"b1"},"A1.A":"bb"}`
	check(t, "jsonB1.Str:%s", jsonB1, c1Str1)
	check(t, "jsonB1.Get(\"A1.A\"):%s", jsonB1.StrN(failDef, "A1.A"), "bb")
	zlog.Info("-------------------------ReNew-----------------------------")
	jsonB1.ReNew(c1Str)
	check(t, "jsonB1.Str:%s", jsonB1, c1Str)
	zlog.Info("-------------------------Set-----------------------------")
	jsonB1.Set(12.3, "@A1.x.d")
	check(t, "jsonB1.Str:%s", jsonB1, unCheck)
	check(t, "jsonB1.Get:%s", jsonB1.StrN(failDef, "@A1.x.d"), "12.3")
	jsonB1.Set(12.3, "@A3.0.x.d")
	jsonB1.Set(33, "@A3.1.x.d")
	check(t, "jsonB1.Str:%s", jsonB1, unCheck)
}

func TestJsonGo4(t *testing.T) {
	zlog.Info("-------------------------Set out of index-----------------------------")
	c1Str := `{"A1":{"A":"a1","B":"b1"}}`
	jsonB1, _ := jsonutil.NewJsonGo(c1Str)
	jsonB1.Set(33, "@A3.3.x.d")
	check(t, "jsonB1.Str:%s", jsonB1, `{"A1":{"A":"a1","B":"b1"},"A3":[]}`)
	zlog.Info("-------------------------reset test -----------------------------")
	jsonB1.ReNew("[12.3]")
	check(t, "jsonB1:%s", jsonB1, "[12.3]")
	jsonB1.Set(12.33, 1)
	check(t, "jsonB1:%s", jsonB1, "[12.3,12.33]")
	check(t, "jsonB1:%s", jsonB1.StrN(failDef, 0), "12.3")
	jsonB1.Set(666, 1)
	check(t, "jsonB1:%s", jsonB1.StrN(failDef, 1), "666")
	zlog.Info("-------------------------errKey index-----------------------------")
	check(t, "jsonB1:%s", jsonB1.StrN(failDef, 9), failDef)
	check(t, "jsonB1.StrN(\"@A2.3.x.d\"):%s", jsonB1.StrN(failDef, "@A2.3.x.d"), failDef)
}

func TestRemove(t *testing.T) {
	input := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}]}`
	jsonB1, _ := jsonutil.NewJsonGo(input)
	check(t, "jsonB1:%s", jsonB1, input)
	zlog.Info("------------------------Set to last------------------------------")
	jsonB1.Set(1, "@A4.-1")
	jsonB1.Set(2, "@A4.-1")
	jsonB1.Set(3, "@A4.-1")
	strSet2 := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":[1,2,3]}`
	check(t, "jsonB1:%s", jsonB1, strSet2)
	zlog.Info("------------------------remove------------------------------")
	jsonB1.Remove("@A4.-1")
	strRemove1 := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":[1,2]}`
	check(t, "jsonB1:%s", jsonB1, strRemove1)
	jsonB1.Remove("@A4")
	check(t, "jsonB1:%s", jsonB1, input)
	zlog.Info("------------------------Set to last------------------------------")
	jsonB1.Set(1, "A4", "-1")
	jsonB1.Set(2, "A4", "-1")
	jsonB1.Set(3, "A4", "-1")
	check(t, "jsonB1:%s", jsonB1, strSet2)
	zlog.Info("------------------------remove------------------------------")
	jsonB1.Remove("@A4.-1")
	check(t, "jsonB1:%s", jsonB1, strRemove1)
	jsonB1.Remove("@A4")
	check(t, "jsonB1:%s", jsonB1, input)
}

func TestNewJsonFile(t *testing.T) {
	input := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}]}`
	jsonB1, _ := jsonutil.NewJsonGo(input)
	file, _ := jsonutil.NewJsonFile("./test.json", jsonB1)
	zlog.Info("------------------------un onload ------------------------------")
	check(t, "jsonB1:%s", jsonB1, input)
	check(t, "file:%s", file, input)
	zlog.Info("------------------------save------------------------------")
	file.SaveFormat()
	jsonB1.Set(12.3, "@A4.www")
	strSet := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":{"www":12.3}}`
	check(t, "file:%s", file, strSet)
	check(t, "jsonB1:%s", jsonB1, strSet)
	zlog.Info("------------------------reload------------------------------")
	file.ReLoad()
	check(t, "file:%s", file, input)
	check(t, "jsonB1:%s", jsonB1, input)
	zlog.Info("------------------------Set to last------------------------------")
	file.Set(1, "@A4.-1")
	file.Set(2, "@A4.-1")
	file.Set(3, "@A4.-1")
	strSet2 := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":[1,2,3]}`
	check(t, "file:%s", file, strSet2)
	check(t, "jsonB1:%s", jsonB1, strSet2)
	zlog.Info("------------------------remove------------------------------")
	file.Remove("@A4.-1")
	strRemove1 := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":[1,2]}`
	check(t, "file:%s", file, strRemove1)
	check(t, "jsonB1:%s", jsonB1, strRemove1)
	file.Remove("@A4")
	check(t, "file:%s", file, input)
	check(t, "jsonB1:%s", jsonB1, input)
}
