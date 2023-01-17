package tools

import (
	"io"
	"log"
	"os"
	"syscall"
	"time"
)

var (
	Warn  *log.Logger
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

func LogInit() {
	var out interface{}
	out = io.Discard

	deleteLog := logCreation()
	if deleteLog {
		os.Remove("server.log")
	}

	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		out = file
	}

	Debug = log.New(out.(io.Writer), "DEBUG: ", log.Ldate|log.Ltime)
	Info = log.New(out.(io.Writer), "INFO:  ", log.Ldate|log.Ltime)
	Warn = log.New(out.(io.Writer), "WARN:  ", log.Ldate|log.Ltime)
	Error = log.New(out.(io.Writer), "ERROR: ", log.Ldate|log.Ltime)

	Debug.Println("")
	Debug.Println("")
	Info.Print("IMMERSIVE SERVER STARTED")
}

func DebugLog(version, httpPort, oscPort, oscIp string) {
	Debug.Println("---------------------------")
	Debug.Println("common data:")
	Debug.Printf("\t\tversion  - %s\n", version)
	Debug.Println("- - - - - - - - - - - - - -")
	Debug.Println("сonnection data:")
	Debug.Printf("\t\tOSC IP - %s\n", oscIp)
	Debug.Printf("\t\tOSC Port - %s\n", oscPort)
	Debug.Printf("\t\tHTTP Port - %s\n", httpPort)
	Debug.Println("---------------------------")
}

func logCreation() bool {
	var deleteLog bool = false

	if log, err := os.Stat("server.log"); err == nil {
		createTime := time.Unix(0, log.Sys().(*syscall.Win32FileAttributeData).CreationTime.Nanoseconds())
		currTime := time.Now()
		diff := currTime.Sub(createTime).Milliseconds()

		if diff > 604800000 {
			deleteLog = true
		}
	}
	return deleteLog
}
