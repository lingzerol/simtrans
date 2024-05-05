package config

import (
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/lingzerol/simtrans/library/utils"
)

type DeviceConfig struct {
	Name      string
	SecretKey string
}

type DatabaseConfig struct {
	DBType   string
	Host     string
	Port     string
	UserName string
	PassWord string
	DBName   string
}

type ServerConfig struct {
	ServiceID       uint64
	Listen          string
	MaxClients      uint64
	MaxLen          uint64
	Timeout         time.Duration
	DataTimeout     time.Duration
	TTL             time.Duration
	AliveSpan       time.Duration
	TrustedDevices  map[string]DeviceConfig
	TrustLogin      bool
	CommonSecretKey string
	DBConfig        DatabaseConfig
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
	if serverConfig.ServiceID <= 0 {
		serverConfig.ServiceID, _ = utils.RandomID()
	}
	return &serverConfig, nil
}

func GetServerConfig() *ServerConfig {
	return &serverConfig
}

func GetSecretKey(deviceName string) string {
	serverConfig := GetServerConfig()
	secretKey := serverConfig.CommonSecretKey
	if serverConfig.TrustLogin {
		deviceConfig, ok := serverConfig.TrustedDevices[deviceName]
		if !ok {
			return ""
		}
		secretKey = deviceConfig.SecretKey
	}
	return secretKey
}
