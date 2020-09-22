package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Iteration int `yaml:iteration`
}

func NewConf() *Config {
	return &Config{}
}

func GetConf() *Config {

	viper.AddConfigPath(".")
	viper.SetConfigName("conf")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	return conf
}
