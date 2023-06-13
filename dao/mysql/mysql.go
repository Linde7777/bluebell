package mysql

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
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
