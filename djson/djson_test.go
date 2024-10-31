package djson_test

import (
	"encoding/json"
	"fmt"
	"gitee.com/dk83/goutils/djson"
	"gitee.com/dk83/goutils/dlog"
	"gitee.com/dk83/goutils/utils/valUtil"
	"reflect"
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
	_got, e := valUtil.Str(got, failDef)
	if e != nil {
		t.Errorf("Str err: got %#q; want %s", got, want)
		return
	}
	if want != unCheck && _got != want {
		t.Errorf("check fail: got %s; want %s", _got, want)
		return
	}
	dlog.TEST.Log(fmt, _got)
}

//构造函数测试
func TestJsonGo1(t *testing.T) {
	a1 := &TypeA{A: "a1", B: "b1"}
	a1Str := `{"A":"a1","B":"b1"}`
	a2 := &TypeA{A: "a2", B: "b2"}
	a2Str := `{"A":"a2","B":"b2"}`
	b1 := &TypeB{A1: a1, A2: a2}
	b1Str := `{"A1":{"A":"a1","B":"b1"},"A2":{"A":"a2","B":"b2"}}`

	jsonB1, _ := djson.NewJsonGo(b1)

	check(t, "a1:%s", a1, a1Str)
	check(t, "a2:%s", a2, a2Str)
	check(t, "b1:%s", b1, b1Str)
	check(t, "jsonB1:%s", jsonB1, b1Str)

	jsonStr1, _ := djson.NewJsonGo("aaa")
	jsonFloat1, _ := djson.NewJsonGo(1.1)
	jsonInt1, _ := djson.NewJsonGo(123.1)
	check(t, "jsonStr1.Str:%s", jsonStr1, "aaa")
	check(t, "jsonStr2.Str:%s", jsonFloat1, "1.1")
	check(t, "jsonInt1.Str:%s", jsonInt1, "123")
	check(t, "jsonFloat1.IntN:%s", jsonFloat1.IntN(2), "2")
	check(t, "jsonInt1.FloatN:%s", jsonInt1.Float64N(1), "123")
}

// TestJsonGo2 As 测试
func TestJsonGo2(t *testing.T) {
	b1Str := `{"A1":{"A":"a1","B":"b1","b":1},"A2":{"A":"a2","B":"b2"}}`
	jsonB1, _ := djson.NewJsonGo(b1Str)
	c2 := &TypeC{}
	c2Str := `{"A1":{"A":"a1","B":"b1","C":""}}`
	jsonB1.As(c2)
	check(t, "TypeC:%s", c2, c2Str)
	dlog.Info("------------------------------------------------------")
	a21 := &TypeA{}
	jsonB1.As(a21, "A1")
	a21Str := `{"A":"a1","B":"b1"}`
	check(t, "TypeA:%s", valUtil.StrN(a21, failDef), a21Str)
	dlog.Info("------------------------------------------------------")
	a3 := &TypeA2{}
	a31 := &TypeA2{}
	a3Str := `{"A":"a1","B":"b1","C":""}`
	jsonB1.As(a3, "A1")
	jsonB1.As(&a31, "A1")
	check(t, "TypeA2:%s", valUtil.StrN(a3, failDef), a3Str)
	check(t, "&TypeA2:%s", valUtil.StrN(a31, failDef), a3Str)
	dlog.Info("------------------------map------------------------------")
	a4 := make(map[string]interface{})
	a4Str := `{"A":"a1","B":"b1","b":1}`
	jsonB1.As(&a4, "A1")
	check(t, "map a4:%s", valUtil.StrN(a4, failDef), a4Str)
}

//Get Set 测试
func TestJsonGo3(t *testing.T) {
	c1Str := `{"A1":{"A":"a1","B":"b1"}}`
	jsonB1, _ := djson.NewJsonGo(c1Str)
	check(t, "jsonB1.Str:%s", jsonB1, c1Str)
	check(t, "jsonB1.Get(\"A1.A\"):%s", jsonB1.StrN(failDef, "A1.A"), failDef)
	check(t, "jsonB1.Get(\"A1.A\"):%s", jsonB1.StrN(failDef, "A1", "A"), "a1")
	check(t, "jsonB1.Get(\"@A1.A\"):%s", jsonB1.StrN(failDef, "@A1.A"), "a1")
	dlog.Info("-------------------------Set-----------------------------")
	jsonB1.Set("bb", "A1.A")
	c1Str1 := `{"A1":{"A":"a1","B":"b1"},"A1.A":"bb"}`
	check(t, "jsonB1.Str:%s", jsonB1, c1Str1)
	check(t, "jsonB1.Get(\"A1.A\"):%s", jsonB1.StrN(failDef, "A1.A"), "bb")
	dlog.Info("-------------------------ReNew-----------------------------")
	jsonB1.ReNew(c1Str)
	check(t, "jsonB1.Str:%s", jsonB1, c1Str)
	dlog.Info("-------------------------Set-----------------------------")
	jsonB1.Set(12.3, "@A1.x.d")
	check(t, "jsonB1.Str:%s", jsonB1, unCheck)
	check(t, "jsonB1.Get:%s", jsonB1.StrN(failDef, "@A1.x.d"), "12.3")
	jsonB1.Set(12.3, "@A3.0.x.d")
	jsonB1.Set(33, "@A3.1.x.d")
	check(t, "jsonB1.Str:%s", jsonB1, unCheck)
}

