package config

import (
	"fmt"
	"log"

	"github.com/adwinugroho/wedding-management-system/internals/logger"
	"github.com/spf13/viper"
)

type EnvAppConfig struct {
	AppVersion  string `mapstructure:"app_version"`
	AppName     string `mapstructure:"app_name"`
	Port        string `mapstructure:"app_port"`
	AppURL      string `mapstructure:"app_url"`
	Environment string `mapstructure:"environment"`
	JWTSecret   string `mapstructure:"jwt_secret"`
}

type EnvSSOConfig struct {
	GoogleClientID     string `mapstructure:"sso_google_client_id"`
	GoogleClientSecret string `mapstructure:"sso_google_client_secret"`
	GoogleRedirectURL  string `mapstructure:"sso_google_redirect_url"`
}

var (
	AppConfig EnvAppConfig
	SSOConfig EnvSSOConfig
)

var configStruct = map[string]interface{}{
	"app-config":        &AppConfig,
	"postgresql-config": &PostgreSQLConfig,
	"sso-config":        &SSOConfig,
}

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	for key, value := range configStruct {
		if err := viper.Unmarshal(value); err != nil {
			logger.LogFatal(fmt.Sprintf("Error loading config %s, cause: %+v\n", key, err))
		}
		logger.LogInfo(fmt.Sprintf("Config loaded successfully: %s, value: %+v", key, value))
	}
}
