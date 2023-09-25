package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

type EnvVars struct {
	GO_ENV string `mapstructure:"GO_ENV" env:"GO_ENV"`
	PORT   string `mapstructure:"PORT" env:"PORT"`
	DB_URL string `mapstructure:"DB_URL" env:"DB_URL"`
}

func LoadConfig() (config EnvVars, err error) {

	env := os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("GO_ENV not set.")
	}

	if env == "production" {
		err = loadEnv_Production(&config)
		return
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	err = checkRequiredFields(&config)

	return
}

func checkRequiredFields(config *EnvVars) (err error) {
	fields := reflect.TypeOf(*config)
	values := reflect.ValueOf(*config)

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		envName := field.Tag.Get("env")

		if envName != "" && value.String() == "" {
			err = fmt.Errorf("%s is required.", envName)
			break
		}
	}

	return
}

func loadEnv_Production(config *EnvVars) error {
	config.PORT = os.Getenv("PORT")
	config.DB_URL = os.Getenv("DB_URL")

	return checkRequiredFields(config)
}
