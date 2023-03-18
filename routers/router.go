package routers

import (
	"github.com/RotSayman/vinbase/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/Add", &controllers.AddController{})
	beego.Router("/Edit/:id", &controllers.EditController{})
	beego.Router("/Delete/:id", &controllers.DeleteController{}, "*:Post")
}
