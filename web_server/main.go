package main

import (
	"flag"
	"fmt"
	"os"
	"vue-admin/web_server/config"
	"vue-admin/web_server/core/jwt"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/mail"
	"vue-admin/web_server/model/a_mongo_index"
	"vue-admin/web_server/model/shareDB"
	"vue-admin/web_server/server"
	"vue-admin/web_server/server/handler/file_handler"

	log "github.com/flywithbug/log4go"
)

//log 启动配置
func setLog() {
	w := log.NewFileWriter()
	w.SetPathPattern(config.Conf().LogPath)
	c := log.NewConsoleWriter()
	c.SetColor(true)
	log.Register(w)
	log.Register(c)
	log.SetLevel(1)
	log.SetLayout("2006-01-02 15:04:05")
}

func setFileConfig() {
	file_handler.SetLocalImageFilePath(config.Conf().LocalFilePath)
}

func setJWTKey() {
	//signingKey read
	jwt.ReadSigningKey(config.Conf().PrivateKeyPath, config.Conf().PublicKeyPath)
}

func init() {
	fmt.Println("----------------init----------------")
	//配置文件
	configPath := flag.String("config", "config.json", "Configuration file to use")
	flag.Parse()
	err := config.ReadConfig(*configPath)
	if err != nil {
		panic(err)
	}

}

func setMongoDB() {
	//mongodb启动连接
	//设置数据库名字
	shareDB.SetDBName(config.Conf().DBConfig.DBName)
	mongo.DialMgo(config.Conf().DBConfig.Url)
	mongo_index.CreateMgoIndex()
}

func setMail() {
	var err error
	mail.Mail, err = config.Conf().MailConfig.Dialer()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
}

func main() {
	//log配置
	setLog()
	defer log.Close()
	//文件存储位置
	setFileConfig()
	setMail()

	//jwt验证
	setJWTKey()

	setMongoDB()

	//启动ApiServer服务
	server.StartServer(config.Conf().Port, config.Conf().StaticPath, config.Conf().RouterPrefix, config.Conf().AuthPrefix)
}
