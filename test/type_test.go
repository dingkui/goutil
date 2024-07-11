package test

import (
	"fmt"
	"gitee.com/dk83/goutils/zlog"
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
	zlog.Info(ints)
}
