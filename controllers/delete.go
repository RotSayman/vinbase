package controllers

import (
	"github.com/astaxie/beego"
)

type DeleteController struct {
	beego.Controller
}

func (c *DeleteController) Post() {
	var id = c.Ctx.Input.Param(":id")
	repo.DeleteVin(id)
	c.Redirect("/", 302)
}
