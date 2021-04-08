package core

import (
	"blog/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

func init() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	fmt.Println(global.GVA_CONFIG)
	global.GVA_VP = v
}
