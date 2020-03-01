package domin

import "github.com/gin-gonic/gin"

//Cors 跨域访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//跨域设置
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Action, Module, X-PINGOTHER, Content-Type, Content-Disposition")
		c.Next()
	}
}
