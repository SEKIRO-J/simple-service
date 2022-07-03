package configs

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	Env            string `mapstructure:"ENV"`
	Network        string `mapstructure:"NETWORK"`
	FlowScannerKey string `mapstructure:"HTTP_EVENT_BROADCASTER_SHARED_SECRET"`
}

func LoadEnvConfig() (EnvConfig, error) {
	var c EnvConfig

	viper.SetDefault("ENV", "dev")
	viper.SetDefault("NETWORK", "testnet")
	viper.SetDefault("HTTP_EVENT_BROADCASTER_SHARED_SECRET", "secret")

	viper.AutomaticEnv()

	err := viper.Unmarshal(&c)
	return c, err
}
