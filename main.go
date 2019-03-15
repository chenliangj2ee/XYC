package main

import (
	_ "XYC/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run("192.168.6.153:8080")
}

