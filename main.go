package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/routes"
	"bluebell/settings"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/spf13/viper"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Println("fail to init setting: ", err)
		return
	}

	if err := logger.Init(); err != nil {
		fmt.Println("fail to init logger: ", err)
		return
	}
	defer zap.L().Sync()

	if err := mysql.Init(); err != nil {
		fmt.Println("fail to init mysql: ", err)
		return
	}
	defer mysql.Close()

	if err := redis.Init(); err != nil {
		fmt.Println("fail to init redis: ", err)
		return
	}
	defer redis.Close()

	r := routes.SetUp()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")

}
