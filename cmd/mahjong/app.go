package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
	"xlib/log"
)

var wg sync.WaitGroup

//StopServer 当前不需要做什么
func StopServer() {

}

func Run(f func()) {
	wg.Add(1)
	defer wg.Done()
	f()
}

func StartServer() {
	Run(Gin)
	Run(NaNo)
}

func Gin() {
	gin.SetMode(cfg.WebConfig.Mode)
	r := gin.New()
	r.Static("/file", "../../web/file")
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error": "page not found",
		})
	})
	if err := r.Run(cfg.Web.IP + strconv.Itoa(cfg.Web.Port)); err != nil {
		log.Error("run http error, error:", err)
		return
	}
}

func NaNo() {

}
