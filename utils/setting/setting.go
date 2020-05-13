package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string

	DbInfo Database
)

type Base struct {
	RunMode string
}

type APP struct {
	JwtSecret string
}

type Server struct {
	HttpPort     int
	ReadTimeout  int
	WriteTimeout int
}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

type LoadSetting struct {
	Base
	APP
	Server
	Database
}

//#debug or release
//RUN_MODE = debug
//
//[app]
//PAGE_SIZE = 10
//JWT_SECRET = 23347$040412
//
//[server]
//HTTP_PORT = 8081
//READ_TIMEOUT = 60
//WRITE_TIMEOUT = 60
//
//[database]
//TYPE = postgres
//USER = postgres
//PASSWORD = 0852
//#127.0.0.1:3306
//HOST = 127.0.0.1:5432
//NAME = postgres

//var Conf = LoadSetting{
//	Base:     Base{"debug"},
//	APP:      APP{"23347$040412"},
//	Server:   Server{8080, 60, 60},
//	Database: Database{"postgres", "postgres", "0852", "postgresql", "5434", "postgres"},
//}

func init() {
	//RunMode = Conf.RunMode
	//
	//HTTPPort = Conf.HttpPort
	//ReadTimeout = time.Duration(Conf.ReadTimeout) * time.Second
	//WriteTimeout = time.Duration(Conf.WriteTimeout) * time.Second
	//
	//JwtSecret = Conf.JwtSecret
	//
	//DbInfo = Database{
	//	Type:     Conf.Type,
	//	User:     Conf.User,
	//	Password: Conf.Password,
	//	Host:     Conf.Host,
	//	Port:     Conf.Port,
	//	DbName:   Conf.DbName,
	//}
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'config/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
