package auth

import (
	"github.com/chenfanlinux/gocmdb/server/controllers/base"
	"github.com/chenfanlinux/gocmdb/server/models"
)

type LoginRequiredController struct {
	base.BaseController
	User *models.User
}


func (c *LoginRequiredController) Prepare() {
	c.BaseController.Prepare()

	if user := DefaultManager.IsLogin(c); user == nil {
		// 未登陆
		DefaultManager.GoToLoginPage(c) // todo 需要修改参数
		c.StopRun()

	} else {
		// 已登陆
		c.User = user
		c.Data["user"] = user
	}
}

// 登录验证
type AuthController struct {
	base.BaseController
}

func (c *AuthController) Login(){
	DefaultManager.Login(c)
}

func (c *AuthController) Logout(){
	DefaultManager.Logout(c)

}
