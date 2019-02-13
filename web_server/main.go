package main

import (
	"flag"
	"fmt"
	"os"
	"vue-admin/web_server/config"
	"vue-admin/web_server/core/jwt"
	"vue-admin/web_server/core/mongo"
	"vue-admin/web_server/email"
	"vue-admin/web_server/log_writer"
	"vue-admin/web_server/model/a_mongo_index"
	"vue-admin/web_server/model/shareDB"
	"vue-admin/web_server/server"
	"vue-admin/web_server/server/handler/file_handler"

	log "github.com/flywithbug/log4go"
)

//log 启动配置
func setLog() {

	fmt.Println("init log")
	//日志保存到db
	w := log_writer.NewDBWriter()
	log.Register(w)

	//log日志控制台输出
	l := log_writer.NewConsoleWriter()
	l.SetColor(true)
	log.Register(l)

	log.SetLevel(1)
	log.SetLayout("2006-01-02 15:04:05")
	fmt.Println("\tcomplete")

}

func setFileConfig() {
	fmt.Println("设置图片文件存储路径")
	file_handler.SetLocalImageFilePath(config.Conf().LocalFilePath)
}

func setJWTKey() {
	//signingKey read
	fmt.Println("配置JWT 秘钥文件")
	jwt.ReadSigningKey(config.Conf().PrivateKeyPath, config.Conf().PublicKeyPath)
}

func init() {
	fmt.Println("----------------init----------------")
	//配置文件
	configPath := flag.String("config", "config.json", "Configuration file to use")
	flag.Parse()
	fmt.Println("ParseConfig")
	err := config.ReadConfig(*configPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("\tComplete")
}

func setMongoDB() {
	//mongodb启动连接
	//设置数据库名字
	fmt.Println("init mongodb")
	shareDB.SetDocMangerDBName(config.Conf().DBConfig.DBName)
	shareDB.SetMonitorDBName(config.Conf().MonitorDBConfig.DBName)

	fmt.Println("\tdoc_manager")
	err := mongo.RegisterMongo(config.Conf().DBConfig.Url, config.Conf().DBConfig.DBName)
	if err != nil {
		panic(err)
	}
	fmt.Println("\tmonitor")
	err = mongo.RegisterMongo(config.Conf().MonitorDBConfig.Url, config.Conf().MonitorDBConfig.DBName)
	if err != nil {
		panic(err)
	}
	fmt.Println("\tindex")

	//模型唯一索引
	mongo_index.CreateMgoIndex()
	fmt.Println("\tcomplete")

}

func setMail() {
	var err error
	email.Mail, err = config.Conf().MailConfig.Dialer()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
}

func main() {
	setMongoDB()

	//log配置
	setLog()
	defer log.Close()
	//文件存储位置
	setFileConfig()

	setMail()

	//jwt验证
	setJWTKey()
	fmt.Println("Server")
	//启动ApiServer服务
	server.StartServer(
		config.Conf().Port,
		config.Conf().StaticPath,
		config.Conf().RouterPrefix,
		config.Conf().AuthPrefix,
		config.Conf().Env)
}
