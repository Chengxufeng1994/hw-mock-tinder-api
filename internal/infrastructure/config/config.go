package config

import (
	"fmt"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type Config struct {
	GinMode  string   `mapstructure:"gin_mode"`
	Logging  Logging  `mapstructure:"logging"`
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
}

type Logging struct {
	Format       string   `mapstructure:"format"`
	Level        string   `mapstructure:"level"`
	Name         string   `mapstructure:"name"`
	Outputs      []string `mapstructure:"outputs"`
	ErrorOutputs []string `mapstructure:"error_outputs"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Database struct {
	LogLevel     string `mapstructure:"log_level"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"database"`
	MaxLifetime  int    `mapstructure:"max_lifetime"`
	MaxIdleTime  int    `mapstructure:"max_idle_time"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

var (
	initOnce sync.Once
)

func InitializationConfig(configFile string) (*Config, error) {
	initOnce.Do(func() {
		if configFile != "" {
			viper.SetConfigFile(configFile)
		} else {
			viper.SetConfigType("yaml")
			viper.SetConfigName("config")
			viper.AddConfigPath(".")
			viper.AddConfigPath(path.Join(getHomeDir(), ".tinder"))
			viper.AddConfigPath("/etc/tinder")
		}

		viper.SetEnvPrefix("TINDER")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // support nested config
		viper.AutomaticEnv()                                   // read in environment variables

		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	})

	var cfg Config

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// getHomeDir find and return the home directory
func getHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Get home directory -", err)
		os.Exit(1)
	}
	return home
}
