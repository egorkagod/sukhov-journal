package voice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"journal/internal/config"
)

type VoiceManager struct {
	VoiceURL  string
	AudioPath string
}

var Manager *VoiceManager

func InitManager(config *config.Config) {
	Manager = &VoiceManager{VoiceURL: config.VoiceURL}
}

func (m *VoiceManager) VoiceOver(text string) ([]byte, error) {
	body := map[string]string{"text": text}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf("%v/voice/tts", m.VoiceURL), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Ошибка при получении аудио озвучки, код ошибки - %v", resp.StatusCode)
	}

	audioData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return audioData, nil
}

func (m *VoiceManager) SaveAudio(data []byte, articleID uint64) (string, error) {
	filename := fmt.Sprintf("%v.wav", articleID)
	filepath := filepath.Join(m.AudioPath, filename)
	err := os.WriteFile(filepath, data, 0644)
	if err != nil {
		return "", err
	}
	return filepath, nil
}
