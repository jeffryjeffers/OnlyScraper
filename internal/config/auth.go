package config

import (
	"encoding/json"
	"os"
)

type Auth struct {
	XBC       string `json:"x_bc"`
	UserAgent string `json:"user_agent"`
	AuthID    string `json:"auth_id"`
	Session   string `json:"sess"`
}

func LoadAuth(filePath string) (*Auth, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var auth Auth
	err = json.Unmarshal(bytes, &auth)
	if err != nil {
		return nil, err
	}

	return &auth, nil
}
