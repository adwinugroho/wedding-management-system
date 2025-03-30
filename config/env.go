package config

import (
	"log"

	"github.com/adwinugroho/wedding-management-system/internals/logger"
	"github.com/spf13/viper"
)

type EnvAppConfig struct {
	AppVersion  string `mapstructure:"app_version"`
	AppName     string `mapstructure:"app_name"`
	Port        string `mapstructure:"app_port"`
	Environment string `mapstructure:"environment"`
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
		log.Println("Loading config: ", key)
		if err := viper.Unmarshal(value); err != nil {
			log.Printf("Error loading config %s, cause: %+v\n", key, err)
			log.Fatal(err)
		}
		log.Printf("%s: %+v\n", key, value)
		logger.LogInfo("Config loaded successfully")
	}
}
