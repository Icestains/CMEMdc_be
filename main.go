package main

import (
	"CMEMdc_be/Database/Postgresql"
	_ "CMEMdc_be/Database/Postgresql"
	"CMEMdc_be/config"
	"CMEMdc_be/router"
	"CMEMdc_be/utils/mqtt"
)

func main() {

	//加载配置信息
	cfg := config.GetConfig()

	//建立数据库链接
	dbConn, err := Postgresql.NewDatabase(&cfg.Database)
	if err != nil {
		panic(err.Error())
	}
	//关闭数据库
	defer dbConn.Close()

	mqtt.TestMqtt()

	//初始化路由
	router.Init(cfg.AppHost, cfg.AppPort, dbConn)
}
