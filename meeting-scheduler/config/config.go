package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
	SqlFilePath string
}

func LoadConfig() (*Config, error) {
	// env := os.Getenv("APP_ENV")
	// if env == "" {
	// 	env = "dev"
	// }

	// envFile := fmt.Sprintf(".env.%s", env)
	// if _, err := os.Stat(envFile); err == nil {
	// 	_ = godotenv.Load(envFile)
	// } else {
	// 	_ = godotenv.Load()
	// }

	viper.SetConfigName("config.dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	databaseURL := LoadPostgresConfig()

	cfg := &Config{
		ServerPort:  viper.GetString("server.port"),
		DatabaseURL: databaseURL,
		SqlFilePath: viper.GetString("sqlFilePath"),
	}

	log.Printf("Configuration loaded for environment")
	return cfg, nil
}

func LoadPostgresConfig() string {
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	database := viper.GetString("database.name")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	sslmode := viper.GetString("database.sslmode")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if sslmode == "" {
		sslmode = "disable"
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, host, port, database, sslmode)
}

func CreateTable(sqlFilePath string, db *sql.DB) {
	sqlData, err := os.ReadFile(sqlFilePath)
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
	}
	_, err = db.Exec(string(sqlData))
	if err != nil {
		log.Fatalf("Failed to execute SQL script: %v", err)
	}

	fmt.Println("Tables created successfully!")
}
