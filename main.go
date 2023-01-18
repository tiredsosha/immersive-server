package main

import (
	"immersive_server/server"
	"immersive_server/tools"
)

const (
	version = "1.0.0.alpha"
)

func main() {
	tools.LogInit()

	cfg := tools.ConfInit()
	tools.DebugLog(version, cfg.HttpPort, cfg.OscPort, cfg.OscIp)

	server.Start(cfg)
}
