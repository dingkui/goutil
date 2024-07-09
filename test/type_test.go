package test

import (
	"fmt"
	"reflect"
	"testing"
)

type typeTest1 struct {
	name string
}
type typeTest1Child struct {
	typeTest1
}

func TestTypeTest1_1(t *testing.T) {
	child := &typeTest1Child{}
	child.name = "t2"
	parent := child.typeTest1 // 正确地进行类型断言
	parent.name = "t1"
	fmt.Printf("parent:%+v t2:%+v", parent, child) // 注意换行符的位置
}
func TestTypeTest1_2(t *testing.T) {
	child := &typeTest1Child{}
	child.name = "t2"
	parent := &child.typeTest1 // 正确地进行类型断言
	parent.name = "t1"
	fmt.Printf("parent:%+v t2:%+v", parent, child) // 注意换行符的位置
}
func checkType(t interface{}) {
	value, ok := t.(typeTest1Child)
	if ok {
		fmt.Println("t is of type typeTest1Child:", value)
	} else {
		fmt.Println("t is not of type typeTest1Child:")
	}
	value2, ok := t.(typeTest1)
	if ok {
		fmt.Println("t is of type typeTest1:", value2)
	} else {
		fmt.Println("t is not of type typeTest1:")
	}
	var d typeTest1Child
	var b interface{} = d
	fmt.Printf("t type is:%s", reflect.TypeOf(t).String())
	if reflect.TypeOf(b).String() == "main.Base" {
		fmt.Println("b is exactly of type Base")
	} else {
		fmt.Println("b is not of type Base")
	}

	if _, ok := b.(typeTest1); ok {
		fmt.Println("b is of type typeTest1")
	} else {
		fmt.Println("b is not of type typeTest1")
	}
}
func checkType2(t interface{}) {
	switch t.(type) {
	case typeTest1:
		fmt.Println("checkType2 typeTest1")
	case typeTest1Child:
		fmt.Println("checkType2 typeTest1Child")
	}
}
func checkType3(t interface{}) {
	switch t.(type) {
	case *typeTest1:
		fmt.Println("checkType3 指针: typeTest1")
	case typeTest1:
		fmt.Println("checkType3 类型: typeTest1")
	case *typeTest1Child:
		fmt.Println("checkType3 指针: typeTest1Child")
	case typeTest1Child:
		fmt.Println("checkType3 类型: typeTest1Child")
	}
}
func TestCheckType(t *testing.T) {
	child := typeTest1Child{}
	child.name = "t2"
	checkType(child) // 注意换行符的位置
}
func TestCheckType2(t *testing.T) {
	child := &typeTest1Child{}
	child.name = "t2"
	checkType3(child)            // 注意换行符的位置
	checkType3(&child.typeTest1) // 注意换行符的位置
}
