package initializers

import "os"

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`
}

func getEnv(key, defVar string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defVar
}

func LoadConfig(path string) (config Config, err error) {
	DBHost := getEnv("POSTGRES_HOST", "localhost")
	DBUserName := getEnv("POSTGRES_USER", "postgres")
	DBUserPassword := getEnv("POSTGRES_PASSWORD", "postgres")
	DBName := getEnv("POSTGRES_DB", "imba")
	DBPort := getEnv("POSTGRES_PORT", "5432")
	ServerPort := getEnv("PORT", "8080")

	config = Config{
		DBHost:         DBHost,
		DBUserName:     DBUserName,
		DBUserPassword: DBUserPassword,
		DBName:         DBName,
		DBPort:         DBPort,
		ServerPort:     ServerPort,
	}
	return
}
