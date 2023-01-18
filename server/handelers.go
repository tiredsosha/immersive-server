package server

import (
	"immersive_server/tools"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hypebeast/go-osc/osc"
)

type Request struct {
	Command string `json:"command" binding:"required"`
}

func themeSender(conn *gin.Context) {
	var req Request
	conn.BindJSON(&req)

	OscPort, _ := strconv.Atoi(config.OscPort)

	tools.Debug.Println("sending command: " + req.Command + " to server")
	client := osc.NewClient(config.OscIp, OscPort)
	msg := osc.NewMessage("/" + req.Command)
	msg.Append(0.5)
	client.Send(msg)

	conn.Data(http.StatusOK, "application/text", []byte("OK")) // Your custom response here
}

func ping(conn *gin.Context) {
	conn.Data(http.StatusOK, "application/text", []byte("pong")) // Your custom response here
}

func index(conn *gin.Context) {
	conn.HTML(http.StatusOK, "index.html", []byte("OK"))
}
