package server

import (
	"immersive_server/tools"
	"strconv"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func routes() {
	router.GET("/ping", ping)
	router.GET("/", index)
	router.POST("/themes", themeSender)
}

func start(port int) {
	router = gin.Default()
	// Обработаем шаблоны вначале, так что их не нужно будет перечитывать
	// ещё раз. Из-за этого вывод HTML-страниц такой быстрый.
	router.LoadHTMLGlob("templates/*")
	routes()
	router.Run(":" + strconv.Itoa(port))
	tools.Debug.Println("http is startesd")
}
