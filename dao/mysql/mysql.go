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

var sqlsFilepath string

const sqlFilesFolderName = "/models/"

// InitData create database and tables from the SQL scripts
// in bluebell/models if they do not exist.
func InitData() (err error) {
	currWorkingDir, err := os.Getwd()
	if err != nil {
		zap.L().Error("os.Getwd fail", zap.Error(err))
	}
	sqlsFilepath = currWorkingDir + sqlFilesFolderName

	if err = createDBIfNotExist(); err != nil {
		zap.L().Error("fail to create DB", zap.Error(err))
		return
	}

	if err = createTablesIfNotExist(); err != nil {
		zap.L().Error("fail to create tables", zap.Error(err))
		return
	}

	if err = insertCommunityIfNotExist(); err != nil {
		zap.L().Error("fail to insert community", zap.Error(err))
		return
	}

	return
}

const createDBFilename = "create_db.sql"

func createDBIfNotExist() error {
	bytes, err := os.ReadFile(sqlsFilepath + createDBFilename)
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

func createTablesIfNotExist() (err error) {
	err = executeSqlScript(sqlsFilepath + createCommunityTableFilename)
	if err != nil {
		return
	}
	err = executeSqlScript(sqlsFilepath + createPostTableFilename)
	if err != nil {
		return err
	}
	err = executeSqlScript(sqlsFilepath + createUserTableFilename)
	return err
}

const insertCommunityTabelFilename = "insert_community_table.sql"

func insertCommunityIfNotExist() (err error) {
	return executeSqlScript(sqlsFilepath + insertCommunityTabelFilename)
}

func executeSqlScript(filepath string) error {
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
