package settings

import (
	"errors"
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	Port      int    `mapstructure:"port"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`

	*AuthConfig  `mapstructure:"auth"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type AuthConfig struct {
	AccTokenExpDurInMinute int `mapstructure:"access_token_expire_duration_minute"`
	RefTokenExpDurInHour   int `mapstructure:"refresh_token_expire_duration_hour"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_conns"`
	MaxIdleConns int    `mapstructure:"idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("settings")
	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("viper.ReadInConfig() failed: ", err)
		return
	}

	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Println("viper.Unmarshal failed: ", err)
	}

	if Conf.Mode != "debug" && Conf.Mode != "release" {
		return errors.New("mode should be [debug] or [release]")
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Configuration file has been modified...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Println("viper.Unmarshal failed: ", err)
		}
	})
	return
}
