package configs

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	Name        string `mapstructure:"RDS_DB_NAME"`
	Host        string `mapstructure:"RDS_HOSTNAME"`
	Port        int    `mapstructure:"RDS_PORT"`
	User        string `mapstructure:"RDS_USERNAME"`
	Pwd         string `mapstructure:"RDS_PASSWORD"`
	SSLMode     string `mapstructure:"RDS_SSLMODE"`
	SSLRootCert string `mapstructure:"RDS_SSLROOTCERT"`
}

func LoadDBConfig() (DBConfig, error) {
	var c DBConfig

	viper.SetDefault("RDS_DB_NAME", "")
	viper.SetDefault("RDS_HOSTNAME", "")
	viper.SetDefault("RDS_PORT", 5432)
	viper.SetDefault("RDS_USERNAME", "")
	viper.SetDefault("RDS_PASSWORD", "")
	viper.SetDefault("RDS_SSLMODE", "disable")
	viper.SetDefault("RDS_SSLROOTCERT", "")

	viper.AutomaticEnv()

	err := viper.Unmarshal(&c)
	return c, err
}
