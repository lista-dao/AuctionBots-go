package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Commands []string `mapstructure:"commands"`
	Log      struct {
		Level  string `mapstructure:"level"`
		Caller bool   `mapstructure:"caller"`
	} `mapstructure:"log"`
	Wallet struct {
		PrivateKey string `mapstructure:"privateKey"`
	} `mapstructure:"wallet"`
	RpcNode struct {
		Http string `mapstructure:"http"`
	} `mapstructure:"rpcNode"`
	Contract struct {
		Interaction string   `mapstructure:"interaction"`
		Collaterals []string `mapstructure:"collaterals"`
		Hay         string   `mapstructure:"hay"`
		FlushBuy    string   `mapstructure:"flushBuy"`
	} `mapstructure:"contract"`
	Settings struct {
		MaxPricePercentage int64 `mapstructure:"maxPricePercentage"`
	} `mapstructure:"settings"`
	Analytics struct {
		Url string `mapstructure:"url"`
	} `mapstructure:"analytics"`
}

func LoadConfig(configFile string, dir string) (*Config, error) {
	configFullPath := filepath.Join(dir, configFile)
	envPath := filepath.Join(dir, ".env")
	configTxtPath := filepath.Join(dir, "config.txt")

	logrus.Infof("configFullPath: %s envPath: %s configTxtPath: %s", configFullPath, envPath, configTxtPath)
	gotenv.Load(envPath)
	gotenv.Load(configTxtPath)

	viper.SetConfigFile(configFullPath)
	viper.SetConfigType("")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	err := viper.ReadInConfig()

	if err != nil {
		return nil, fmt.Errorf("viper.ReadInConfig err: %v", err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)

	return &cfg, err
}
