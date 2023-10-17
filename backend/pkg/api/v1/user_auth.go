package v1

import (
	"backend/pkg/model/req"
	"backend/pkg/utils"
	"backend/pkg/utils/r"

	"github.com/gin-gonic/gin"
)

type UserAuth struct{}

// 登录
func (*UserAuth) Login(c *gin.Context) {
	loginReq := utils.BindValidJson[req.Login](c)
	loginVo, code := userService.Login(c, loginReq.Username, loginReq.Password)
	r.SendData(c, code, loginVo)
}
