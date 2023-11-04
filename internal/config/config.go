package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Types struct {
		Timeline    string `json:"timeline"`
		Pinned      string `json:"pinned"`
		Paid        string `json:"paid"`
		PaidPreview string `json:"paidPreview"`
		Highlights  string `json:"highlights"`
		Archived    string `json:"archived"`
		Message     string `json:"message"`
	} `json:"types"`
	MediaType struct {
		Images string `json:"images"`
		Videos string `json:"videos"`
		Audios string `json:"audios"`
	} `json:"mediaTypes"`
}

func LoadConfig(filePath string) (*Config, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
