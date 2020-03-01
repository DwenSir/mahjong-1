package user

import (
	"github.com/XzavierLuo/mahjong/internal/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"xlib/log"
)

//Login 登陆接口
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		//先获取code
		var Code struct{
			Code string `json:"code"`
		}
		if err := c.BindJSON(&Code); err != nil {
			log.Error("get code error, error:", err)
			c.JSON(http.StatusOK, Reply{Code:router.CodeNoCode, Error:router.ErrNoCode, Data:nil})
			return
		}

		//调用微信接口，获取对应openid
		//如果数据库中有openid，那么直接获取对应openid,如果没有就新建用户，生成id
		//返回数据
	}
}
