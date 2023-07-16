package mysql

import (
	"bluebell/settings"
	"fmt"
	"os"

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

const sqlFilesDir = "../../models"
const createDBFilename = "create_db.sql"

func CreateDBIfNotExist() error {
	bytes, err := os.ReadFile(sqlFilesDir + createDBFilename)
	if err != nil {
		zap.L().Error("os.ReadFile failed: ", zap.Error(err))
		return err
	}
	script := string(bytes)
	_, err = db.Exec(script)
	if err != nil {
		zap.L().Error("db.Exec failed: ", zap.Error(err))
		return err
	}

	return err
}

const createCommunityTableFilename = "create_community_table.sql"
const createPostTableFilename = "create_post_table.sql"
const createUserTableFilename = "create_user_tabel.sql"

func CreateTablesIfNotExist() (err error) {
	err = createTableIfNotExist(sqlFilesDir + createCommunityTableFilename)
	if err != nil {
		return
	}
	err = createTableIfNotExist(sqlFilesDir + createPostTableFilename)
	if err != nil {
		return err
	}
	err = createTableIfNotExist(sqlFilesDir + createUserTableFilename)
	return err
}

func InsertCommunityIfNotExist() {

}

func createTableIfNotExist(filepath string) error {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		zap.L().Error("os.ReadFile failed: ", zap.Error(err))
		return err
	}
	script := string(bytes)
	_, err = db.Exec(script)
	if err != nil {
		zap.L().Error("db.Exec failed: ", zap.Error(err))
		return err
	}

	return err
}

func Close() {
	if err := db.Close(); err != nil {
		zap.L().Fatal("fail to close mysql", zap.Error(err))
	}
}
