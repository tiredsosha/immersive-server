package server

import (
	"immersive_server/tools"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var config *tools.Conf

func routes() {
	router.GET("/ping", ping)
	router.GET("/", index)
	router.POST("/themes", themeSender)
}

func Start(cfg *tools.Conf) {
	gin.SetMode(gin.ReleaseMode)

	config = cfg
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")
	routes()
	router.Run(":" + cfg.HttpPort)
	tools.Debug.Println("http is started")
}
