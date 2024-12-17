package app

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	// vn *viper.Viper

	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Score    int    `mapstructure:"score"`
}

func (c *Config) Init(env string) error {
	vn := viper.New()

	vn.AddConfigPath(filepath.Join(".", "configs"))
	vn.SetConfigType("yml")
	vn.SetConfigName(fmt.Sprintf("config.%s.yml", env))

	if err := vn.ReadInConfig(); err != nil {
		return err
	}

	if err := vn.Unmarshal(&c); err != nil {
		return err
	}

	return nil
}
