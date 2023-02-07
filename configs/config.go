package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

// use viper to map .env
func InitializeAppConfig() {
	vp := viper.New()

	vp.SetConfigName("app.env")
	vp.SetConfigType("env")
	vp.AddConfigPath("../app.env")

	err := vp.ReadInConfig()

	if err != nil {
		fmt.Print("Failed in loading the env", err)
	}
}
