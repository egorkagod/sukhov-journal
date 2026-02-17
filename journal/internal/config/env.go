package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Debug bool

	VoiceURL  string
	AudioPath string

	JWTSecret []byte

	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     int64
}

func getEnv(key string, byDefault string) string {
	value := os.Getenv(key)
	if value == "" && byDefault == "" {
		log.Fatalf("Не найден %v в окружении", key)
	} else if value == "" {
		return byDefault
	}
	return value
}

func Load() (*Config, error) {
	cnf := &Config{}
	_ = godotenv.Load()

	debug, err := strconv.ParseBool(getEnv("DEBUG", "false"))
	if err != nil {
		log.Fatal("Debug должен быть bool")
	}
	cnf.Debug = debug

	cnf.AudioPath = getEnv("AUDIO_PATH", "/var/www/sukhov-jornal/audio")
	cnf.VoiceURL = getEnv("VOICE_URL", "")
	cnf.JWTSecret = []byte(getEnv("JWT_SECRET", ""))

	cnf.DBHost = getEnv("DB_HOST", "localhost")
	cnf.DBUser = getEnv("DB_USER", "localhost")
	cnf.DBPassword = getEnv("DB_PASSWORD", "localhost")
	cnf.DBName = getEnv("DB_NAME", "localhost")
	cnf.DBPort, err = strconv.ParseInt(getEnv("DB_PORT", "localhost"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Порт базы данных должен быть числом")
	}
	return cnf, nil
}
