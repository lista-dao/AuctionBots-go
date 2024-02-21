package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Commands []string `mapstructure:"commands"`
	Log      struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"log"`
	Wallet struct {
		PrivateKey string `mapstructure:"privateKey"`
	} `mapstructure:"wallet"`
	RpcNode struct {
		Ws   string `mapstructure:"ws"`
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

func LoadConfig(configFile string) (*Config, error) {
	viper.SetConfigFile(configFile)
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
