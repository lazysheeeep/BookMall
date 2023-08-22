package config

import (
	"BookMall/dao"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	Debug    string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	Host       string
	AvatarPath string
	BookPath   string

	SignName   string
	TemplateId string
	SdkAppId   string
	SecretId   string
	SecretKey  string
)

func Init() {
	//读取本地配置文件
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadMysql(file)
	LoadPath(file)
	LoadSms(file)
	//读写分离
	//root:1234@tcp(127.0.0.1:3306)/book_mall?charset=utf8mb4&parseTime=True&loc=Local
	pathRead := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	pathWrite := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	dao.DataBase(pathRead, pathWrite)
}

func LoadServer(file *ini.File) {
	Debug = file.Section("service").Key("Debug").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbName = file.Section("mysql").Key("DbName").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
}

func LoadPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
	BookPath = file.Section("path").Key("BookPath").String()
}

func LoadSms(file *ini.File) {
	SignName = file.Section("sms").Key("SignName").String()
	TemplateId = file.Section("sms").Key("TemplateId").String()
	SdkAppId = file.Section("sms").Key("SdkAppId").String()
	SecretId = file.Section("sms").Key("SecretId").String()
	SecretKey = file.Section("sms").Key("SecretKey").String()
}
