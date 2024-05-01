package config

import (
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type ServerConfig struct {
	Listen         string
	MaxClients     uint64
	MaxLen         uint64
	Timeout        time.Duration
	DataTimeout    time.Duration
	TTL            time.Duration
	AliveSpan      time.Duration
	TrustedDevices []string
}

var (
	serverConfig ServerConfig
)

func InitServerConfig(configPath string) (*ServerConfig, error) {
	tomlData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	if _, err = toml.Decode(string(tomlData), &serverConfig); err != nil {
		return nil, err
	}
	return &serverConfig, nil
}

func GetServerConfig() *ServerConfig {
	return &serverConfig
}
