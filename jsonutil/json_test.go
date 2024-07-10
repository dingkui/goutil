package jsonutil

import (
	"fmt"
	"testing"
)

type TypeA struct {
	a string
	B string
}
type TypeB struct {
	a1 TypeA
	A2 TypeA
}

func TestTypeTest1_1(t *testing.T) {
	//a1 := &TypeA{
	//	a: "a1",
	//	B: "b1",
	//}
	a2 := &TypeA{
		a: "a2",
		B: "b2",
	}
	//b1 := &TypeB{
	//	a1: *a1,
	//	A2: *a2,
	//}
	jsonB1 := NewJSON(a2)

	//fmt.Printf("a1:%+v\n", a1)
	//fmt.Printf("a2:%+v\n", a2)
	//fmt.Printf("b1:%+v\n", b1)
	//fmt.Printf("jsonB1:%+v\n", jsonB1)
	fmt.Printf("jsonB1.Str:%s\n", jsonB1.Str())

	//
	//jsonStr1 := NewJSON("aaa")
	//jsonStr2 := NewJSON("123")
	//jsonInt1 := NewJSON(123)
	//fmt.Printf("jsonStr1:%+v\n", jsonStr1)
	//fmt.Printf("jsonStr2:%+v\n", jsonStr2)
	//fmt.Printf("jsonInt1:%+v\n", jsonInt1)
}
func TestJsonData1(t *testing.T) {
	//a1 := &TypeA{
	//	a: "a1",
	//	B: "b1",
	//}
	a2 := &TypeA{
		a: "a2",
		B: "b2",
	}
	//b1 := &TypeB{
	//	a1: *a1,
	//	A2: *a2,
	//}
	jsonB1 := NewJSON(a2)

	//fmt.Printf("a1:%+v\n", a1)
	//fmt.Printf("a2:%+v\n", a2)
	//fmt.Printf("b1:%+v\n", b1)
	//fmt.Printf("jsonB1:%+v\n", jsonB1)
	fmt.Printf("jsonB1.Str:%s\n", jsonB1.Str())

	//
	//jsonStr1 := NewJSON("aaa")
	//jsonStr2 := NewJSON("123")
	//jsonInt1 := NewJSON(123)
	//fmt.Printf("jsonStr1:%+v\n", jsonStr1)
	//fmt.Printf("jsonStr2:%+v\n", jsonStr2)
	//fmt.Printf("jsonInt1:%+v\n", jsonInt1)
}
