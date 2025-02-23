package config

import "github.com/spf13/viper"

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
}

var AppConfig Config

func LoadConfig() {
	viper.SetDefault("DATABASE_URL", "postgres://postgres:postgres@localhost:55432/task_db")
	viper.SetDefault("JWT_SECRET", "your-secret-key")
	viper.AutomaticEnv()

	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic("Failed to load config: " + err.Error())
	}
}
