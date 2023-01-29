package http

import (
	"ccloud/web/http/handler"
	"ccloud/web/log"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HttpServerConfig struct {
	IP   string
	Port int
}

type HttpServer interface {
	Config(HttpServerConfig)
	Init()
	Start()
	Shutdown()
	RegisterHandler(handler.HttpHandler)
}

type httpserver struct {
	mux    *http.ServeMux
	config HttpServerConfig
}

func NewHttpServer() HttpServer {
	return &httpserver{}
}

func (server *httpserver) Config(config HttpServerConfig) {
	log.Logger.Info("start config log")
	server.config.IP = config.IP
	server.config.Port = config.Port
	log.Logger.Info("config log completed")
}

func (server *httpserver) Init() {
	log.Logger.Info("start initliazing log")
	server.mux = http.NewServeMux()
	log.Logger.Info("initliazing log completed")
}

func (server *httpserver) Start() {
	log.Logger.Info("start http server")
	addr := server.buildAddr()

	srv := &http.Server{
		Addr:    addr,
		Handler: server.mux,
	}

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-done
		log.Logger.Info("shutdown http server")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Logger.Fatal("Shutdown server:", err)
		}
		log.Logger.Info("shutdown http server completed")
	}()

	err := srv.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Logger.Info("Server closed under request")
		} else {
			log.Logger.Info("Server closed unexpected")
		}
	}
	log.Logger.Info("start http server completed")
}

func (server *httpserver) RegisterHandler(handler handler.HttpHandler) {
	server.mux.Handle(handler.Pattern(), handler.Handler())
}

func (server *httpserver) Shutdown() {
}

func (server httpserver) buildAddr() string {
	return server.config.IP + ":" + fmt.Sprintf("%v", server.config.Port)
}
