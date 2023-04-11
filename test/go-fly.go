package main

import (
	"fmt"
	"gitee.com/dk83/goutils/jsonutil"
	"gitee.com/dk83/goutils/logutil"
)

func main() {
	item := jsonutil.SetItem("[1,2,3,4]", 22, 3)
	fmt.Print(item)
	//config.InitLog()
	logutil.ErrorLn(1)
}
