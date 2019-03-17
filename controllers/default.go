package controllers

import (
	"XYC/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {
	c.TplName = "index.html"
}

func (c *MainController) Shop() {
	c.Data["ps"] = models.GetProduct()
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
