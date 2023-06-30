package main

import (
	"gitee.com/dk83/goutils/config"
	"gitee.com/dk83/goutils/jsonutil"
	"gitee.com/dk83/goutils/logutil"
)

func main() {
	config.InitLog()
	//item := jsonutil.SetItem("[1,2,3,4]", 22, 3)
	//fmt.Print(item)
	////config.InitLog()
	//logutil.Error(1)
	//
	//for i := 0; i < 10000; i++ {
	//	n := rand.Int31n(5)
	//	if n == 5 {
	//		logutil.Error(n)
	//	}
	//}
	//logutil.Info(httputil.GetAsStr("https://chat.chan3d.com/index.html"))
	//re := httputil.Get2Folder("d:\\x.jpg1", "https://chat.chan3d.com/static/images/5.jpg")
	//logutil.Info("download re:", re)

	json := jsonutil.MkJSON("{\"a\":1,\"b\":[{\"a\":\"1\"},{\"a\":\"2\"},{\"a\":\"3\"}]}")
	logutil.Info("json.Str()：", json.Str())
	a, err := json.Array()
	logutil.Info("json.Array()：", a, err)
	b, err := json.Map()
	logutil.Info("json.Map()：", b, err)
	c := &Test{}
	err = json.As(c)
	logutil.Info("json.As()：c", c, err)

	i := json.GetItem("b", 4)
	logutil.Info("json.GetItem(\"b\", 2)：", i)

	json1 := jsonutil.MkJSON("[{\"a\":\"1\"},{\"a\":\"2\"},{\"a\":\"3\"}]")
	getItem, e := json1.Array()
	logutil.Info("json.GetItem(\"b\", 2)：", getItem, e)

	json2 := jsonutil.MkJSON(nil)
	getItem2, e := json2.Array()
	logutil.Info("json.GetItem(\"b\", 2)：", getItem2, e)

}

type Test struct {
	A int           `json:"a"`
	B []interface{} `json:"b"`
}
