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

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
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
