package djson_test

import (
	"encoding/json"
	"fmt"
	"github.com/dingkui/goutil/djson"
	"github.com/dingkui/goutil/dlog"
	"github.com/dingkui/goutil/utils/testUtil"
	"github.com/dingkui/goutil/utils/valUtil/forceVal"
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

//构造函数测试
func TestJsonGo1(t *testing.T) {
	a1 := &TypeA{A: "a1", B: "b1"}
	a1Str := `{"A":"a1","B":"b1"}`
	a2 := &TypeA{A: "a2", B: "b2"}
	a2Str := `{"A":"a2","B":"b2"}`
	b1 := &TypeB{A1: a1, A2: a2}
	b1Str := `{"A1":{"A":"a1","B":"b1"},"A2":{"A":"a2","B":"b2"}}`

	jsonB1, _ := djson.NewJsonGo(b1)

	testUtil.Check(t, "a1:%s", a1, a1Str)
	testUtil.Check(t, "a2:%s", a2, a2Str)
	testUtil.Check(t, "b1:%s", b1, b1Str)
	testUtil.Check(t, "jsonB1:%s", jsonB1, b1Str)

	jsonStr1, _ := djson.NewJsonGo("aaa")
	jsonFloat1, _ := djson.NewJsonGo(1.1)
	jsonInt1, _ := djson.NewJsonGo(123.1)
	testUtil.Check(t, "jsonStr1.Str:%s", jsonStr1, "aaa")
	testUtil.Check(t, "jsonStr2.Str:%s", jsonFloat1, "1.1")
	testUtil.Check(t, "jsonInt1.Str:%s", jsonInt1, "123.1")
	testUtil.Check(t, "jsonFloat1.IntN:%s", jsonFloat1.IntN(2), "2")
	testUtil.Check(t, "jsonInt1.FloatN:%s", jsonInt1.Float64N(1), "123.1")
}

// TestJsonGo2 As 测试
func TestJsonGo2(t *testing.T) {
	b1Str := `{"A1":{"A":"a1","B":"b1","b":1},"A2":{"A":"a2","B":"b2"}}`
	jsonB1, _ := djson.NewJsonGo(b1Str)
	c2 := &TypeC{}
	c2Str := `{"A1":{"A":"a1","B":"b1","C":""}}`
	jsonB1.As(c2)
	testUtil.Check(t, "TypeC:%s", c2, c2Str)
	dlog.Info("------------------------------------------------------")
	a21 := &TypeA{}
	jsonB1.As(a21, "A1")
	a21Str := `{"A":"a1","B":"b1"}`
	testUtil.Check(t, "TypeA:%s", forceVal.Str(a21, testUtil.FailDef), a21Str)
	dlog.Info("------------------------------------------------------")
	a3 := &TypeA2{}
	a31 := &TypeA2{}
	a3Str := `{"A":"a1","B":"b1","C":""}`
	jsonB1.As(a3, "A1")
	jsonB1.As(&a31, "A1")
	testUtil.Check(t, "TypeA2:%s", forceVal.Str(a3, testUtil.FailDef), a3Str)
	testUtil.Check(t, "&TypeA2:%s", forceVal.Str(a31, testUtil.FailDef), a3Str)
	dlog.Info("------------------------map------------------------------")
	a4 := make(map[string]interface{})
	a4Str := `{"A":"a1","B":"b1","b":1}`
	jsonB1.As(&a4, "A1")
	testUtil.Check(t, "map a4:%s", forceVal.Str(a4, testUtil.FailDef), a4Str)
}

//Get Set 测试
func TestJsonGo3(t *testing.T) {
	c1Str := `{"A1":{"A":"a1","B":"b1"}}`
	jsonB1, _ := djson.NewJsonGo(c1Str)
	testUtil.Check(t, "jsonB1.Str:%s", jsonB1, c1Str)
	testUtil.Check(t, "jsonB1.Get(\"A1.A\"):%s", jsonB1.StrN(testUtil.FailDef, "A1.A"), testUtil.FailDef)
	testUtil.Check(t, "jsonB1.Get(\"A1.A\"):%s", jsonB1.StrN(testUtil.FailDef, "A1", "A"), "a1")
	testUtil.Check(t, "jsonB1.Get(\"@A1.A\"):%s", jsonB1.StrN(testUtil.FailDef, "@A1.A"), "a1")
	dlog.Info("-------------------------Set-----------------------------")
	jsonB1.Set("bb", "A1.A")
	c1Str1 := `{"A1":{"A":"a1","B":"b1"},"A1.A":"bb"}`
	testUtil.Check(t, "jsonB1.Str:%s", jsonB1, c1Str1)
	testUtil.Check(t, "jsonB1.Get(\"A1.A\"):%s", jsonB1.StrN(testUtil.FailDef, "A1.A"), "bb")
	dlog.Info("-------------------------ReNew-----------------------------")
	jsonB1.ReNew(c1Str)
	testUtil.Check(t, "jsonB1.Str:%s", jsonB1, c1Str)
	dlog.Info("-------------------------Set-----------------------------")
	jsonB1.Set(12.3, "@A1.x.d")
	testUtil.Check(t, "jsonB1.Str:%s", jsonB1, testUtil.UnCheck)
	testUtil.Check(t, "jsonB1.Get:%s", jsonB1.StrN(testUtil.FailDef, "@A1.x.d"), "12.3")
	jsonB1.Set(12.3, "@A3.0.x.d")
	jsonB1.Set(33, "@A3.1.x.d")
	testUtil.Check(t, "jsonB1.Str:%s", jsonB1, testUtil.UnCheck)
}

func TestJsonGo4(t *testing.T) {
	dlog.Info("-------------------------Set out of index-----------------------------")
	c1Str := `{"A1":{"A":"a1","B":"b1"}}`
	jsonB1, _ := djson.NewJsonGo(c1Str)
	jsonB1.Set(33, "@A3.3.x.d")
	testUtil.Check(t, "jsonB1.Str:%s", jsonB1, `{"A1":{"A":"a1","B":"b1"},"A3":[]}`)
	dlog.Info("-------------------------reset test -----------------------------")
	jsonB1.ReNew("[12.3]")
	testUtil.Check(t, "jsonB1:%s", jsonB1, "[12.3]")
	jsonB1.Set(12.33, 1)
	testUtil.Check(t, "jsonB1:%s", jsonB1, "[12.3,12.33]")
	testUtil.Check(t, "jsonB1:%s", jsonB1.StrN(testUtil.FailDef, 0), "12.3")
	jsonB1.Set(666, 1)
	testUtil.Check(t, "jsonB1:%s", jsonB1.StrN(testUtil.FailDef, 1), "666")
	dlog.Info("-------------------------errKey index-----------------------------")
	testUtil.Check(t, "jsonB1:%s", jsonB1.StrN(testUtil.FailDef, 9), testUtil.FailDef)
	testUtil.Check(t, "jsonB1:%s", forceVal.Int("xxx", 9), "9")
	testUtil.Check(t, "jsonB1.StrN(\"@A2.3.x.d\"):%s", jsonB1.StrN(testUtil.FailDef, "@A2.3.x.d"), testUtil.FailDef)
}

func TestRemove(t *testing.T) {
	input := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}]}`
	jsonB1, _ := djson.NewJsonGo(input)
	testUtil.Check(t, "jsonB1:%s", jsonB1, input)
	dlog.Info("------------------------Set to last------------------------------")
	jsonB1.Set(1, "@A4.-2")
	jsonB1.Set(2, "@A4.-2")
	jsonB1.Set(3, "@A4.-2")
	strSet2 := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":[1,2,3]}`
	testUtil.Check(t, "jsonB1:%s", jsonB1, strSet2)
	dlog.Info("------------------------remove------------------------------")
	jsonB1.Remove("@A4.-2")
	strRemove1 := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":[1,2]}`
	testUtil.Check(t, "jsonB1:%s", jsonB1, strRemove1)
	jsonB1.Remove("@A4")
	testUtil.Check(t, "jsonB1:%s", jsonB1, input)
	dlog.Info("------------------------Set to first------------------------------")
	jsonB1.Set(1, "A4", "-1")
	jsonB1.Set(2, "A4", "-1")
	jsonB1.Set(3, "A4", "-1")
	strfirst1 := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":[3,2,1]}`
	testUtil.Check(t, "jsonB1:%s", jsonB1, strfirst1)
	dlog.Info("------------------------remove------------------------------")
	jsonB1.Remove("@A4.-1")
	strfirst2 := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":[2,1]}`
	testUtil.Check(t, "jsonB1:%s", jsonB1, strfirst2)
	jsonB1.Remove("@A4")
	testUtil.Check(t, "jsonB1:%s", jsonB1, input)
}

