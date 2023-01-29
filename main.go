package main

import (
	"ccloud/web/config"
	"ccloud/web/http"
	"ccloud/web/http/handler"
	"ccloud/web/log"
	"ccloud/web/service"
)

func main() {
	initConfig()
	initlog()
	initUser()
	starthttpserver()
}

func initConfig() {
	err := config.GetInstance().LoadConfig()
	if err != nil {
		panic(err)
	}
}

func initlog() {
	log.InitLogger()
	defer log.Logger.Sync()
}

func initUser() {
	loginService, _ := service.NewLoginService()
	username := config.GetInstance().AccountConfig.Username
	password := config.GetInstance().AccountConfig.Password
	err := loginService.AddUser(username, password)
	if err != nil {
		panic(err)
	}
}

func starthttpserver() {
	log.Logger.Info("Start http server.")
	httpconfig := http.HttpServerConfig{
		IP:   config.GetInstance().ServerConfig.Host,
		Port: config.GetInstance().ServerConfig.Port,
	}

	httpserver := http.NewHttpServer()
	httpserver.Config(httpconfig)
	httpserver.Init()
	registerhandler(httpserver)
	httpserver.Start()
	log.Logger.Info("http server start success, ready for connection.")
}

func registerhandler(httpserver http.HttpServer) {
	httpserver.RegisterHandler(handler.NewIndexHandler())

	// 注册login
	login, err := handler.NewLoginHandler()
	if err != nil {
		panic(err)
	}
	httpserver.RegisterHandler(login)

	// 注册文件上传
	upload, err := handler.NewUploadFileHandler()
	if err != nil {
		panic(err)
	}
	wrapUpload := handler.BaseHandler{
		Handle: upload,
	}
	httpserver.RegisterHandler(wrapUpload)
}
