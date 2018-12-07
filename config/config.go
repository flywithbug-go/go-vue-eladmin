package config

import (
	"encoding/json"
	"io/ioutil"
	"gopkg.in/gomail.v2"
)

var config *Config

func Conf() *Config {
	if config == nil{
		config = new(Config)
	}
	return config
}


type Config struct {
	ServerPort 		string		`json:"server_port"`			//httpServer
	DBConfig   		DBConfig 	`json:"db_config"`   			//数据库配置
	RSAConfig		RSAConfig									//加密解密
	PrivateName		string		`json:"private_name"`
	RouterPrefix 	[]string 	`json:"router_prefix"` 			//api前缀
	AuthFilterWhite []string 	`json:"auth_filter_white"` 		//白名单
	MailConfig      MailConfig	`json:"mail_config"`  			//邮箱
	AppConfig		AppConfig	`json:"app_config"`

}

type AppConfig struct {
	ApiHost 		string  	`json:"api_host"`    //api请求host
	DomainName		string		`json:"domain_name"`  //域名
	Version			string		`json:"version"`	//版本
}

/*
数据库配置结构体
*/
type DBConfig struct {
	Url         string 	`json:"url"`           //连接地址
	DBName     	string 	`json:"db_name"`       //用户名
	ForceSync	bool	`json:"force_sync"`
}


type MailConfig struct {
	Host      string 	`json:"host"`
	Port      int		`json:"port"`
	Username  string	`json:"username"`
	Password  string	`json:"password"`
}

type RSAConfig struct {
	Public			[]byte
	Private			[]byte
}

/*
读取配置文件
*/
func ReadConfig(path string) error {
	config = new(Config)
	err := config.Parse(path)
	return err
}


/*
解析配置文件
*/
func (this *Config) Parse(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &this)
	if err != nil {
		return err
	}
	b, _ := ioutil.ReadFile(this.PrivateName)
	this.RSAConfig.Private = b
	return nil
}

func (this *MailConfig)Dialer()(*gomail.Dialer,error)  {
	d := gomail.NewDialer(this.Host,this.Port,this.Username,this.Password)
	return d,nil
}
