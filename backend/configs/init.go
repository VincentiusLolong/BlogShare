package configs

import (
	"fmt"
	"log"
	"mailinglist/infrastructure/models"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func AllEnv(s string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	value, errs := os.LookupEnv(s)
	if !errs {
		log.Fatal("Env Variable Empty or Not available")
	}
	return value
}

func DBInit() error {
	viper.SetConfigName("app")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fatal error config file: %s", err)
	}

	models.PostgreSetting.Connstr = viper.GetString("postgresql.PSQLstring")
	models.PostgreSetting.DBname = viper.GetString("postgresql.PSQLname")

	return nil
}
