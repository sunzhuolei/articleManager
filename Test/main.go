package main

import (
	_ "Test/routers"
	"github.com/astaxie/beego"
	_"Test/models"
	"time"
)
func FormatTime(timestamp int)(formattime string){
	formattime = time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
	return
}
func main() {
	beego.AddFuncMap("FormatTime",FormatTime)
	beego.Run()
}

