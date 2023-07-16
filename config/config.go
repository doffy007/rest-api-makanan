package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var ConfigureApp Config

type Config struct {
	Port        int
	Environment string

	DBHost     string
	DBPort     int
	DBName     string
	DBUsername string
	DBPassword string

	DBConnectionIdle     int
	DBConnectionLifetime int
	DBMaxIdle            int
	DBMaxOpen            int
}

func Init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("[INIT] Unable to load configuration. %+v\n", err)
	}

	// Port
	ConfigureApp.Port = viper.GetInt("PORT")

	// Environment
	if ConfigureApp.Environment = viper.GetString("ENVIRONMENT"); ConfigureApp.Environment == "" {
		log.Fatalln("[INIT] ENVIRONMENT is missing in .env")
	}

	// DBHostname
	if ConfigureApp.DBHost = viper.GetString("DB_HOST"); ConfigureApp.DBHost == "" {
		log.Fatalln("[INIT] DB_HOST is missing in .env")
	}

	// DBPort
	ConfigureApp.DBPort = viper.GetInt("DB_PORT")

	// DBDatabase
	if ConfigureApp.DBName = viper.GetString("DB_NAME"); ConfigureApp.DBName == "" {
		log.Fatalln("[INIT] DB_NAME is missing in .env")
	}

	// DBUsername
	if ConfigureApp.DBUsername = viper.GetString("DB_USERNAME"); ConfigureApp.DBUsername == "" {
		log.Fatalln("[INIT] DB_USERNAME is missing in .env")
	}

	// DBPassword
	if ConfigureApp.DBPassword = viper.GetString("DB_PASSWORD"); ConfigureApp.DBPassword == "" {
		fmt.Println("[INIT] DB_PASSWORD is missing in .env")
	}

	if ConfigureApp.DBConnectionIdle = viper.GetInt("DB_CONNECTION_IDDLE"); ConfigureApp.DBConnectionIdle == 0 {
		ConfigureApp.DBConnectionIdle = 1
		fmt.Println("[INIT] DB_CONNECTION_IDDLE is missing in .env, set to default")
	}

	if ConfigureApp.DBConnectionLifetime = viper.GetInt("DB_CONNECTION_IDDLE_LIFETIME"); ConfigureApp.DBConnectionLifetime == 0 {
		ConfigureApp.DBConnectionLifetime = 5
		fmt.Println("[INIT] DB_CONNECTION_IDDLE_LIFETIME is missing in .env, set to default")
	}

	if ConfigureApp.DBMaxIdle = viper.GetInt("DB_MAX_IDDLE"); ConfigureApp.DBMaxIdle == 0 {
		ConfigureApp.DBMaxIdle = 30
		fmt.Println("[INIT] DB_MAX_IDDLE is missing in .env, set to default")
	}

	if ConfigureApp.DBMaxOpen = viper.GetInt("DB_MAX_OPEN"); ConfigureApp.DBMaxOpen == 0 {
		ConfigureApp.DBMaxOpen = 90
		fmt.Println("[INIT] DB_MAX_OPEN is missing in .env, set to default")
	}

	log.Printf("[INIT] Configuration loaded from %s\n", viper.ConfigFileUsed())
}
