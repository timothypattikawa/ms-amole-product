package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadViper(env string) *viper.Viper {

	v := viper.New()
	v.SetConfigName(fmt.Sprintf("application-%s", env))
	v.AddConfigPath(".")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		return nil
	}

	return v
}
