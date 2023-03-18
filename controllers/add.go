package controllers

import (
	"strings"

	"github.com/astaxie/beego"
)

type AddController struct {
	beego.Controller
}

func (c *AddController) Get() {
	c.TplName = "add.tpl"
}

func (c *AddController) Post() {

	var name = c.GetString("name")
	var vin = strings.ToUpper(c.GetString("vin"))
	repo.AddVin(name, vin)
	c.Redirect("/", 302)
}
