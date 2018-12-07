package main

import (
	"doc-manager/config"
	"doc-manager/server"
	"flag"
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
	configPath := flag.String("config", "config.json", "Configuration file to use")
	flag.Parse()
	err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatal("读取配置文件错误:", err.Error())
	}
	conf := config.Conf()
	SetLog()
	defer log.Close()
	config.DailMongo()
	server.Start(conf.ServerPort,conf.RouterPrefix)
}
