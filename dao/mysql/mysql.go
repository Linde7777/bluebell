package mysql

import (
	"bluebell/settings"
	"fmt"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		settings.Conf.MySQLConfig.User,
		settings.Conf.MySQLConfig.Password,
		settings.Conf.MySQLConfig.Host,
		settings.Conf.MySQLConfig.Port,
		settings.Conf.MySQLConfig.DBName,
	)

	if db, err = sqlx.Connect("mysql", dsn); err != nil {
		zap.L().Error("sqlx.Connect failed: ", zap.Error(err))
		return
	}

	db.SetMaxOpenConns(viper.GetInt("mysql.max_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.idle_conns"))
	return
}

func Close() {
	if err := db.Close(); err != nil {
		zap.L().Fatal("fail to close mysql", zap.Error(err))
	}
}
