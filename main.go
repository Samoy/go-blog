package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/samoy/go-blog/models"
	"github.com/samoy/go-blog/pkg/gredis"
	"github.com/samoy/go-blog/pkg/logging"
	"github.com/samoy/go-blog/pkg/setting"
	"github.com/samoy/go-blog/pkg/util"
	"github.com/samoy/go-blog/routers"
)

func init() {
	setting.Setup()
	logging.Setup()
	models.Setup()
	gredis.Setup()
	util.Setup()
}

func main() {

	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			logging.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logging.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logging.Fatalf("Server Shutdown:", err)
	}

	logging.Info("Server exiting")
}
