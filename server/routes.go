package server

import (
	"immersive_server/tools"

	"github.com/gin-contrib/cors"
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
	config = cfg

	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.Use(cors.Default())

	router.LoadHTMLGlob("templates/*")
	router.Static("/img", "./img")
	router.Static("/css", "./css")
	router.Static("/fonts", "./fonts")

	routes()

	router.Run(":" + cfg.HttpPort)
	tools.Debug.Println("http is started")
}
