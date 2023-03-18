package controllers

import (
	"github.com/RotSayman/vinbase/repository"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

var repo *repository.Repository

func init() {
	repo = repository.GetRepo()
}

func (c *MainController) Get() {
	search := c.GetString("search")
	if search == "" {
		c.Data["vins"] = repo.GetAllVins()
	} else {
		c.Data["Search"] = search
		c.Data["vins"] = repo.FindByName(search)
	}
	c.TplName = "index.tpl"
}
