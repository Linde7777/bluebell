package mysql

import (
	"bluebell/settings"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/exec"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	if err = createDBAndTablesIfNotExist(cfg.User, cfg.Password); err != nil {
		zap.L().Error("createDBAndTablesIfNotExist: ", zap.Error(err))
		return err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName)

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

const sqlFilesFolderName = "models"
const createDBIfNotExist = "create_db_if_not_exist.sql"
const createCommunityTableIfNotExist = "create_community_table_if_not_exist.sql"
const createPostTableIfNotExist = "create_post_table_if_not_exist.sql"
const createUserTableIfNotExist = "create_user_table_if_not_exist.sql"
const insertCommunityTableIfNotExist = "insert_community_table_if_not_exist"

func createDBAndTablesIfNotExist(username, password string) (err error) {
	cmd := exec.Command("mysql -u root -p")
	_, err = cmd.Output()
	if err != nil {
		zap.L().Error("fail to login mysql", zap.Error(err))
		return err
	}

	currWorkDir, err := os.Getwd()
	if err != nil {
		zap.L().Error("os.Getwd", zap.Error(err))
		return err
	}
	sqlFilePath := currWorkDir + "/" + sqlFilesFolderName + "/"

	err = executeSqlInCMD(sqlFilePath + createDBIfNotExist)
	if err != nil {
		zap.L().Error("createDBIfNotExist", zap.Error(err))
		return err
	}

	err = executeSqlInCMD(sqlFilePath + createCommunityTableIfNotExist)
	if err != nil {
		zap.L().Error("createCommunityTableIfNotExist", zap.Error(err))
		return err
	}

	err = executeSqlInCMD(sqlFilePath + createUserTableIfNotExist)
	if err != nil {
		zap.L().Error("createUserTableIfNotExist", zap.Error(err))
		return err
	}

	err = executeSqlInCMD(sqlFilePath + createPostTableIfNotExist)
	if err != nil {
		zap.L().Error("createPostTableIfNotExist", zap.Error(err))
		return err
	}

	err = executeSqlInCMD(sqlFilePath + insertCommunityTableIfNotExist)
	if err != nil {
		zap.L().Error("insertCommunityTableIfNotExist", zap.Error(err))
		return err
	}

	return
}

func executeSqlInCMD(filepath string) (err error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return
	}
	cmd := exec.Command(string(bytes))
	_, err = cmd.Output()
	if err != nil {
		return
	}

	return
}
