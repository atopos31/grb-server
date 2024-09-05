package app

import (
	"context"
	"gvb/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	// Conf 全局配置
	Conf *config.Config
	// Log 日志
	Log *logrus.Logger
	// 程序退出
	Exit bool
)

func Run(engine *gin.Engine) {
	srv := &http.Server{
		Addr:    Conf.Sys.Addr(),
		Handler: engine,
	}
	go func() {
		printLogo()
		Log.Info("\nGRB listening and serving HTTP on: ", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	Exit = true
	Log.Info("Shutting down GRB server...")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		Log.Fatal("GRB server forced to shutdown:", err)
	}
	Log.Info("GRB server exiting")
}

func printLogo() {
	println(`
	
   _____   _____    ____              _____                                     
  / ____| |  __ \  |  _ \            / ____|                                    
 | |  __  | |__) | | |_) |  ______  | (___     ___   _ __  __   __   ___   _ __ 
 | | |_ | |  _  /  |  _ <  |______|  \___ \   / _ \ | '__| \ \ / /  / _ \ | '__|
 | |__| | | | \ \  | |_) |           ____) | |  __/ | |     \ V /  |  __/ | |   
  \_____| |_|  \_\ |____/           |_____/   \___| |_|      \_/    \___| |_|   
                                                                                
                                                                                
	`)
}