func TestNewJsonFile(t *testing.T) {
	input := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}]}`
	jsonB1, _ := djson.NewJsonGo(input)
	file, _ := djson.NewJsonFile("./test.json", jsonB1)
	dlog.Info("------------------------un onload ------------------------------")
	testUtil.Check(t, "jsonB1:%s", jsonB1, input)
	testUtil.Check(t, "file:%s", file, input)
	dlog.Info("------------------------save------------------------------")
	file.SaveFormat()
	jsonB1.Set(12.3, "@A4.www")
	strSet := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":{"www":12.3}}`
	testUtil.Check(t, "file:%s", file, strSet)
	testUtil.Check(t, "jsonB1:%s", jsonB1, strSet)
	dlog.Info("------------------------reload------------------------------")
	file.ReLoad()
	testUtil.Check(t, "file:%s", file, input)
	testUtil.Check(t, "jsonB1:%s", jsonB1, input)
	dlog.Info("------------------------Set to first------------------------------")
	file.Set(1, "@A4.-1")
	file.Set(2, "@A4.-1")
	file.Set(3, "@A4.-1")
	strSet2 := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":[3,2,1]}`
	testUtil.Check(t, "file:%s", file, strSet2)
	testUtil.Check(t, "jsonB1:%s", jsonB1, strSet2)
	dlog.Info("------------------------remove first------------------------------")
	file.Remove("@A4.-1")
	strRemove1 := `{"A1":{"A":"a1"},"A2":[{"x":{"d":12.3}}],"A4":[2,1]}`
	testUtil.Check(t, "file:%s", file, strRemove1)
	testUtil.Check(t, "jsonB1:%s", jsonB1, strRemove1)
	file.Remove("@A4")
	testUtil.Check(t, "file:%s", file, input)
	testUtil.Check(t, "jsonB1:%s", jsonB1, input)
}

//func TestInt(t *testing.T) {
//
//	a := 123.3
//	b := djson.N0.IntN(a)
//	c := djson.N0.FloatN(a)
//	d := int(c)
//	testUtil.Check(t, "a:%v", a, "123.3")
//	//testUtil.Check(t, "b:%v", b, djson.N0.ToStr())
//	testUtil.Check(t, "c:%v", c, "123.3")
//	testUtil.Check(t, "d:%v", d, "123")
//	testUtil.Check(t, "float64(b):%v", float64(b), testUtil.UnCheck)
//	testUtil.Check(t, "float64(a):%v", float64(a), testUtil.UnCheck)
//	dlog.Info(c == float64(b))
//
//}
func TestInt2(t *testing.T) {
	testUtil.Check(t, "N0:%v", forceVal.Int(float64(21.1)), testUtil.UnCheck)
	//testUtil.Check(t, "StrEmpty:%v",  djson.StrEmpty.ToStr(), testUtil.UnCheck)
	//testUtil.Check(t, "StrEmpty:%v",  djson.N0.IntN(0), testUtil.UnCheck)
	//testUtil.Check(t, "BoolFalse:%v",  djson.BoolFalse.ToStr(), testUtil.UnCheck)

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
