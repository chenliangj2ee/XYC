package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.TplName = "index.html"
}

func (c *MainController) Shop() {
	c.TplName = "shop.html"
}

func (c *MainController) About() {
	c.TplName = "about-us.html"
}

func (c *MainController) Blog() {
	c.TplName = "blog.html"
}
func (c *MainController) Contact() {
	c.TplName = "contact-us.html"
}
