package main

import (
	"github.com/blog-service/global"
	"github.com/blog-service/internal/model"
	"github.com/blog-service/internal/routers"
	"github.com/blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init(){
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil{
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error{
	s, err := setting.NewSetting()
	if err != nil{
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil{
		return err
	}
	err = s.ReadSection("App", &global.ServerSetting)
	if err != nil{
		return err
	}

	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil{
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error{
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil{
		return err
	}
	return nil
}