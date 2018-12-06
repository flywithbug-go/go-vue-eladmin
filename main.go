package main

import (
	"doc_manager/server"
	log "github.com/flywithbug/log4go"
)
func SetLog() {
	w := log.NewFileWriter()
	w.SetPathPattern("./log/log-%Y%M%D.log")
	c := log.NewConsoleWriter()
	c.SetColor(true)
	log.Register(w)
	log.Register(c)
	//log.SetLevel(config.Conf().LogLevel%4)
	log.SetLayout("2006-01-02 15:04:05")
}

func main()  {
	SetLog()
	server.Start(":8081",nil)
}
