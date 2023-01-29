package main

import (
	"ccloud/web/http"
	"ccloud/web/http/handler"
	"ccloud/web/log"
	"ccloud/web/security"
	"fmt"
)

func main() {
	passwordencoder := security.NewBcryptPasswordEncoder()
	fmt.Println(passwordencoder.Encode("123456"))
	initlog()
	starthttpserver()
}

func initlog() {
	log.InitLogger()
	defer log.Logger.Sync()
}

func starthttpserver() {
	log.Logger.Info("Start http server.")
	httpconfig := http.HttpServerConfig{
		IP:   "localhost",
		Port: 8080,
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
