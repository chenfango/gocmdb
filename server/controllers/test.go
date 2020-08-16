package controllers

import (
	"fmt"
	"github.com/chenfanlinux/gocmdb/server/controllers/auth"
	"time"
)

type TestController struct {
	auth.LoginRequiredController
}

type TestPageController struct {
	LayoutController
}

func (c *TestPageController) Index(){
	c.Data["expand"] = "system_management"
	c.Data["menu"] = "user_management"
	c.TplName = "test_page/index.html"

}

func (c *TestController) Test(){
	fmt.Println(c.Ctx.Input.RequestBody)
	c.Data["json"] = map[string]interface{}{"now":time.Now()}
	c.ServeJSON()
}