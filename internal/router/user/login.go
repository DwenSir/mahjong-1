package user

import (
	"encoding/json"
	"github.com/XzavierLuo/mahjong/internal/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
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
			c.JSON(http.StatusOK, router.Reply{Code:router.CodeNoCode, Error:router.ErrNoCode, Data:nil})
			return
		}

		//调用微信接口，获取对应openid
		u, err := url.Parse("https://api.weixin.qq.com/sns/jscode2session")
		if err != nil {
			log.Error("get token error, error:", err)
			c.JSON(http.StatusOK, router.Reply{Code:router.CodeGetToken, Error:router.ErrGetToken, Data:nil})
			return
		}

		paras := &url.Values{}
		paras.Set("appid", "123456")
		paras.Set("secret", "123456")
		paras.Set("grant_type", "authorization_code")
		u.RawQuery = paras.Encode()
		resp, err := http.Get(u.String())
		if err != nil || resp == nil || resp.Body == nil {
			log.Error("get token error, error:", err)
			c.JSON(http.StatusOK, router.Reply{Code:router.CodeGetToken, Error:router.ErrGetToken, Data:nil})
			return
		}
		defer resp.Body.Close()

		var Res struct{
			OpenId string `json:"openid"`
			SessionKey string `json:"session_key"`
			UnionId string `json:"unionid"`
			ErrCode int `json:"errcode"`
			ErrMsg string `json:"errmsg"`
		}

		if err = json.NewDecoder(resp.Body).Decode(&Res); err != nil {
			log.Error("get token error, error:", err)
			c.JSON(http.StatusOK, router.Reply{Code:router.CodeGetToken, Error:router.ErrGetToken, Data:nil})
			return
		}

		//获取失败，返回对应数据
		if Res.ErrCode != 0 {
			log.Error("get token error, error:", err)
			c.JSON(http.StatusOK, router.Reply{Code:router.CodeGetToken, Error:router.ErrGetToken, Data: nil})
			return
		}

		//如果数据库中有openid，那么直接获取对应openid,如果没有就新建用户，生成id
		//返回数据
	}
}
