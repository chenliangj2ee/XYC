package main

import (
	_ "XYC/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run("localhost:8080")
}
