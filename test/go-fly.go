package main

import (
	"fmt"
	"gitee.com/dk83/goutils/config"
	"gitee.com/dk83/goutils/httputil"
	"gitee.com/dk83/goutils/jsonutil"
	"gitee.com/dk83/goutils/logutil"
	"math/rand"
)

func main() {
	config.InitLog()
	item := jsonutil.SetItem("[1,2,3,4]", 22, 3)
	fmt.Print(item)
	//config.InitLog()
	logutil.ErrorLn(1)

	for i := 0; i < 10000; i++ {
		n := rand.Int31n(5)
		if n == 5 {
			logutil.ErrorLn(n)
		}
	}
	logutil.Info(httputil.GetAsStr("https://chat.chan3d.com/index.html"))
	re := httputil.Get2Folder("d:\\x.jpg1", "https://chat.chan3d.com/static/images/5.jpg")
	logutil.Info("download re:", re)
}