func TestJsonGo4(t *testing.T) {
	dlog.Info("-------------------------Set out of index-----------------------------")
	c1Str := `{"A1":{"A":"a1","B":"b1"}}`
	jsonB1, _ := djson.NewJsonGo(c1Str)
	jsonB1.Set(33, "@A3.3.x.d")
	check(t, "jsonB1.Str:%s", jsonB1, `{"A1":{"A":"a1","B":"b1"},"A3":[]}`)
	dlog.Info("-------------------------reset test -----------------------------")
	jsonB1.ReNew("[12.3]")
	check(t, "jsonB1:%s", jsonB1, "[12.3]")
	jsonB1.Set(12.33, 1)
	check(t, "jsonB1:%s", jsonB1, "[12.3,12.33]")
	check(t, "jsonB1:%s", jsonB1.StrN(failDef, 0), "12.3")
	jsonB1.Set(666, 1)
	check(t, "jsonB1:%s", jsonB1.StrN(failDef, 1), "666")
	dlog.Info("-------------------------errKey index-----------------------------")
	check(t, "jsonB1:%s", jsonB1.StrN(failDef, 9), failDef)
	check(t, "jsonB1.StrN(\"@A2.3.x.d\"):%s", jsonB1.StrN(failDef, "@A2.3.x.d"), failDef)
}

func TestRemove(t *testing.T) {
	input := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}]}`
	jsonB1, _ := djson.NewJsonGo(input)
	check(t, "jsonB1:%s", jsonB1, input)
	dlog.Info("------------------------Set to last------------------------------")
	jsonB1.Set(1, "@A4.-1")
	jsonB1.Set(2, "@A4.-1")
	jsonB1.Set(3, "@A4.-1")
	strSet2 := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":[1,2,3]}`
	check(t, "jsonB1:%s", jsonB1, strSet2)
	dlog.Info("------------------------remove------------------------------")
	jsonB1.Remove("@A4.-1")
	strRemove1 := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":[1,2]}`
	check(t, "jsonB1:%s", jsonB1, strRemove1)
	jsonB1.Remove("@A4")
	check(t, "jsonB1:%s", jsonB1, input)
	dlog.Info("------------------------Set to last------------------------------")
	jsonB1.Set(1, "A4", "-1")
	jsonB1.Set(2, "A4", "-1")
	jsonB1.Set(3, "A4", "-1")
	check(t, "jsonB1:%s", jsonB1, strSet2)
	dlog.Info("------------------------remove------------------------------")
	jsonB1.Remove("@A4.-1")
	check(t, "jsonB1:%s", jsonB1, strRemove1)
	jsonB1.Remove("@A4")
	check(t, "jsonB1:%s", jsonB1, input)
}

func TestNewJsonFile(t *testing.T) {
	input := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}]}`
	jsonB1, _ := djson.NewJsonGo(input)
	file, _ := djson.NewJsonFile("./test.json", jsonB1)
	dlog.Info("------------------------un onload ------------------------------")
	check(t, "jsonB1:%s", jsonB1, input)
	check(t, "file:%s", file, input)
	dlog.Info("------------------------save------------------------------")
	file.SaveFormat()
	jsonB1.Set(12.3, "@A4.www")
	strSet := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":{"www":12.3}}`
	check(t, "file:%s", file, strSet)
	check(t, "jsonB1:%s", jsonB1, strSet)
	dlog.Info("------------------------reload------------------------------")
	file.ReLoad()
	check(t, "file:%s", file, input)
	check(t, "jsonB1:%s", jsonB1, input)
	dlog.Info("------------------------Set to last------------------------------")
	file.Set(1, "@A4.-1")
	file.Set(2, "@A4.-1")
	file.Set(3, "@A4.-1")
	strSet2 := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":[1,2,3]}`
	check(t, "file:%s", file, strSet2)
	check(t, "jsonB1:%s", jsonB1, strSet2)
	dlog.Info("------------------------remove------------------------------")
	file.Remove("@A4.-1")
	strRemove1 := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":[1,2]}`
	check(t, "file:%s", file, strRemove1)
	check(t, "jsonB1:%s", jsonB1, strRemove1)
	file.Remove("@A4")
	check(t, "file:%s", file, input)
	check(t, "jsonB1:%s", jsonB1, input)
}

