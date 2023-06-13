package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("settings/config.yaml")
	if err = viper.ReadInConfig(); err != nil {
		fmt.Println("viper.ReadInConfig() failed: ", err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Configuration file has been modified...")

	})
	return
}
