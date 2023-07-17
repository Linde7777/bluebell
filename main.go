package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
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
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Println("fail to init setting: ", err)
		return
	}

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Println("fail to init logger: ", err)
		return
	}
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {
			fmt.Println("Fail to sync zap")
		}
	}(zap.L())

	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Println("fail to init mysql: ", err)
		return
	}
	defer mysql.Close()

	if err := mysql.InitData(); err != nil {
		fmt.Println("fail to init data ", err)
		return
	}

	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println("fail to init redis: ", err)
		return
	}
	defer redis.Close()

	err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID)
	if err != nil {
		fmt.Println("fail to init snowflake")
	}

	r := routes.SetUp(settings.Conf.Mode)
	err = r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		zap.L().Fatal("Fail to run router,", zap.Error(err))
		return
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
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