//func TestInt(t *testing.T) {
//
//	a := 123.3
//	b := djson.N0.IntN(a)
//	c := djson.N0.FloatN(a)
//	d := int(c)
//	check(t, "a:%v", a, "123.3")
//	//check(t, "b:%v", b, djson.N0.ToStr())
//	check(t, "c:%v", c, "123.3")
//	check(t, "d:%v", d, "123")
//	check(t, "float64(b):%v", float64(b), unCheck)
//	check(t, "float64(a):%v", float64(a), unCheck)
//	dlog.Info(c == float64(b))
//
//}
func TestInt2(t *testing.T) {
	check(t, "N0:%v", valUtil.IntN(float64(21.1)), unCheck)
	//check(t, "StrEmpty:%v",  djson.StrEmpty.ToStr(), unCheck)
	//check(t, "StrEmpty:%v",  djson.N0.IntN(0), unCheck)
	//check(t, "BoolFalse:%v",  djson.BoolFalse.ToStr(), unCheck)

	//of := reflect.TypeOf("xxx")
	of := reflect.ValueOf(123.1)

	marshal, _ := json.Marshal(float64(0))
	dlog.Info(string(marshal))
	marshal, _ = json.Marshal(float64(0.1))
	dlog.Info(string(marshal))
	marshal, _ = json.Marshal(float32(0.0))
	dlog.Info(string(marshal))

	dlog.Info(fmt.Sprintf("%T", float32(0.0)))
	dlog.Info(fmt.Sprintf("%v", float64(0)))
	dlog.Info(of.Float())
}
func TestTypes(t *testing.T) {
	x, _ := djson.NewJsonFile("", nil)
	x.Set("x", "xxx")
	var intf interface{} = x
	x1, ok := intf.(*djson.JsonFile)
	if ok {
		dlog.Info("1")
	}
	dlog.Info(x1.StrN("x"))
	_, ok = intf.(*djson.JsonGo)
	if ok {
		dlog.Info("2")
	}
}
func TestTypeB(t *testing.T) {
	a1 := &TypeA{A: "a1", B: "b1"}
	gojson := &TypeB{A1: a1}
	gojson2 := &TypeB{}

	x, err := djson.NewJsonGo(gojson)
	dlog.Info(err)
	dlog.Info(x)
	bytes, _ := x.Bytes()
	dlog.Info(string(bytes))
	x2, err := djson.NewJsonGo(bytes)
	dlog.Info(err)
	dlog.Info(x2)
	x2.As(gojson2)
	fmt.Println(gojson2)

}
func TestMars(t *testing.T) {
	jsonGo := djson.NewJsonMap()
	jsonGo.Set("xxValue", "@xx.xx")

	dlog.Info(jsonGo)
	marshal, _ := json.Marshal(jsonGo)
	dlog.Info(string(marshal))
	dlog.Info(jsonGo)
	json.Unmarshal(marshal, jsonGo)
	dlog.Info(jsonGo)
}

func TestJsonArray9(t *testing.T) {
	jsonGo, _ := djson.NewJsonArray()
	jsonGo.Set("xxValue", -1)
	jsonGo.Set("xxValue2", -1)
	jsonGo.Set("xxValue3", -1)

	dlog.Info(jsonGo)
	marshal, _ := json.Marshal(jsonGo)
	dlog.Info(string(marshal))
	dlog.Info(jsonGo)
	json.Unmarshal(marshal, jsonGo)
	dlog.Info(jsonGo)
	jsonGo.Set("xxValue22", "@xx.xx")
	dlog.Info(jsonGo.ToStr())
}
func TestJson9(t *testing.T) {
	jsonGo, _ := djson.NewJsonFile("", make([]string, 0))
	jsonGo.Set("xxValue", "@1")

	dlog.Info(jsonGo)
	marshal, _ := json.Marshal(jsonGo)
	dlog.Info(string(marshal))
	dlog.Info(jsonGo)
	json.Unmarshal(marshal, jsonGo)
	dlog.Info(jsonGo)
	jsonGo.Set("xxValue22", "@xx.xx")
	dlog.Info(jsonGo.ToStr())
}
