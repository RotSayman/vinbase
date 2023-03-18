package controllers

import (
	"strings"

	"github.com/astaxie/beego"
)

type EditController struct {
	beego.Controller
}

func (c *EditController) Get() {
	var id = c.Ctx.Input.Param(":id")
	var vin = repo.GetVin(id)
	c.Data["Name"] = vin.Name
	c.Data["VinID"] = vin.VinID
	c.TplName = "edit.tpl"
}

func (c *EditController) Post() {
	var id = c.Ctx.Input.Param(":id")
	var name = c.GetString("name")
	var vin = strings.ToUpper(c.GetString("vin"))
	repo.UpdateVin(id, name, vin)
	c.Redirect("/", 302)
}
