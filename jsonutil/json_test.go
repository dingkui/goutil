package jsonutil

import (
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

func TestJsonGo1(t *testing.T) {
	a1 := &TypeA{
		A: "a1",
		B: "b1",
	}
	a2 := &TypeA{
		A: "a2",
		B: "b2",
	}
	b1 := &TypeB{
		A1: a1,
		A2: a2,
	}
	jsonB1 := NewJsonGo(b1)

	zlog.Info("a1:%v", a1)
	zlog.Info("a2:%v", a2)
	zlog.Info("b1:%v", b1)
	zlog.Info("jsonB1:%v", jsonB1)
	zlog.Info("jsonB1.Str:%s", jsonB1.Str())

	jsonStr1 := NewJsonGo("aaa")
	jsonFloat1 := NewJsonGo(1.1)
	jsonInt1 := NewJsonGo(123)
	zlog.Info("jsonStr1:%v", jsonStr1)
	zlog.Info("jsonStr2:%v", jsonFloat1)
	zlog.Info("jsonInt1:%v", jsonInt1)
	zlog.Info("jsonStr1.Str:%s", jsonStr1.Str())
	zlog.Info("jsonStr2.Str:%s", jsonFloat1.Str())
	zlog.Info("jsonInt1.Str:%s", jsonInt1.Str())
}
func TestJsonGo2(t *testing.T) {
	b1 := &TypeB{
		A1: &TypeA{
			A: "a1",
			B: "b1",
		},
		A2: &TypeA{
			A: "a2",
			B: "b2",
		},
	}
	jsonB1 := NewJsonGo(b1)
	c2 := &TypeC{}
	jsonB1.As(c2)
	zlog.Info("c2:%s", Data2Json(c2))
	zlog.Info("------------------------------------------------------")
	a21 := &TypeA{}
	jsonB1.GetN("A1").As(a21)
	zlog.Info("a2:%s", Data2Json(a21))
	zlog.Info("------------------------------------------------------")
	a3 := &TypeA2{}
	jsonB1.GetN("A1").As(a3)
	zlog.Info("a3:%s", Data2Json(a3))
	zlog.Info("------------------------------------------------------")
	a31 := &TypeA2{}
	jsonB1.GetN("A1").As(&a31)
	zlog.Info("a31:%s", Data2Json(a31))
	zlog.Info("------------------------------------------------------")
	a4 := make(map[string]interface{})
	jsonB1.GetN("A1").As(&a4)
	zlog.Info("a4:%s", Data2Json(a4))
}
func TestJsonGo3(t *testing.T) {
	c1 := &TypeC{
		A1: &TypeA2{
			A: "a1",
		},
	}
	jsonB1 := NewJsonGo(c1)
	zlog.Info("jsonB1:%v", jsonB1)
	zlog.Info("jsonB1.Str:%s", jsonB1.Str())
	zlog.Info("jsonB1.GetStr(\"A1.A\"):%s", jsonB1.GetN("A1.A"))
	zlog.Info("------------------------------------------------------")
	jsonB1.Set("bbb", "A1.A")
	jsonB1.Set("cc", "A1.C")
	jsonB1.Set("dd", "A1.d")
	zlog.Info("jsonB1.Str:%s", jsonB1.Str())
	zlog.Info("jsonB1.GetStr(\"A1.A\"):%s", jsonB1.GetN("A1.A").Str())
	zlog.Info("------------------------------------------------------")
	c2 := &TypeC{}
	jsonB1.As(c2)
	zlog.Info("c2:%s", Data2Json(c2))
	zlog.Info("------------------------------------------------------")
	a2 := &TypeA{}
	jsonB1.GetN("A1").As(a2)
	zlog.Info("a2:%s", Data2Json(a2))
	zlog.Info("------------------------------------------------------")
	a3 := &TypeA2{}
	jsonB1.GetN("A1").As(a3)
	zlog.Info("a3:%s", Data2Json(a3))
	zlog.Info("------------------------------------------------------")
	a4 := make(map[string]interface{})
	jsonB1.GetN("A1").As(&a4)
	zlog.Info("a4:%s", Data2Json(a4))
}

func TestJsonGo4(t *testing.T) {
	jsonB1 := NewJsonGo("{\"A1\":{\"A\":\"a1\"},\"A2\":[{\"x\":{\"d\":12.3}}]}")
	//zlog.Info("jsonB1.GetN(\"@A1.A\").Str():%s", jsonB1.GetN("@A1.A").Str())
	//zlog.Info("jsonB1.Str(\"@A1.A\"):%s", jsonB1.Str("@A1.A"))
	//zlog.Info("jsonB1.Str(\"@A1\"):%s", jsonB1.Str("@A1"))
	//zlog.Info("jsonB1.Str():%s", jsonB1.Str())
	//zlog.Info("------------------------------------------------------")
	//jsonB1.Set(12.3,"@A1.x.d")
	//zlog.Info("jsonB1.Str():%s", jsonB1.Str())
	//zlog.Info("jsonB1.Str(\"A1.x\"):%s", jsonB1.Str("@A1.x.d"))
	//zlog.Info("------------------------------------------------------")
	jsonB1.Set(12.3, "@A3.0.x.d")
	jsonB1.Set(33, "@A3.1.x.d")
	jsonB1.GetN("@A3.1.x").Set(12.3, "d1")
	zlog.Info("jsonB1.Str():%s", jsonB1.Str())
	zlog.Info("jsonB1.Str(\"@A2.3.x.d\"):%s", jsonB1.Str("@A2.0.x.d"))
	zlog.Info("jsonB1.Str(\"@A2.3.x.d\"):%s", jsonB1.Str("@A3.1.x.d"))
	zlog.Info("------------------------------------------------------")
}

func TestJsonGo5(t *testing.T) {
	//arr := make([]interface{}, 0)
	//arr = append(arr, 12.3)
	jsonB1 := NewJsonGo("[12.3]")
	zlog.Info("jsonB1.Str():%s", jsonB1.Str())
	jsonB1.Set(12.33, 1)
	zlog.Info("jsonB1.Str():%s", jsonB1.Str())
	//zlog.Info("jsonB1.Str():%s", jsonB1.Str(0))
	//zlog.Info("jsonB1.Str(\"@A2.3.x.d\"):%s", jsonB1.Str("@A2.3.x.d"))
	zlog.Info("------------------------------------------------------")
}

func TestNewJsonFile(t *testing.T) {
	jsonB1 := NewJsonGo("{\"A1\":{\"A\":\"a1\"},\"A2\":[{\"x\":{\"d\":12.3}}]}")
	file := NewJsonFile("./test.json", jsonB1)
	zlog.Info("file.Str():%s", file.Str())
	zlog.Info("------------------------save------------------------------")
	file.SaveFormat()
	jsonB1.Set(12.3, "@A4.www")
	zlog.Info("file.Str():%s", file.Str())
	zlog.Info("file.Str():%s", jsonB1.Str())
	file.Reload()
	zlog.Info("------------------------reload------------------------------")
	zlog.Info("file.Str():%s", file.Str())
	file.Set(1, "@A4.-1")
	file.Set(2, "@A4.-1")
	file.Set(3, "@A4.-1")
	zlog.Info("file.Str():%s", file.Str())
	zlog.Info("------------------------remove------------------------------")
	file.Remove("@A4.-1")
	zlog.Info("file.Str():%s", file.Str())
	file.Remove("@A4")
	zlog.Info("file.Str():%s", file.Str())
}
