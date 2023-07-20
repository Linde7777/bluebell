package mysql

import (
	"bluebell/settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"strings"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	if err = createDBAndTablesIfNotExist(cfg.User, cfg.Password, cfg.Host, cfg.Port); err != nil {
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
const (
	createDBIfNotExist             = "create_db_if_not_exist.sql"
	createCommunityTableIfNotExist = "create_community_table_if_not_exist.sql"
	createPostTableIfNotExist      = "create_post_table_if_not_exist.sql"
	createUserTableIfNotExist      = "create_user_table_if_not_exist.sql"
	insertCommunityTableIfNotExist = "insert_community_table_if_not_exist.sql"
)

func createDBAndTablesIfNotExist(username, password, host string, port int) error {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/", username, password, host, port)
	tempDB, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		zap.L().Error("Failed to connect to MySQL", zap.Error(err))
		return err
	}
	defer tempDB.Close()

	currWorkDir, err := os.Getwd()
	if err != nil {
		zap.L().Error("os.Getwd", zap.Error(err))
		return err
	}

	var sqlFiles = []string{
		createDBIfNotExist,
		createCommunityTableIfNotExist,
		createPostTableIfNotExist,
		createUserTableIfNotExist,
		insertCommunityTableIfNotExist,
	}
	sqlFilePathPrefix := currWorkDir + "/" + sqlFilesFolderName + "/"
	for _, sqlFile := range sqlFiles {
		if err := execSql(tempDB, sqlFilePathPrefix+sqlFile); err != nil {
			zap.L().Error("execSql:", zap.Error(err))
			return err
		}
	}

	return nil
}

func execSql(tempDB *sqlx.DB, sqlFilePath string) error {
	rawContent, err := readString(sqlFilePath)
	if err != nil {
		zap.L().Error("fail to read sql string:", zap.Error(err))
		return err
	}

	rawSqlStrSlice := strings.Split(rawContent, ";")

	// remove the empty string
	sqlStrSlice := rawSqlStrSlice[:len(rawSqlStrSlice)-1]

	for _, sqlStr := range sqlStrSlice {
		_, err = tempDB.Exec(sqlStr)
		if err != nil {
			zap.L().Error("fail to exec sql string:", zap.Error(err))
			return err
		}
	}

	return err
}

func readString(sqlFilePath string) (string, error) {
	bytes, err := os.ReadFile(sqlFilePath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
