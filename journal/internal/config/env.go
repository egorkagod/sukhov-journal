package config

import (
	"errors"
	// "os"
)

type Config struct {
	TTSPath  string
	VoiceURL string
	JWTSecret []byte
}

func Load() (*Config, error) {
	cnf := &Config{}

	// if cnf.VoiceURL = os.Getenv("VOICE_URL"); cnf.VoiceURL == "" {
	// 	return nil, errors.New("В окружении нет ссылки на сервис озвучки")
	// }

	if cnf.TTSPath == "" {
		cnf.TTSPath = "/voice/tts"
	}

	// secretStr := os.Getenv("JWT_SECRET")
	secretStr := "TOPSECRETtbank"
	if secretStr == "" {
		return nil, errors.New("В окружении нет JWT секрета")
	}
	cnf.JWTSecret = []byte(secretStr)

	return cnf, nil
}
