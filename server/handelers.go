package server

import (
	"immersive_server/tools"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hypebeast/go-osc/osc"
)

type Request struct {
	command string `json:"command" binding:"required"`
}

func themeSender(conn *gin.Context) {
	var req Request
	conn.BindJSON(&req)

	tools.Debug.Println("sending command: " + req.command + " to server")

	// Пойнтер? наверное да пойнтер на осц порт и айпи
	client := osc.NewClient(ip, port)
	msg := osc.NewMessage("/" + req.command)
	msg.Append(0.5)
	client.Send(msg)

	conn.Data(http.StatusOK, "application/text", []byte("OK")) // Your custom response here
}

func ping(conn *gin.Context) {
	conn.Data(http.StatusOK, "application/text", []byte("pong")) // Your custom response here
}

func index(conn *gin.Context) {
	conn.Data(http.StatusOK, "application/text", []byte("pong")) // Your custom response here
}
