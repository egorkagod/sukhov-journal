package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct { //TODO Дописать валидацию DB значений
	TTSPath   string
	VoiceURL  string
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Не удалось загрузить .env файл в окружение")
	}

	cnf.TTSPath = getEnv("TTS_PATH", "/voice/tts")
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
